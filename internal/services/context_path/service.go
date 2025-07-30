package context_path

import (
	"context"
	"go-skeleton/internal/repositories/context_path"
	"go-skeleton/pkg/utils/null"
	"sort"
)

func (s *contextPathService) GetListPath(ctx context.Context, RoleID uint) (*AuthorizedPathResponse, error) {
	frontendPath, err := s.contextPathRepo.GetByContextTag(ctx, RoleID, "frontend_path")
	if err != nil {
		return nil, err
	}
	backendPath, err := s.contextPathRepo.GetByContextTag(ctx, RoleID, "backend_path")
	if err != nil {
		return nil, err
	}

	frontendList := []string{}
	for _, v := range frontendPath {
		frontendList = append(frontendList, v.Name)
	}
	backendList := []BackendPath{}
	for _, v := range backendPath {
		backendList = append(backendList, BackendPath{Name: v.Name, Method: v.Method})
	}

	return &AuthorizedPathResponse{
		FrontendPath: frontendList,
		BackendPath:  backendList,
	}, nil
}

func (s *contextPathService) CreateRBAC(ctx context.Context, roleID uint, rbacs []map[string]ModulePermission) error {
	// Frontend setup
	var frontendPageGroups []string

	// Backend setup
	backendMap := make(map[string][]string) // map[PageGroup][]Labels

	for _, moduleMap := range rbacs {
		for pageGroup, permission := range moduleMap {
			// Handle frontend_path
			if null.DerefBool(permission.FrontendPath) {
				frontendPageGroups = append(frontendPageGroups, pageGroup)
			}

			// Handle backend_path (each map is an action like read, create, etc.)
			for _, labelMap := range permission.BackendPath {
				for label, enabled := range labelMap {
					if enabled {
						backendMap[pageGroup] = append(backendMap[pageGroup], label)
					}
				}
			}
		}
	}

	// Delete all access paths for the role to upsert
	if err := s.contextPathRepo.DeleteByRoleID(ctx, roleID); err != nil {
		return err
	}

	// Call frontend repo
	if err := s.contextPathRepo.CreateFrontendRBAC(ctx, roleID, frontendPageGroups); err != nil {
		return err
	}

	// Call backend repo
	if err := s.contextPathRepo.CreateBackendRBAC(ctx, roleID, backendMap); err != nil {
		return err
	}

	return nil
}

func (s *contextPathService) GetListRBAC(ctx context.Context, roleID uint) ([]map[string]ModulePermission, error) {
	allData, err := s.contextPathRepo.GetListRBAC(ctx)
	if err != nil {
		return nil, err
	}

	var roleData []context_path.ListRBAC
	if roleID != 0 {
		roleData, err = s.contextPathRepo.GetRBACByRoleID(ctx, roleID)
		if err != nil {
			return nil, err
		}
	}

	// Track all existing page groups in allData
	existingPageGroups := make(map[string]bool)
	for _, v := range allData {
		existingPageGroups[v.PageGroup] = true
	}

	// Create a quick lookup for role permissions
	roleMap := make(map[string]map[string]bool) // pageGroup -> label -> true
	for _, r := range roleData {
		if _, ok := roleMap[r.PageGroup]; !ok {
			roleMap[r.PageGroup] = make(map[string]bool)
		}
		if r.Label == nil {
			roleMap[r.PageGroup]["frontend_path"] = true
		} else {
			roleMap[r.PageGroup][*r.Label] = true
		}
	}

	grouped := make(map[string]ModulePermission)
	labelTracker := make(map[string]map[string]bool) // to avoid duplicate label inserts

	// Process all data from allData (existing page groups)
	for _, v := range allData {
		perm := grouped[v.PageGroup]

		// Ensure labelTracker is initialized
		if labelTracker[v.PageGroup] == nil {
			labelTracker[v.PageGroup] = make(map[string]bool)
		}

		// Set FrontendPath: properly handle existence
		if v.Label == nil {
			if roleMapPageGroup, exists := roleMap[v.PageGroup]; exists {
				value := roleMapPageGroup["frontend_path"]
				perm.FrontendPath = &value
			} else {
				value := false
				perm.FrontendPath = &value
			}
		} else {
			label := *v.Label
			if !labelTracker[v.PageGroup][label] {
				hasAccess := false
				if pageGroupMap, exists := roleMap[v.PageGroup]; exists {
					hasAccess = pageGroupMap[label]
				}
				perm.BackendPath = append(perm.BackendPath, map[string]bool{label: hasAccess})
				labelTracker[v.PageGroup][label] = true
			}
		}

		grouped[v.PageGroup] = perm
	}

	// Handle page groups that are in roleMap but not in allData (non-existent)
	// These should get nil values
	for pageGroup := range roleMap {
		if !existingPageGroups[pageGroup] {
			grouped[pageGroup] = ModulePermission{
				FrontendPath: nil,
				BackendPath:  nil,
			}
		}
	}

	// Flatten to the format: []map[string]ModulePermission
	var result []map[string]ModulePermission
	for pageGroup, permission := range grouped {
		result = append(result, map[string]ModulePermission{
			pageGroup: permission,
		})
	}

	// Sort by pageGroup key
	sort.Slice(result, func(i, j int) bool {
		var ki, kj string
		for k := range result[i] {
			ki = k
		}
		for k := range result[j] {
			kj = k
		}
		return ki < kj
	})

	return result, nil
}

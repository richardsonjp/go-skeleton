package context_path

import (
	"context"
	"fmt"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/utils/errors"
)

func (r *contextPathRepo) Create(ctx context.Context, m model.ContextPath) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *contextPathRepo) GetByID(ctx context.Context, ID uint) (*model.ContextPath, error) {
	contextPath := model.ContextPath{}
	contextPath.ID = ID

	query := r.dbdget.Get(ctx).Where(contextPath)

	if err := query.Find(&contextPath).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &contextPath, nil
}

func (r *contextPathRepo) GetByContextTag(ctx context.Context, roleID uint, contextTag string) ([]*PathList, error) {
	var pathList []*PathList
	contextPathModel := model.ContextPath{}
	q := r.dbdget.Get(ctx).Model(contextPathModel)
	frontendModel := model.FrontendPath{}
	backendModel := model.BackendPath{}

	var qselect []string
	if contextTag == frontendModel.TableName() {
		qselect = []string{
			"frontend_path.name as name",
		}
	} else if contextTag == backendModel.TableName() {
		qselect = []string{
			"backend_path.name as name",
			"backend_path.method as method",
		}
	}

	q.
		Where("context_path.context_tag = ?", contextTag).
		Where("context_path.role_id = ?", roleID).
		Joins(fmt.Sprintf("JOIN %s ON %s.id = context_path.path_id", contextTag, contextTag))

	err := q.
		Select(qselect).
		Find(&pathList).Error
	if err != nil {
		return nil, err
	}

	return pathList, nil
}

func (r *contextPathRepo) DeleteByRoleID(ctx context.Context, roleID uint) error {
	if err := r.dbdget.Get(ctx).
		Delete(&model.ContextPath{}, "role_id = ?", roleID).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *contextPathRepo) GetListRBAC(ctx context.Context) ([]ListRBAC, error) {
	var allRBAC []ListRBAC
	q := r.dbdget.Get(ctx)

	// Get all frontend paths
	var frontendPaths []ListRBAC
	if err := q.Table("frontend_path").
		Select("'frontend_path' AS context_path, page_group, NULL AS label").
		Scan(&frontendPaths).Error; err != nil {
		return nil, fmt.Errorf("error fetching frontend paths: %w", err)
	}

	// Get all backend paths
	var backendPaths []ListRBAC
	if err := q.Table("backend_path").
		Select("'backend_path' AS context_path, page_group, label").
		Scan(&backendPaths).Error; err != nil {
		return nil, fmt.Errorf("error fetching backend paths: %w", err)
	}

	// Combine all paths
	allRBAC = append(frontendPaths, backendPaths...)

	return allRBAC, nil
}

func (r *contextPathRepo) GetRBACByRoleID(ctx context.Context, roleID uint) ([]ListRBAC, error) {
	var allRBAC []ListRBAC
	q := r.dbdget.Get(ctx)

	// Get frontend_path RBAC
	var frontendPaths []ListRBAC
	if err := q.Table("context_path AS cp").
		Select("'frontend_path' AS context_path, fp.page_group, NULL AS label").
		Joins("JOIN frontend_path AS fp ON cp.path_id = fp.id").
		Where("cp.role_id = ? AND cp.context_tag = ?", roleID, "frontend_path").
		Scan(&frontendPaths).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch frontend_path: %w", err)
	}

	// Get backend_path RBAC
	var backendPaths []ListRBAC
	if err := q.Table("context_path AS cp").
		Select("'backend_path' AS context_path, bp.page_group, bp.label").
		Joins("JOIN backend_path AS bp ON cp.path_id = bp.id").
		Where("cp.role_id = ? AND cp.context_tag = ?", roleID, "backend_path").
		Scan(&backendPaths).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch backend_path: %w", err)
	}

	allRBAC = append(frontendPaths, backendPaths...)
	return allRBAC, nil
}

func (r *contextPathRepo) CreateBackendRBAC(ctx context.Context, roleID uint, pageGroupToLabels map[string][]string) error {
	contextTag := "backend_path"

	for pageGroup, labels := range pageGroupToLabels {
		query := `
			INSERT INTO context_path (context_tag, role_id, path_id)
			SELECT ?, ?, id
			FROM backend_path
			WHERE label IN ?
			AND page_group = ?
		`

		if err := r.dbdget.Get(ctx).Exec(query, contextTag, roleID, labels, pageGroup).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *contextPathRepo) CreateFrontendRBAC(ctx context.Context, roleID uint, pageGroups []string) error {
	contextTag := "frontend_path"

	query := `
		INSERT INTO context_path (context_tag, role_id, path_id)
		SELECT ?, ?, id
		FROM frontend_path
		WHERE page_group IN ?
	`

	if err := r.dbdget.Get(ctx).Exec(query, contextTag, roleID, pageGroups).Error; err != nil {
		return err
	}

	return nil
}

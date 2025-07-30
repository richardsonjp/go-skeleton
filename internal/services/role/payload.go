package role

import "go-skeleton/internal/services/context_path"

type RoleGetFilterByPayload struct {
	Pagination struct {
		Page  int    `form:"page" query:"page" validate:"omitempty,numeric"`
		Limit int    `form:"limit" query:"limit" validate:"omitempty,numeric"`
		Sort  string `form:"sort" query:"sort" validate:"omitempty"`
	}
	Status string `form:"status" query:"status"`
}

type RoleCreatePayload struct {
	Name   string                                     `json:"name"`
	Status string                                     `json:"status"`
	RBAC   []map[string]context_path.ModulePermission `json:"rbac"`
}

type RoleUpdatePayload struct {
	ID     uint                                       `json:"id"`
	Name   string                                     `json:"name"`
	Status string                                     `json:"status"`
	RBAC   []map[string]context_path.ModulePermission `json:"rbac"`
}

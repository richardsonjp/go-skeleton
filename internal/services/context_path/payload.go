package context_path

type CreateRBAC struct {
	PageGroup      string   `json:"page_group"`
	EnableFrontend bool     `json:"enable_frontend"`
	BackendLabels  []string `json:"backend_labels"`
}

type ModulePermission struct {
	FrontendPath *bool             `json:"frontend_path"`
	BackendPath  []map[string]bool `json:"backend_path"`
}

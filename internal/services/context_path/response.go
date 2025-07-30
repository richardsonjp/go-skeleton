package context_path

type AuthorizedPathResponse struct {
	FrontendPath []string      `json:"frontend_path"`
	BackendPath  []BackendPath `json:"-"`
}

type BackendPath struct {
	Name   string
	Method string
}

type GroupedRBAC struct {
	PageGroup     string   `json:"page_group"`
	HasFrontend   bool     `json:"has_frontend"`
	BackendLabels []string `json:"backend_labels"`
}

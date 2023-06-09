package branch

type BranchResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	UniqueID string `json:"unique_id"`
	Status   string `json:"status"`
}

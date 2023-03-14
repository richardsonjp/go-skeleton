package admin

type MCAdminUpdatePayload struct {
	UniqueID    string    `json:"unique_id" binding:"required"`
	Name        *string   `json:"name" validate:"omitempty,min=3,excludesall=0x7C"`
	Email       *string   `json:"email" validate:"omitempty,email=true,min=8,max=50"`
	Password    *string   `json:"password"`
	PhoneNumber *string   `json:"phone_number" validate:"omitempty,min=8,max=20"`
	Role        *[]string `json:"role" validate:"omitempty,min=1,unique,dive,min=1"`
	Status      *string   `json:"status" validate:"omitempty,oneof='active' 'inactive'"`
}

type MCAdminCreatePayload struct {
	Email       string   `json:"email" binding:"required" validate:"email=true,min=8,max=50"`
	Name        string   `json:"name" binding:"required" validate:"min=3,max=100"`
	Password    string   `json:"password" binding:"required"`
	PhoneNumber string   `json:"phone_number" binding:"required" validate:"min=8,max=20"`
	Role        []string `json:"role" validate:"min=1,unique,dive,min=1"`
}

type MCGetListFilterQuery struct {
	Pagination struct {
		Page  string `form:"page" binding:"omitempty,numeric"`
		Limit string `form:"limit" binding:"omitempty,numeric"`
	}
	Search string `form:"search"`
}

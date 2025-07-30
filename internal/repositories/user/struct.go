package user

type UserList struct {
	ID          uint   `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	Phone       string `gorm:"column:phone"`
	RoleName    string `gorm:"column:role_name"`
	Status      string `gorm:"column:status"`
	LastLoginAt string `gorm:"column:last_login_at"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
}

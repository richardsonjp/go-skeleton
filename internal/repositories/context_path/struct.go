package context_path

type PathList struct {
	Name   string `gorm:"column:name"`
	Method string `gorm:"column:method"`
}

type ListRBAC struct {
	ContextPath string  `gorm:"column:context_path"`
	PageGroup   string  `gorm:"column:page_group"`
	Label       *string `gorm:"column:label"`
}

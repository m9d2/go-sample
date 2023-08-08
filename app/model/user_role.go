package model

type UserRole struct {
	User User `gorm:"foreignkey:UserID"`
	Role Role `gorm:"foreignkey:RoleID"`
}

func (UserRole *UserRole) TableName() string {
	return "user_role"
}

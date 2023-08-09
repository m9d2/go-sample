package model

type UserRole struct {
	UserID uint `json:"userId"`
	RoleID uint `json:"roleId"`
	Role   Role
}

func (UserRole *UserRole) TableName() string {
	return "user_role"
}

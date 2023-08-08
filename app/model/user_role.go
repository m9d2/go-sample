package model

type UserRole struct {
	Role Role
	User User
}

func (UserRole *UserRole) TableName() string {
	return "user_role"
}

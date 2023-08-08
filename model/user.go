package model

type User struct {
	Model
	Name        string `json:"name" gorm:"size:255"`
	DeptId      uint   `json:"dept_id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	UserType    uint8  `json:"user_type"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Sex         uint8  `json:"sex"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password"`
	Status      uint8  `json:"status"`
	LastLoginIp string `json:"last_login_ip"`
	Remark      string `json:"remark"`
	//UserRole    []UserRole
}

func (u *User) TableName() string {
	return "user"
}

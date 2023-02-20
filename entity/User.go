package entity

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (user User) TableName() string {
	return "t_user"
}

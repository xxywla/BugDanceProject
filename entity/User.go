package entity

type User struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (user User) TableName() string {
	return "t_user"
}

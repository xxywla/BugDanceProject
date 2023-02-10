package entity

type Favorite struct {
	Id      int64
	VideoId int64
	UserId  int64
}

func (favorite Favorite) TableName() string {
	return "t_favorite"
}

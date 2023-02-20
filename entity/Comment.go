package entity

type Comment struct {
	Id         int64  `json:"id"`
	VideoId    int64  `json:"video_id"`
	UserId     int64  `json:"user_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

func (comment Comment) TableName() string {
	return "t_comment"
}

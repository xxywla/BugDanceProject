package entity

type CommentVo struct {
	Id         int64  `json:"id"`
	User       UserVo `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

package entity

type Video struct {
	Id          int64  `json:"id"`
	AuthorId    int64  `json:"author_id"`
	PlayUrl     string `json:"play_url"`
	CoverUrl    string `json:"cover_url"`
	Title       string `json:"title"`
	PublishTime int64  `json:"publish_time"`
}

func (video Video) TableName() string {
	return "t_video"
}

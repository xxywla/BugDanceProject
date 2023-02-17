package entity

type Comment struct {
}

func (comment Comment) TableName() string {
	return "t_comment"
}

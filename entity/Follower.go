package entity

type Follower struct {
	Id         int64 `json:"id"`
	UserId     int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
}

func (follower Follower) TableName() string {
	return "t_follower"
}

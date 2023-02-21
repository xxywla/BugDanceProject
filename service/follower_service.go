package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

// FollowerList 获取指定用户的关注者列表
func FollowerList(userId int64) ([]*entity.Follower, error) {
	return repository.NewFollowerDaoInstance().FollowerList(userId), nil
}

// SaveFollower 把关注者信息保存到数据库
func SaveFollower(comment entity.Follower) error {
	err := repository.NewFollowerDaoInstance().AddFollower(comment.UserId, comment.FollowerId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFollower 从数据库中删除关注着信息
func DeleteFollower(follower entity.Follower) error {
	err := repository.NewFollowerDaoInstance().DeleteFollower(follower.UserId, follower.FollowerId)
	if err != nil {
		return err
	}
	return nil
}

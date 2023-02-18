package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

// FollowerList 获取指定用户发布的视频列表
func FollowerList(userId int64) ([]*entity.Follower, error) {
	return repository.NewFollowerDaoInstance().FollowerList(userId), nil
}

// SaveFollower 把视频信息保存到数据库
func SaveFollower(comment entity.Follower) error {
	err := repository.NewFollowerDaoInstance().AddFollower(comment.UserId, comment.FollowerId)
	if err != nil {
		return err
	}
	return nil
}

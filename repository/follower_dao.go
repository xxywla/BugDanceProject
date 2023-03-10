package repository

import (
	"douyinapp/entity"
	"fmt"
	"sync"
)

type FollowerDao struct {
}

var (
	followerDao  *FollowerDao
	followerOnce sync.Once
)

func NewFollowerDaoInstance() *FollowerDao {
	followerOnce.Do(func() {
		followerDao = &FollowerDao{}
	})
	return followerDao
}

// AddFollower 在关注表添加一项数据
func (*FollowerDao) AddFollower(userId int64, followerId int64) error {
	fav := &entity.Follower{UserId: userId, FollowerId: followerId}
	db.Create(fav)
	fmt.Println(fav.Id)
	return nil
}

// DeleteFollower 在关注表删除一条数据
func (*FollowerDao) DeleteFollower(userId int64, followerId int64) {
	db.Delete(entity.Follower{}, "user_id = ? and follower_id = ?", userId, followerId)

}

// FollowerList 根据用户Id获取该用户的所有关注者Id
func (*FollowerDao) FollowerList(userId int64) []*entity.Follower {

	followerList := make([]*entity.Follower, 0)

	db.Where("user_id = ?", userId).Find(&followerList)

	return followerList
}

// FollowList 根据用户Id获取该用户关注的所有人的Id
func (*FollowerDao) FollowList(followerId int64) []int64 {

	followerList := make([]*entity.Follower, 0)

	db.Where("follower_id = ?", followerId).Find(&followerList)

	followerIdList := make([]int64, 0)
	for _, follower := range followerList {
		followerIdList = append(followerIdList, follower.Id)
	}
	return followerIdList
}

// QueryVideoFollowerCount 获取指定用户的关注者数量
func (*FollowerDao) QueryVideoFollowerCount(userId int64) int64 {

	var count int64

	db.Model(&entity.Follower{}).Where("user_id = ?", userId).Count(&count)

	return count
}

// QueryFollowCount 获取指定用户关注的用户数量
func (*FollowerDao) QueryFollowCount(followerId int64) int64 {

	var count int64

	db.Model(&entity.Follower{}).Where("follower_id = ?", followerId).Count(&count)

	return count
}

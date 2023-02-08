package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (user User) TableName() string {
	return "t_user"
}

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) QueryUserInfoById(userId int64) (User, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return User{}, err
	}
	var user User
	db.First(&user, userId)
	return user, nil
}

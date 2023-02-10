package repository

import (
	"douyinapp/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

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

// QueryUserInfoById 根据用户Id查询用户信息
func (*UserDao) QueryUserInfoById(userId int64) (entity.User, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return entity.User{}, err
	}
	var user entity.User
	db.First(&user, userId)
	return user, nil
}

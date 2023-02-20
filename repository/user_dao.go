package repository

import (
	"douyinapp/entity"
	"fmt"
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
	var user entity.User
	db.First(&user, userId)
	return user, nil
}

// Yimin code

func (*UserDao) CreateUser(user *entity.User) error {
	tx := db.Begin()
	if err := tx.Create(user).Error; err != nil {
		fmt.Println("failed to create user info: ", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (*UserDao) GetUserInfoById(id int64) (*entity.User, error) {
	user := &entity.User{Id: id}
	if err := db.First(&user).Error; err != nil {
		fmt.Println("user not found: ", err)
		return nil, err
	}
	return user, nil
}

func (*UserDao) GetUserInfoByName(username string) (*entity.User, error) {
	user := &entity.User{}
	if err := db.Where("name = ?", username).First(&user).Error; err != nil {
		fmt.Println("user not found: ", err)
		return nil, err
	}
	return user, nil
}

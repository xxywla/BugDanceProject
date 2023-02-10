package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

func UserInfo(userId int64) (entity.User, error) {
	return repository.NewUserDaoInstance().QueryUserInfoById(userId)
}

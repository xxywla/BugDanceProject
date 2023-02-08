package service

import (
	"douyinapp/repository"
)

func UserInfo(user_id int64) (repository.User, error) {
	return repository.NewUserDaoInstance().QueryUserInfoById(user_id)
}

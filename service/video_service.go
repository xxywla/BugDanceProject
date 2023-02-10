package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

func PublishList(userId int64, token string) ([]*entity.Video, error) {
	return repository.NewVideoDaoInstance().QueryVideoListById(userId)
}

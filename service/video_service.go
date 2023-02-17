package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

// PublishList 获取指定用户发布的视频列表
func PublishList(userId int64, token string) ([]*entity.Video, error) {
	return repository.NewVideoDaoInstance().QueryVideoListById(userId)
}

// SaveVideo 把视频信息保存到数据库
func SaveVideo(video entity.Video) error {
	err := repository.NewVideoDaoInstance().SaveVideo(video)
	if err != nil {
		return err
	}
	return nil
}

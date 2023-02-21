package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
	"fmt"
	"math"
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

// Yimin code

func Feed(latestTime int64) ([]entity.VideoVo, int64, error) {
	videos, err := repository.NewVideoDaoInstance().GetVidoes(latestTime)
	if err != nil || videos == nil || len(videos) == 0 {
		return nil, 0, fmt.Errorf("no videos found")
	}
	videoVos := make([]entity.VideoVo, 0)
	var nextTime int64 = math.MaxInt64
	for _, video := range videos {
		videoVo := video2Vo(&video)
		if videoVo != nil {
			videoVos = append(videoVos, *videoVo)
			if video.PublishTime < nextTime {
				nextTime = video.PublishTime
			}
		}
	}
	return videoVos, nextTime, nil
}

func video2Vo(video *entity.Video) *entity.VideoVo {
	userVo, err := GetUserInfoById(video.AuthorId)
	if err != nil || userVo == nil {
		fmt.Print(err)
		return nil
	}
	v := entity.VideoVo{
		Id:            video.Id,
		Author:        *userVo,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         video.Title,
	}
	return &v
}

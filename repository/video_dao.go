package repository

import (
	"douyinapp/entity"
	"sync"
)

type VideoDao struct {
}

var (
	videoDao  *VideoDao
	videoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

// QueryVideoListById 根据用户Id获取用户上传的视频列表
func (*VideoDao) QueryVideoListById(userId int64) []int64 {
	videoList := make([]*entity.Video, 0)
	db.Where("author_id = ?", userId).Find(&videoList)
	videoIdList := make([]int64, 0)
	for _, video := range videoList {
		videoIdList = append(videoIdList, video.Id)
	}
	return videoIdList
}

// QueryVideoById 根据视频Id获取视频信息
func (*VideoDao) QueryVideoById(videoId int64) entity.Video {
	var video entity.Video
	db.Where("id = ?", videoId).Find(&video)
	return video
}

// SaveVideo 把视频信息保存到数据库
func (*VideoDao) SaveVideo(video entity.Video) error {
	db.Create(&video)

	return nil
}

// Yimin code

func (*VideoDao) GetVidoes(latestTime int64) ([]entity.Video, error) {
	video := make([]entity.Video, 0)
	if err := db.Where("publish_time < ?", latestTime).Order("publish_time desc").Limit(30).Find(&video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

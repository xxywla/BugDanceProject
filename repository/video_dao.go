package repository

import (
	"douyinapp/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func (*VideoDao) QueryVideoListById(userId int64) ([]*entity.Video, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"
	videoList := make([]*entity.Video, 0)

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Where("author_id = ?", userId).Find(&videoList)
	return videoList, nil
}

// QueryVideoById 根据视频Id获取视频信息
func (*VideoDao) QueryVideoById(videoId int64) entity.Video {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var video entity.Video
	db.Where("id = ?", videoId).Find(&video)
	return video
}

// SaveVideo 把视频信息保存到数据库
func (*VideoDao) SaveVideo(video entity.Video) error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Create(&video)

	return nil
}

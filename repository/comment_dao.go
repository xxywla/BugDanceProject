package repository

import (
	"douyinapp/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type CommentDao struct {
}

var (
	commentDao  *CommentDao
	commentOnce sync.Once
)

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(func() {
		commentDao = &CommentDao{}
	})
	return commentDao
}

// QueryVideoCommentCount 获取指定视频的评论数量
func (*CommentDao) QueryVideoCommentCount(videoId int64) int64 {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var count int64

	db.Model(&entity.Comment{}).Where("video_id = ?", videoId).Count(&count)

	return count
}

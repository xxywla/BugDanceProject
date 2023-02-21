package repository

import (
	"douyinapp/entity"
	"fmt"
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

// AddComment 评论表添加一项数据
func (*CommentDao) AddComment(videoId int64, userId int64, content string, creatDate string) error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fav := &entity.Comment{VideoId: videoId, UserId: userId, Content: content, CreateDate: creatDate}
	db.Create(fav)
	fmt.Println(fav.Id)
	return nil
}

// DeleteComment 在评论表删除一条数据
func (*CommentDao) DeleteComment(commentId int64) error {

	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Delete(entity.Comment{}, "comment_id = ?", commentId)

	return nil
}

// CommentList 根据用户Id获取该用户评论的所有视频Id
func (*CommentDao) CommentList(userId int64) []*entity.Comment {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	commentList := make([]*entity.Comment, 0)

	db.Where("user_id = ? ORDER BY 'create_date'", userId).Find(&commentList)

	return commentList
}

// QueryVideoCommentCount 获取指定视频的评论数
func (*CommentDao) QueryVideoCommentCount(videoId int64) int64 {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var count int64

	db.Model(&entity.Comment{}).Where("video_id = ?", videoId).Count(&count)

	return count
}

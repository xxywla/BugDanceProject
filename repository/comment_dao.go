package repository

import (
	"douyinapp/entity"
	"fmt"
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
	var count int64
	db.Model(&entity.Comment{}).Where("video_id = ?", videoId).Count(&count)
	return count
}

// AddComment 评论表添加一项数据
func (*CommentDao) AddComment(videoId int64, userId int64, content string, creatDate string) error {
	fav := &entity.Comment{VideoId: videoId, UserId: userId, Content: content, CreateDate: creatDate}
	db.Create(fav)
	fmt.Println(fav.Id)
	return nil
}

// DeleteComment 在评论表删除一条数据
func (*CommentDao) DeleteComment(videoId int64, userId int64, content string, creatDate string) {
	db.Delete(entity.Comment{}, "video_id = ? and user_id = ? and content = ? and create_date = ?", videoId, userId, content, creatDate)
}

// CommentList 根据用户Id获取该用户评论的所有视频Id
func (*CommentDao) CommentList(userId int64) []*entity.Comment {
	commentList := make([]*entity.Comment, 0)
	db.Where("user_id = ?", userId).Find(&commentList)
	return commentList
}

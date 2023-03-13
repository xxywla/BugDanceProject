package repository

import (
	"douyinapp/entity"
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
func (*CommentDao) AddComment(comment *entity.Comment) error {
	err := db.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteComment 在评论表删除一条数据
func (*CommentDao) DeleteComment(commentId int64) error {
	err := db.Delete(entity.Comment{}, "id = ?", commentId).Error
	return err
}

// 根据video id获取该视频的所有评论，按发布时间倒序
func (*CommentDao) CommentList(videoId int64) ([]entity.Comment, error) {
	commentList := make([]entity.Comment, 0)
	if err := db.Order("create_date desc").Where("video_id = ?", videoId).Find(&commentList).Error; err != nil {
		return nil, err
	}
	return commentList, nil
}

package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

// CommentList 获取指定视频的评论
func CommentList(userId int64) ([]*entity.Comment, error) {
	return repository.NewCommentDaoInstance().CommentList(userId), nil
}

// SaveComment 把评论保存到数据库
func SaveComment(comment entity.Comment) error {
	err := repository.NewCommentDaoInstance().AddComment(comment.VideoId, comment.UserId, comment.Content, comment.CreateDate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteComment 从数据库删除评论
func DeleteComment(commentId int64) error {
	err := repository.NewCommentDaoInstance().DeleteComment(commentId)
	if err != nil {
		return err
	}
	return nil
}

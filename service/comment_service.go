package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

// CommentList 获取指定用户发布的视频列表
func CommentList(userId int64) ([]*entity.Comment, error) {
	return repository.NewCommentDaoInstance().CommentList(userId), nil
}

// SaveComment 把视频信息保存到数据库
func SaveComment(comment entity.Comment) error {
	err := repository.NewCommentDaoInstance().AddComment(comment.VideoId, comment.UserId, comment.Content, comment.CreateDate)
	if err != nil {
		return err
	}
	return nil
}

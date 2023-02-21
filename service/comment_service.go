package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
	"fmt"
	"time"
)

// CommentList 获取指定用户发布的视频列表
func CommentList(userId int64) ([]*entity.Comment, error) {
	return repository.NewCommentDaoInstance().CommentList(userId), nil
}

// 保存评论
func SaveComment(videoId int64, user entity.User, commentText string) (*entity.CommentVo, error) {
	createDate := time.Now().Format("01-02")
	comment := &entity.Comment{VideoId: videoId, UserId: user.Id, Content: commentText, CreateDate: createDate}

	comment, err := repository.NewCommentDaoInstance().AddComment(comment)
	commentVo := entity.CommentVo{Id: comment.Id}
	if err != nil {
		fmt.Printf("保存评论失败: %s", err)
		return nil, err
	}
	return &commentVo, nil
}

// 删除评论
func DeleteComment(commentId int64) error {
	err := repository.NewCommentDaoInstance().DeleteComment(commentId)
	return err
}

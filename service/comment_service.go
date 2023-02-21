package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
	"fmt"
	"time"
)

// 获取当前视频的评论列表
func CommentList(videoId int64) ([]entity.CommentVo, error) {
	commentList, err := repository.NewCommentDaoInstance().CommentList(videoId)
	if err != nil {
		return nil, err
	}
	commentVoList := make([]entity.CommentVo, 0)
	for _, comment := range commentList {
		commentVo := comment2Vo(&comment)
		if commentVo != nil {
			commentVoList = append(commentVoList, *commentVo)
		}
	}
	return commentVoList, nil
}

// 保存评论
func SaveComment(videoId int64, user entity.User, commentText string) (*entity.CommentVo, error) {
	createDate := time.Now().Format("01-02")
	comment := &entity.Comment{VideoId: videoId, UserId: user.Id, Content: commentText, CreateDate: createDate}

	comment, err := repository.NewCommentDaoInstance().AddComment(comment)
	userVo := user2Vo(&user)
	commentVo := entity.CommentVo{Id: comment.Id, User: *userVo, Content: commentText, CreateDate: createDate}
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

func comment2Vo(comment *entity.Comment) *entity.CommentVo {
	userVo, err := GetUserInfoById(comment.UserId)
	if err != nil || userVo == nil {
		fmt.Print(err)
		return nil
	}
	commentVo := entity.CommentVo{
		Id:         comment.Id,
		User:       *userVo,
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
	}
	return &commentVo
}

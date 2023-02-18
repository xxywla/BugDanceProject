package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	StatusCode  int32             `json:"status_code"`
	StatusMsg   string            `json:"status_msg"`
	CommentList []*entity.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	StatusCode int32          `json:"status_code"`
	StatusMsg  string         `json:"status_msg"`
	Comment    entity.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 0,
				StatusMsg: "评论成功",
				Comment: entity.Comment{
					Id:         1,
					UserId:     user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(videoId int64) *CommentListResponse {
	commentList, _ := service.CommentList(videoId)
	return &CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "查询成功",
		CommentList: commentList,
	}
}

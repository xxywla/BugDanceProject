package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
			comment := entity.Comment{
				Id:         1,
				UserId:     user,
				Content:    text,
				CreateDate: "05-01",
			}
			service.SaveComment(comment)
			c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 0,
				StatusMsg: "评论成功",
				Comment:   comment,
			})
			return
		} else if actionType == "2" {
			comId := c.Query("comment_id")
			commentId, err := strconv.ParseInt(comId, 10, 64)
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "发生错误"})
				return
			}
			err = service.DeleteComment(commentId)
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "发生错误"})
				return
			}
			c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 0,
				StatusMsg: "删除成功"})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "操作类型不存在"})
	} else {
		c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		vId := c.Query("video_id")
		vedioId, err := strconv.ParseInt(vId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, CommentListResponse{StatusCode: 1, StatusMsg: "发生错误"})
			return
		}
		commentList, err := service.CommentList(vedioId)
		if err != nil {
			c.JSON(http.StatusOK, CommentListResponse{StatusCode: 1, StatusMsg: "发生错误"})
			return
		}
		c.JSON(http.StatusOK, CommentListResponse{StatusCode: 0,
			StatusMsg:   "查询成功",
			CommentList: commentList})
		return
	} else {
		c.JSON(http.StatusOK, CommentListResponse{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}

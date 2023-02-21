package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	StatusCode  int32              `json:"status_code"`
	StatusMsg   string             `json:"status_msg"`
	CommentList []entity.CommentVo `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	StatusCode int32            `json:"status_code"`
	StatusMsg  string           `json:"status_msg"`
	Comment    entity.CommentVo `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	session := sessions.Default(c)
	user := session.Get(token)

	if user == nil {
		c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "未登录或登陆过期，请先登录"})
		return
	}

	actionType := c.Query("action_type")
	fmt.Println(actionType)
	if actionType == "1" {
		videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
		commentText := c.Query("comment_text")
		commentVo, err := service.SaveComment(videoId, user.(entity.User), commentText)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "评论失败"})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 0,
			StatusMsg: "评论成功",
			Comment:   *commentVo})
		return
	} else {
		if actionType == "2" {
			commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			err := service.DeleteComment(commentId)
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "删除评论失败"})
				return
			} else {
				c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 0, StatusMsg: "删除评论成功"})
				return
			}
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{StatusCode: 1, StatusMsg: "action_type不合法"})
		}
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentList, err := service.CommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			StatusCode: 1,
			StatusMsg:  "无法获取评论",
		})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			StatusCode:  0,
			CommentList: commentList,
		})
	}
}

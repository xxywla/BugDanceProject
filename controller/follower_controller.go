package controller

import (
	"douyinapp/entity"
	"douyinapp/service"

	"github.com/gin-gonic/gin"
)

type FollowerListResponse struct {
	StatusCode   int32              `json:"status_code"`
	StatusMsg    string             `json:"status_msg"`
	FollowerList []*entity.Follower `json:"comment_list,omitempty"`
}

type FollowerActionResponse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	Follower   entity.Follower `json:"comment,omitempty"`
}

// FollowerAction no practical effect, just check if token is valid
func FollowerAction(c *gin.Context) {
	// token := c.Query("token")
	// actionType := c.Query("action_type")

	// if user, exist := usersLoginInfo[token]; exist {
	// 	if actionType == "1" {
	// 		c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 0,
	// 			StatusMsg: "评论成功",
	// 			Follower: entity.Follower{
	// 				Id:         1,
	// 				UserId:     user,
	// 				FollowerId: 10,
	// 			}})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 0})
	// } else {
	// 	c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 1, StatusMsg: "用户不存在"})
	// }
}

// FollowerList all videos have same demo comment list
func FollowerList(videoId int64) *FollowerListResponse {
	commentList, _ := service.FollowerList(videoId)
	return &FollowerListResponse{
		StatusCode:   0,
		StatusMsg:    "查询成功",
		FollowerList: commentList,
	}
}

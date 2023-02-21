package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			follower := entity.Follower{
				Id:         1,
				UserId:     user,
				FollowerId: 10,
			}
			err := service.SaveFollower(follower)
			if err != nil {
				c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 1, StatusMsg: "发生错误"})
			}
			c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 0,
				StatusMsg: "关注成功"})
			return
		} else if actionType == "2" {
			follower := entity.Follower{
				Id:         1,
				UserId:     user,
				FollowerId: 10,
			}
			err := service.DeleteFollower(follower)
			if err != nil {
				c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 1, StatusMsg: "发生错误"})
			}
			c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 0,
				StatusMsg: "已取消关注"})
			return
		}
		c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 1, StatusMsg: "发生错误"})
	} else {
		c.JSON(http.StatusOK, FollowerActionResponse{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}

// FollowerList all videos have same demo comment list
func FollowerList(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		uId := c.Query("video_id")
		userId, err := strconv.ParseInt(uId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, FollowerListResponse{StatusCode: 1, StatusMsg: "发生错误"})
			return
		}
		followerList, err := service.FollowerList(userId)
		if err != nil {
			c.JSON(http.StatusOK, FollowerListResponse{StatusCode: 1, StatusMsg: "发生错误"})
			return
		}
		c.JSON(http.StatusOK, FollowerListResponse{StatusCode: 0,
			StatusMsg:    "查询成功",
			FollowerList: followerList})
		return
	} else {
		c.JSON(http.StatusOK, FollowerListResponse{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}

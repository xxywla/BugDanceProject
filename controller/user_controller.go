package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type UserResponse struct {
//	StatusCode int32       `json:"status_code"`
//	StatusMsg  string      `json:"status_msg"`
//	User       entity.User `json:"user"`
//}

// UserInfo 根据用户Id获取用户信息
//func UserInfo(userId int64) *UserResponse {
//	user, err := service.UserInfo(userId)
//	if err != nil {
//		return &UserResponse{StatusCode: -1,
//			StatusMsg: "没找到用户",
//			User:      user}
//	}
//	return &UserResponse{
//		StatusCode: 0,
//		StatusMsg:  "查询成功",
//		User:       user}
//}

//type FavoriteActionResponse struct {
//	StatusCode int32  `json:"status_code"`
//	StatusMsg  string `json:"status_msg"`
//}

// FavoriteAction 赞操作
//func FavoriteAction(userId int64, videoId int64, actionType int32) FavoriteActionResponse {
//
//	service.FavoriteAction(userId, videoId, actionType)
//	return FavoriteActionResponse{0, "成功操作"}
//}

type FavoriteListResponse struct {
	StatusCode int32            `json:"status_code"`
	StatusMsg  string           `json:"status_msg"`
	VideoList  []entity.VideoVo `json:"video_list"`
}

// FavoriteList 获取指定用户所有喜欢的视频
//func FavoriteList(userId int64) FavoriteListResponse {
//	videoList := service.FavoriteListByUserId(userId)
//	return FavoriteListResponse{
//		StatusCode: 0,
//		StatusMsg:  "查找成功",
//		VideoList:  videoList,
//	}
//}

// Yimin code

type UserResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	User       entity.User `json:"user"`
}

type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := service.EncryptPassword(username + password)
	if err != nil || token == "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to create token",
		})
	} else {
		if _, err := service.GetUserInfoByName(username); err != nil {
			if user, err := service.CreateUser(username, password); err != nil || user == nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					StatusCode: 1,
					StatusMsg:  "Failed to create account",
				})
			} else {
				session := sessions.Default(c)
				session.Options(sessions.Options{
					Path:   "/",
					MaxAge: int(3600),
				})
				session.Set(token, 1)
				session.Save()
				c.JSON(http.StatusOK, UserLoginResponse{
					StatusCode: 0,
					UserId:     user.Id,
					Token:      token,
				})
			}
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			})
		}
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := service.EncryptPassword(username + password)
	if err != nil || token == "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to create token",
		})
	} else {
		session := sessions.Default(c)
		session.Options(sessions.Options{
			Path:   "/",
			MaxAge: int(3600),
		})
		session.Set(token, 1)
		session.Save()

		if err := service.VerifyAccount(username, password); err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				StatusCode: 1,
				StatusMsg:  "Wrong username or password",
			})
		} else {
			user, err := service.GetUserInfoByName(username)
			if err != nil || user == nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					StatusCode: 1,
					StatusMsg:  "Failed to login",
				})
			} else {
				c.JSON(http.StatusOK, UserLoginResponse{
					StatusCode: 0,
					UserId:     user.Id,
					Token:      token,
				})
			}
		}
	}
}

func UserInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	session := sessions.Default(c)
	// session.Options(sessions.Options{
	// 	Path:   "/",
	// 	MaxAge: int(3600),
	// })
	session.Set(token, 1)
	session.Save()
	flag := session.Get(token)
	if err != nil || flag == nil {
		c.JSON(http.StatusOK, UserResponse{
			StatusCode: 1,
			StatusMsg:  "Cannot parse user_id or token",
		})
	} else {
		if user, err := service.GetUserInfoById(id); err != nil || user == nil {
			c.JSON(http.StatusOK, UserResponse{
				StatusCode: 1,
				StatusMsg:  "User doesn't exist",
			})
		} else {
			c.JSON(http.StatusOK, UserResponse{
				StatusCode: 0,
				User:       *user,
			})
		}
	}

}

type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// FavoriteAction 赞操作
func FavoriteAction(userId int64, videoId int64, actionType int32) FavoriteActionResponse {

	service.FavoriteAction(userId, videoId, actionType)
	return FavoriteActionResponse{0, "成功操作"}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		VideoList:  nil,
	})
}

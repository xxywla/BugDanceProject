package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	User       entity.UserVo `json:"user"`
}

type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   int64  `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(username string, password string, userId int64) (string, error) {
	expireTime := time.Now().Add(time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("setting.JwtSecret"))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("setting.JwtSecret"), nil
	})
	if err != nil || tokenClaims == nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token校验失败")
	}
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if _, err := service.GetUserInfoByName(username); err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "User already exist",
		})
		return
	}

	if userVo, err := service.CreateUser(username, password); err != nil || userVo == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to create account",
		})
	} else {
		token, err := GenerateToken(username, password, userVo.Id)
		if err != nil || token == "" {
			c.JSON(http.StatusOK, UserLoginResponse{
				StatusCode: 1,
				StatusMsg:  "Failed to create token",
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 0,
			UserId:     userVo.Id,
			Token:      token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if err := service.VerifyAccount(username, password); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "Wrong username or password",
		})
		return
	}

	if userVo, err := service.GetUserInfoByName(username); err != nil || userVo == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "Failed to login",
		})
	} else {
		token, err := GenerateToken(username, password, userVo.Id)
		if err != nil || token == "" {
			c.JSON(http.StatusOK, UserLoginResponse{
				StatusCode: 1,
				StatusMsg:  "Failed to create token",
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			StatusCode: 0,
			UserId:     userVo.Id,
			Token:      token,
		})
	}
}

func UserInfo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	tokenClaims, err := ParseToken(token)
	if err != nil || tokenClaims == nil {
		c.JSON(http.StatusOK, UserResponse{
			StatusCode: 1,
			StatusMsg:  "Cannot parse user_id or token",
		})
		return
	}

	if userVo, err := service.GetUserInfoById(id); err != nil || userVo == nil {
		c.JSON(http.StatusOK, UserResponse{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			StatusCode: 0,
			User:       *userVo,
		})
	}
}

type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// FavoriteAction 赞操作
func FavoriteAction(c *gin.Context) {

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	token := c.Query("token")

	tokenClaims, err := ParseToken(token)
	if err != nil || tokenClaims == nil {
		c.JSON(http.StatusOK, FavoriteActionResponse{1, "操作失败"})
		return
	}

	userId := tokenClaims.UserId

	service.FavoriteAction(userId, videoId, int32(actionType))
	c.JSON(200, FavoriteActionResponse{0, "成功操作"})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		VideoList:  nil,
	})
}

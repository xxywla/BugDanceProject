package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
)

type UserResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	User       entity.User `json:"user"`
}

// UserInfo 根据用户Id获取用户信息
func UserInfo(userId int64) *UserResponse {
	user, err := service.UserInfo(userId)
	if err != nil {
		return &UserResponse{StatusCode: -1,
			StatusMsg: "没找到用户",
			User:      user}
	}
	return &UserResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		User:       user}
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

type FavoriteListResponse struct {
	StatusCode int32            `json:"status_code"`
	StatusMsg  string           `json:"status_msg"`
	VideoList  []entity.VideoVo `json:"video_list"`
}

// FavoriteList 获取指定用户所有喜欢的视频
func FavoriteList(userId int64) FavoriteListResponse {
	videoList := service.FavoriteListByUserId(userId)
	return FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "查找成功",
		VideoList:  videoList,
	}
}

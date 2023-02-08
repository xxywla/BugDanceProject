package controller

import (
	"douyinapp/repository"
	"douyinapp/service"
)

type UserRequest struct {
	UserId int64
	Token  string
}
type UserResponse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	User       repository.User `json:"user"`
}

func UserInfo(request UserRequest) *UserResponse {
	user, err := service.UserInfo(request.UserId)
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

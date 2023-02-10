package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
)

type PublishListResponse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	VideoList  []*entity.Video `json:"video_list"`
}

func PublishList(userId int64, token string) *PublishListResponse {
	publishList, _ := service.PublishList(userId, token)
	return &PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		VideoList:  publishList,
	}
}

package controller

import (
	"douyinapp/entity"
	"douyinapp/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

type PublishActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// PublishAction 上传视频
func PublishAction(c *gin.Context) {

	session := sessions.Default(c)

	file, err := c.FormFile("data")

	if err != nil {
		c.JSON(500, &PublishActionResponse{StatusCode: 1, StatusMsg: "视频有问题"})
		return
	}

	// 把文件写入对应位置
	err = c.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		c.JSON(500, &PublishActionResponse{StatusCode: 1, StatusMsg: "视频保存有问题"})
		return
	}

	// 调用 service 把视频信息保存到数据库
	var video entity.Video
	video.PlayUrl = "./" + file.Filename
	video.Title = c.PostForm("title")
	video.CoverUrl = "www.picture.com"

	token := c.PostForm("token")

	tmp := session.Get(token)
	fmt.Println(tmp)
	video.AuthorId = int64(session.Get(token).(int))

	err = service.SaveVideo(video)
	if err != nil {
		c.JSON(500, &PublishActionResponse{StatusCode: 1, StatusMsg: "视频保存有问题"})
		return
	}

	c.JSON(200, &PublishActionResponse{StatusCode: 0, StatusMsg: "视频上传成功"})
}

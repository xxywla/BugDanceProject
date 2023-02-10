package main

import (
	"douyinapp/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()

	// 用户信息
	r.GET("/douyin/user", func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
		if err != nil {

		}
		data := controller.UserInfo(userId)
		c.JSON(200, data)
	})

	// 发布列表
	r.GET("/douyin/publish/list", func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
		token := c.Query("token")
		if err != nil {

		}
		data := controller.PublishList(userId, token)
		c.JSON(200, data)
	})

	// 赞操作
	r.POST("/douyin/favorite/action", func(c *gin.Context) {
		videoId, _ := strconv.ParseInt(c.PostForm("video_id"), 10, 64)
		actionType, _ := strconv.ParseInt(c.PostForm("action_type"), 10, 32)
		token := c.PostForm("token")

		session := sessions.Default(c)

		userId := session.Get(token)

		if userId == nil {
			return
		}

		data := controller.FavoriteAction(userId.(int64), videoId, int32(actionType))
		c.JSON(200, data)
	})

	// 喜欢列表

	err := r.Run()
	if err != nil {
		return
	}
}

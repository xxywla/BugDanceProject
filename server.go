package main

import (
	"douyinapp/controller"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("shuiche"))
	r.Use(sessions.Sessions("mysession", store))

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

		session := sessions.Default(c)
		session.Set("111", 3)
		err = session.Save()
		if err != nil {
			return
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
	r.GET("/douyin/favorite/list", func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
		if err != nil {
		}

		token := c.Query("token")
		session := sessions.Default(c)
		tmp := session.Get(token)
		fmt.Println(tmp)

		data := controller.FavoriteList(userId)
		c.JSON(200, data)
	})

	// 视频投稿
	r.POST("/douyin/publish/action", func(c *gin.Context) {
		controller.PublishAction(c)
	})

	err := r.Run()
	if err != nil {
		return
	}
}

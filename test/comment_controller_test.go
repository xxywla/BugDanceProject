package test

import (
	"net/http"
	"testing"
)

func TestCommentAction(t *testing.T) {
	e := newExpect(t)

	loginResp := e.POST("/douyin/user/login/").
		WithQuery("username", "cym001").WithQuery("password", "123456").
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	commentResp := e.POST("/douyin/comment/action/").
		WithQuery("token", loginResp.Value("token").String().Raw()).
		WithQuery("video_id", 1).
		WithQuery("action_type", 1).
		WithQuery("comment_text", "test").
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	commentResp.Value("status_code").Number().Equal(0)
	commentResp.Value("comment").Object().NotEmpty()
}

func TestCommentList(t *testing.T) {
	e := newExpect(t)

	commentListResp := e.GET("/douyin/comment/list/").
		WithQuery("video_id", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	commentListResp.Value("status_code").Number().Equal(0)
}

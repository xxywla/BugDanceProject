package test

import (
	"net/http"
	"testing"
)

func TestFeeding(t *testing.T) {
	e := newExpect(t)

	feedResp := e.GET("/douyin/feed/").
		WithQuery("last_time", 0).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	feedResp.Value("status_code").Number().IsEqual(0)
}

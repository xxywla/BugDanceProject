package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
	"fmt"
	"math"
)

// PublishList 获取指定用户发布的视频列表
func PublishList(userId int64) []entity.VideoVo {

	videoDao := repository.NewVideoDaoInstance()

	favoriteDao := repository.NewFavoriteDaoInstance()

	commentDao := repository.NewCommentDaoInstance()

	videoIdList := videoDao.QueryVideoListById(userId)

	// 根据视频Id查询对应的视频信息
	videoVoList := make([]entity.VideoVo, 0)

	for _, videoId := range videoIdList {

		video := videoDao.QueryVideoById(videoId)

		var videoVo entity.VideoVo

		videoVo.Id = video.Id
		videoVo.Author.Id = video.AuthorId
		videoVo.PlayUrl = video.PlayUrl
		videoVo.CoverUrl = video.CoverUrl
		videoVo.Title = video.Title

		// 获取视频作者的信息

		// 获取视频的点赞数
		videoVo.FavoriteCount = favoriteDao.QueryVideoFavoriteCount(videoId)

		// 获取视频的评论数
		videoVo.CommentCount = commentDao.QueryVideoCommentCount(videoId)

		// 当前用户一定已经点赞了
		videoVo.IsFavorite = true

		videoVoList = append(videoVoList, videoVo)
	}

	return videoVoList
}

// SaveVideo 把视频信息保存到数据库
func SaveVideo(video entity.Video) error {
	err := repository.NewVideoDaoInstance().SaveVideo(video)
	if err != nil {
		return err
	}
	return nil
}

// Yimin code

func Feed(latestTime int64) ([]entity.VideoVo, int64, error) {
	videos, err := repository.NewVideoDaoInstance().GetVidoes(latestTime)
	if err != nil || videos == nil || len(videos) == 0 {
		return nil, 0, fmt.Errorf("no videos found")
	}
	videoVos := make([]entity.VideoVo, 0)
	var nextTime int64 = math.MaxInt64
	for _, video := range videos {
		videoVo := video2Vo(&video)
		if videoVo != nil {
			videoVos = append(videoVos, *videoVo)
			if video.PublishTime < nextTime {
				nextTime = video.PublishTime
			}
		}
	}
	return videoVos, nextTime, nil
}

func video2Vo(video *entity.Video) *entity.VideoVo {
	userVo, err := GetUserInfoById(video.AuthorId)
	if err != nil || userVo == nil {
		fmt.Print(err)
		return nil
	}
	v := entity.VideoVo{
		Id:            video.Id,
		Author:        *userVo,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         video.Title,
	}
	return &v
}

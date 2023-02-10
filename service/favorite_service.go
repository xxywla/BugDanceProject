package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
)

// FavoriteAction 处理用户点赞或取消赞的操作
func FavoriteAction(userId int64, videoId int64, actionType int32) {

	favoriteDao := repository.NewFavoriteDaoInstance()

	if actionType == 1 {
		// 点赞
		favoriteDao.AddFavorite(userId, videoId)
	} else if actionType == 2 {
		// 取消点赞
		favoriteDao.DeleteFavorite(userId, videoId)
	}

}

// FavoriteListByUserId 获取指定用户Id的所有喜欢的视频信息
func FavoriteListByUserId(userId int64) []entity.VideoVo {

	favoriteDao := repository.NewFavoriteDaoInstance()

	//userDao := repository.NewUserDaoInstance()

	videoDao := repository.NewVideoDaoInstance()

	// 获取用户喜欢的视频Id列表
	videoIdList := favoriteDao.FavoriteList(userId)

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

		// 获取视频的评论数

		// 当前用户一定已经点赞了
		videoVo.IsFavorite = true

		videoVoList = append(videoVoList)
	}

	return videoVoList

}

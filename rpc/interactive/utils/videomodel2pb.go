package utils

import (
	"douyin/model/video"
	"douyin/rpc/interactive/pb"
)

func VideoModelPb(mvideo *video.Video, author *pb.User) (pvideo *pb.Video, err error) {
	var isFavorite bool
	if mvideo.IsFavorite == "true" {
		isFavorite = true
	} else {
		isFavorite = false
	}
	pvideo = &pb.Video{
		Id:            mvideo.Id,
		Author:        author,
		PlayUrl:       mvideo.PlayUrl,
		CoverUrl:      mvideo.CoverUrl,
		FavoriteCount: mvideo.FavoriteCount,
		CommentCount:  mvideo.CommentCount,
		IsFavorite:    isFavorite,
		Title:         mvideo.Title,
	}
	return
}

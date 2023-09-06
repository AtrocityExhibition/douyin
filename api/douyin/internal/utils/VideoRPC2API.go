package utils

import (
	"douyin/api/douyin/internal/types"
	"douyin/rpc/interactive/pb"
)

func VideoRPC2API(pvideo *pb.Video) (video *types.Video) {
	user := UserRPC2API(pvideo.Author)
	return &types.Video{
		Id:             pvideo.Id,
		Author:         user,
		Play_url:       pvideo.PlayUrl,
		Cover_url:      pvideo.CoverUrl,
		Favorite_count: pvideo.FavoriteCount,
		Comment_count:  pvideo.CommentCount,
		Is_favorite:    pvideo.IsFavorite,
		Title:          pvideo.Title,
	}
}

package utils

import (
	"douyin/api/douyin/internal/types"
	"douyin/rpc/interactive/pb"
)

func UserRPC2API(puser *pb.User) (user *types.User) {
	return &types.User{
		Id:               puser.Id,
		Name:             puser.Name,
		Follow_count:     puser.FollowCount,
		Follower_count:   puser.FollowerCount,
		Is_follow:        puser.IsFollow,
		Avatar:           puser.Avatar,
		Background_image: puser.BackgroundImage,
		Signature:        puser.Signature,
		Total_favorited:  puser.TotalFavorited,
		Work_count:       puser.WorkCount,
		Favorite_count:   puser.FavoriteCount,
	}
}

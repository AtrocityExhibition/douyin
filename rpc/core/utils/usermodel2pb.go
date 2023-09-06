package utils

import (
	"douyin/model/user"
	"douyin/rpc/core/pb"
)

func UserModelPb(muser *user.User) (puser *pb.User, err error) {
	var isFollow bool
	if muser.IsFollow == "true" {
		isFollow = true
	} else {
		isFollow = false
	}
	puser = &pb.User{
		Id:              muser.Id,
		Name:            muser.Name,
		FollowCount:     muser.FollowCount,
		FollowerCount:   muser.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          muser.Avatar.String,
		BackgroundImage: muser.BackgroundImage.String,
		Signature:       muser.Signature.String,
		TotalFavorited:  muser.TotalFavorited,
		WorkCount:       muser.WorkCount,
		FavoriteCount:   muser.FavoriteCount,
	}
	return
}

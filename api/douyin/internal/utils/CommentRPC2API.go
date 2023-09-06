package utils

import (
	"douyin/api/douyin/internal/types"
	"douyin/rpc/interactive/pb"
)

func CommentRPC2API(comment *pb.Comment) types.Comment {
	user := UserRPC2API(comment.User)
	return types.Comment{
		Id:          comment.Id,
		User:        *user,
		Content:     comment.Content,
		Create_date: comment.CreateDate,
	}
}

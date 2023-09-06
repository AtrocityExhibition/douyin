package logic

import (
	"context"
	"database/sql"
	"douyin/model/favorite"
	"douyin/rpc/interactive/internal/JWT"
	"errors"
	"fmt"

	"douyin/rpc/interactive/internal/svc"
	"douyin/rpc/interactive/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinFavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinFavoriteActionLogic {
	return &DouyinFavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinFavoriteActionLogic) DouyinFavoriteAction(in *pb.DouyinFavoriteActionRequest) (*pb.DouyinFavoriteActionResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入 rpc DouyinFavoriteAction")
	if err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token); err != nil {
		fmt.Println("token鉴权出错了")
		return &pb.DouyinFavoriteActionResponse{
			StatusCode: 1,
			StatusMsg:  "token鉴权出错了",
		}, err
	}
	// 从token中解析出用户名
	claims, err := JWT.JWTAuth(in.Token)
	username := claims["Username"].(string)
	password := claims["Password"].(string)
	user, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, username, password)
	if err != nil {
		return &pb.DouyinFavoriteActionResponse{
			StatusCode: 1,
			StatusMsg:  "根据用户名和密码查找用户出错了",
		}, err
	}
	fmt.Println("action_type: ", in.ActionType)
	if in.ActionType == 1 {
		// 1.点赞 将 video的favorite_count+1
		video, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
		if err != nil {
			fmt.Println("查询视频失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "查询视频失败",
			}, err
		}
		video.FavoriteCount++
		// 更新
		err = l.svcCtx.VideoModel.Update(l.ctx, video)
		if err != nil {
			fmt.Println("视频喜欢数更新失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "视频喜欢数更新失败",
			}, err
		}

		// 2.将user_id和video_id加入favorite中
		data := favorite.Favorite{
			UserId:  sql.NullInt64{user.Id, true},
			VideoId: sql.NullInt64{video.Id, true},
		}
		_, err = l.svcCtx.FavoriteModel.Insert(l.ctx, &data)
		if err != nil {
			fmt.Println("插入favorite失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "插入favorite失败",
			}, err
		}

		// 3. 将user的favorite_count+1
		// 从token中解析出用户名和密码，查找到user
		// 更新
		user.FavoriteCount++
		err = l.svcCtx.UserModel.Update(l.ctx, user)
		if err != nil {
			fmt.Println("更新user失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "更新user失败",
			}, err
		}

		// 4.将作品作者的total_favorited+1
		author, err := l.svcCtx.UserModel.FindOne(l.ctx, video.Author)
		if err != nil {
			fmt.Println("查找author失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "查找author失败",
			}, err
		}
		author.TotalFavorited++
		// 更新
		err = l.svcCtx.UserModel.Update(l.ctx, author)
		if err != nil {
			fmt.Println("更新author失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "更新author失败",
			}, err
		}
	} else if in.ActionType == 2 {
		// 取消 将video的favorite_count-1
		video, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
		if err != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "查询视频失败",
			}, err
		}
		video.FavoriteCount--
		// 更新
		err = l.svcCtx.VideoModel.Update(l.ctx, video)
		if err != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "视频喜欢数更新失败",
			}, err
		}

		// 2.将favorite中的表项删除
		// 通过user_id和video_id查找
		result, err := l.svcCtx.FavoriteModel.FindOneByUserVideo(l.ctx, user.Id, in.VideoId)
		err = l.svcCtx.FavoriteModel.Delete(l.ctx, result.Id)
		if err != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "从favorite中删除失败",
			}, err
		}
		// 3. 将user的favorite_count-1
		// 从token中解析出用户名和密码，查找到user
		// 更新
		user.FavoriteCount--
		err = l.svcCtx.UserModel.Update(l.ctx, user)
		if err != nil {
			fmt.Println("更新user失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "更新user失败",
			}, err
		}

		// 4.将作品作者的total_favorited-1
		author, err := l.svcCtx.UserModel.FindOne(l.ctx, video.Author)
		if err != nil {
			fmt.Println("查找author失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "查找author失败",
			}, err
		}
		author.TotalFavorited--
		// 更新
		err = l.svcCtx.UserModel.Update(l.ctx, author)
		if err != nil {
			fmt.Println("更新author失败")
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: 1,
				StatusMsg:  "更新author失败",
			}, err
		}
	} else {
		return &pb.DouyinFavoriteActionResponse{
			StatusCode: 1,
			StatusMsg:  "action_type 错误",
		}, errors.New("action_type 错误")
	}
	return &pb.DouyinFavoriteActionResponse{
		StatusCode: 0,
	}, nil
}

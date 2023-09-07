package logic

import (
	"bytes"
	"context"
	"douyin/model/video"
	"douyin/rpc/core/internal/JWT"
	oss "douyin/rpc/core/internal/OSSClient"
	"douyin/rpc/core/internal/svc"
	"douyin/rpc/core/pb"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DouyinPublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinPublishActionLogic {
	return &DouyinPublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinPublishActionLogic) DouyinPublishAction(in *pb.DouyinPublishActionRequest) (*pb.DouyinPublishActionResponse, error) {
	// todo: add your logic here and delete this line
	// tood: 使用事务，下面的操作1234
	fmt.Println("进入 rpc DouyinPublishAction")
	if err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token); err != nil {
		fmt.Println("token鉴权出错了")
		return &pb.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "token鉴权出错了",
		}, err
	}
	// 从token中解析出用户名
	claims, err := JWT.JWTAuth(in.Token)
	if err != nil {
		return &pb.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "鉴权错误",
		}, err
	}
	username := claims["Username"].(string)
	password := claims["Password"].(string)
	user, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, username, password)
	if err != nil {
		return &pb.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "根据用户名和密码查找用户出错了",
		}, err
	}
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectName := in.Title
	uploadPath := objectName + ".mp4"
	// author 从token中解析出用户名和密码，然后查找user_id
	insetVideo := &video.Video{
		Author:        user.Id,
		PlayUrl:       "http://little-douyin.oss-cn-chengdu.aliyuncs.com/" + uploadPath,
		CoverUrl:      "http://little-douyin.oss-cn-chengdu.aliyuncs.com/" + uploadPath + "?x-oss-process=video/snapshot,t_0,f_jpg,w_800,h_600",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         in.Title,
	}

	if err := l.svcCtx.UserModel.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		fmt.Println("进入 TransactCtx ")
		// 1.将video插入数据库中
		_, err := l.svcCtx.VideoModel.TransactionInsert(l.ctx, session, insetVideo)
		if err != nil {
			fmt.Println("1.将video插入数据库中出错了")
			return err
		}
		// 2.修改user的work_count
		user.WorkCount++
		err = l.svcCtx.UserModel.TransactionUpdate(l.ctx, session, user)
		if err != nil {
			fmt.Println("更新user的work_count出错")
			return err
		}
		return nil
		// 3.加入到 VideoList 中
		/*
			id, err := result.LastInsertId()
			if err != nil {
				fmt.Println("video 查询刚刚插入的视频出错了")
				return err
			}
			uploadVideo := pb.Video{
				Id: int64(len(oss.VideoList)),
				Author: &pb.User{
					Id:   id,
					Name: username,
				},
				PlayUrl:       "http://little-douyin.oss-cn-chengdu.aliyuncs.com/" + uploadPath,
				CoverUrl:      "http://little-douyin.oss-cn-chengdu.aliyuncs.com/" + uploadPath + "?x-oss-process=video/snapshot,t_0,f_jpg,w_800,h_600",
				Title:         in.Title,
				FavoriteCount: 0,
				CommentCount:  0,
				IsFavorite:    false,
			}
			oss.VideoList = append(oss.VideoList, &uploadVideo)
			//fmt.Println(oss.VideoList)
			return nil
		}*/
	}); err != nil {
		return &pb.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "上传视频失败了",
		}, err
	}

	// 4.上传到oss中
	//localFileName := "D:\\code\\GoLandProj\\simple-demo-main\\public\\bear.mp4"
	//err = oss.Bucket.PutObject(uploadPath, bytes.NewReader([]byte("in.Data")))
	err = oss.Bucket.PutObject(uploadPath, bytes.NewReader(in.Data))
	if err != nil {
		return &pb.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "上传到oss中出错了",
		}, err
	}
	return &pb.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "发布成功",
	}, nil
}

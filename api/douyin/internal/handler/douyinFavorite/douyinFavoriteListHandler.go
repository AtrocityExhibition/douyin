package douyinFavorite

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/douyinFavorite"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinFavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinFavoriteListRequest
		vars := r.URL.Query()
		if token, ok := vars["token"]; ok {
			req.Token = token[0]
		} else {
			fmt.Println("token错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("token错误"))
		}

		if userIdStr, ok := vars["user_id"]; ok {
			userId, err := strconv.ParseInt(userIdStr[0], 10, 64)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, errors.New("video_id格式错误"))
			}
			req.User_id = userId
		} else {
			fmt.Println("video_id错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id错误"))
		}

		l := douyinFavorite.NewDouyinFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.DouyinFavoriteList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

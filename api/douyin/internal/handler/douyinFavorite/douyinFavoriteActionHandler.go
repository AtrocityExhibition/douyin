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

func DouyinFavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinFavoriteActionRequest
		req.Token = r.FormValue("token")
		videoId, err := strconv.ParseInt(r.FormValue("video_id"), 10, 64)
		if err != nil {
			fmt.Println("video_id 格式错误", err)
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id 格式错误"))
		}
		req.Video_id = videoId
		actionType, err := strconv.ParseInt(r.FormValue("action_type"), 10, 64)
		if err != nil {
			fmt.Println("action_type 格式错误", err)
			httpx.ErrorCtx(r.Context(), w, errors.New("action_type 格式错误"))
		}
		req.Action_type = int32(actionType)

		l := douyinFavorite.NewDouyinFavoriteActionLogic(r.Context(), svcCtx)
		fmt.Println("调用 api DouyinFavoriteAction")
		resp, err := l.DouyinFavoriteAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

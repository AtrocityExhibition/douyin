package douyinComment

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/douyinComment"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinCommentListRequest
		req.Token = r.FormValue("token")
		videoId, err := strconv.ParseInt(r.FormValue("video_id"), 10, 64)
		if err != nil {
			fmt.Println("video_id格式错误", err)
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id格式错误"))
		}
		req.Video_id = videoId

		l := douyinComment.NewDouyinCommentListLogic(r.Context(), svcCtx)
		resp, err := l.DouyinCommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

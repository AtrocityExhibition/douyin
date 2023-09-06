package feed

import (
	"douyin/api/core/internal/logic/feed"
	"douyin/api/core/internal/svc"
	"douyin/api/core/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

func DouyinFeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinFeedReq
		req.Token = r.FormValue("token")
		latestTime, err := strconv.ParseInt(r.FormValue("latest_time"), 10, 64)
		if err != nil {
			fmt.Println("latest_time格式错误")
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.Latest_time = latestTime

		l := feed.NewDouyinFeedLogic(r.Context(), svcCtx)
		resp, err := l.DouyinFeed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

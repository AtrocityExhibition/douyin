package feed

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"douyin/api/douyin/internal/logic/feed"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinFeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinFeedReq
		fmt.Println("httpx.ParseForm 开始解析请求参数")
		vars := r.URL.Query()
		if token, ok := vars["token"]; ok {
			req.Token = token[0]
		} else {
			req.Token = ""
		}
		defaultTime := time.Now().UnixMilli()
		if latestTimeStr, ok := vars["latest_time"]; ok {
			latestTime, err := strconv.ParseInt(latestTimeStr[0], 10, 64)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			req.Latest_time = latestTime
		} else {
			req.Latest_time = defaultTime
		}

		fmt.Println("解析完请求参数")
		l := feed.NewDouyinFeedLogic(r.Context(), svcCtx)
		resp, err := l.DouyinFeed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

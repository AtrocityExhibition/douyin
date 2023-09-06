package publish

import (
	"bytes"
	"io"
	"net/http"

	"douyin/api/douyin/internal/logic/publish"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinPublishActionrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinPublishActionrReq
		file, _, err := r.FormFile("data")
		if err != nil {
			return
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return
		}
		req.Data = buf.Bytes()
		req.Title = r.FormValue("title")
		req.Token = r.FormValue("token")

		l := publish.NewDouyinPublishActionrLogic(r.Context(), svcCtx)
		resp, err := l.DouyinPublishActionr(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

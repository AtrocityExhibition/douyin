package publish

import (
	"bytes"
	"io"
	"net/http"

	"douyin/api/core/internal/logic/publish"
	"douyin/api/core/internal/svc"
	"douyin/api/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinPublishActionrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinPublishActionrReq
		if err := r.ParseMultipartForm(1024); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publish.NewDouyinPublishActionrLogic(r.Context(), svcCtx)
		file, _, err := r.FormFile("data")
		defer file.Close()
		if err != nil {
			return
		}

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			return
		}

		req.Data = buf.Bytes()
		req.Title = r.FormValue("title")
		req.Token = r.FormValue("token")

		resp, err := l.DouyinPublishActionr(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

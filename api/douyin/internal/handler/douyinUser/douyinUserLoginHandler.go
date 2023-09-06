package douyinUser

import (
	"errors"
	"fmt"
	"net/http"

	"douyin/api/douyin/internal/logic/douyinUser"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinUserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinUserLoginReq
		vars := r.URL.Query()
		if username, ok := vars["username"]; ok {
			req.Username = username[0]
		} else {
			fmt.Println("用户名错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("用户名错误"))
		}
		if password, ok := vars["password"]; ok {
			req.Password = password[0]
		} else {
			fmt.Println("密码错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("密码错误"))
		}

		l := douyinUser.NewDouyinUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.DouyinUserLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

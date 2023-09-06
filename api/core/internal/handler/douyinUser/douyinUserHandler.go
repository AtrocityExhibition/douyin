package douyinUser

import (
	"douyin/api/core/internal/logic/douyinUser"
	"douyin/api/core/internal/svc"
	"douyin/api/core/internal/types"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

func DouyinUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinUserReq
		vars := r.URL.Query()
		if token, ok := vars["token"]; ok {
			req.Token = token[0]
		} else {
			fmt.Println("未输入token")
			httpx.ErrorCtx(r.Context(), w, errors.New("未输入token"))
		}
		if userIdStr, ok := vars["user_id"]; ok {
			userId, err := strconv.ParseInt(userIdStr[0], 10, 64)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			req.User_id = userId
		} else {
			fmt.Println("未输入user_id")
			httpx.ErrorCtx(r.Context(), w, errors.New("未输入user_id"))
		}

		l := douyinUser.NewDouyinUserLogic(r.Context(), svcCtx)
		resp, err := l.DouyinUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

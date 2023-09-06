package publish

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/publish"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinPublishListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinPublishListReq
		req.Token = r.FormValue("token")
		userId, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
		if err != nil {
			fmt.Println("user_id格式错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("user_id格式错误"))
		}
		req.User_id = userId
		fmt.Println("user_id = ", userId)

		l := publish.NewDouyinPublishListLogic(r.Context(), svcCtx)
		fmt.Println("调用 api 的 DouyinPublishList 返回")
		resp, err := l.DouyinPublishList(&req)
		fmt.Println("调用 api 的 DouyinPublishList 返回")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

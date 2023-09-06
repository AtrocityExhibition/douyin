package douyinRelation

import (
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/douyinRelation"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinRelationFollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinRelationFollowListRequest
		req.Token = r.FormValue("token")
		userId, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
		if err != nil {
			fmt.Println("user_id 格式错误")
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.User_id = userId

		l := douyinRelation.NewDouyinRelationFollowListLogic(r.Context(), svcCtx)
		resp, err := l.DouyinRelationFollowList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

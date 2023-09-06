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

func DouyinRelationActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinRelationActionRequest
		req.Token = r.FormValue("token")
		toUserId, err := strconv.ParseInt(r.FormValue("to_user_id"), 10, 64)
		if err != nil {
			fmt.Println("to_user_id 格式错误")
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.To_user_id = toUserId
		actionType, err := strconv.ParseInt(r.FormValue("action_type"), 10, 64)
		if err != nil {
			fmt.Println("action_type 格式错误")
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.Action_type = int32(actionType)
		fmt.Println(req.Token, " to_user_id: ", req.To_user_id, " action_type: ", req.Action_type)

		l := douyinRelation.NewDouyinRelationActionLogic(r.Context(), svcCtx)
		resp, err := l.DouyinRelationAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

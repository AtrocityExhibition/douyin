package douyinMessage

import (
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/douyinMessage"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinMessageChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinMessageChatRequest
		req.Token = r.FormValue("token")
		toUserId, err := strconv.ParseInt(r.FormValue("to_user_id"), 10, 64)
		if err != nil {
			fmt.Println("to_user_id格式错误")
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.To_user_id = toUserId
		preMsgTime, err := strconv.ParseInt(r.FormValue("pre_msg_time"), 10, 64)
		if err != nil {
			fmt.Println("pre_msg_time格式错误")
			httpx.ErrorCtx(r.Context(), w, err)
		}
		req.Pre_msg_time = preMsgTime
		fmt.Println("token = ", req.Token, "to_user_id = ", req.To_user_id, "pre_msg_time = ", req.Pre_msg_time)
		l := douyinMessage.NewDouyinMessageChatLogic(r.Context(), svcCtx)
		resp, err := l.DouyinMessageChat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

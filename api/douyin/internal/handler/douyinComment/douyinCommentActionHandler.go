package douyinComment

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"douyin/api/douyin/internal/logic/douyinComment"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DouyinCommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DouyinCommentActionRequest
		req.Token = r.FormValue("token")
		videoId, err := strconv.ParseInt(r.FormValue("video_id"), 10, 64)
		if err != nil {
			fmt.Println("video_id格式错误", err)
			httpx.ErrorCtx(r.Context(), w, errors.New("video_id格式错误"))
		}
		req.Video_id = videoId
		actionType, err := strconv.ParseInt(r.FormValue("action_type"), 10, 64)
		if err != nil {
			fmt.Println("action_type格式错误", err)
			httpx.ErrorCtx(r.Context(), w, errors.New("action_type格式错误"))
		}
		req.Action_type = int32(actionType)
		req.Comment_text = r.FormValue("comment_text")
		commentId := r.FormValue("comment_id")
		if commentId != "" {
			req.Comment_id, err = strconv.ParseInt(commentId, 10, 64)
			if err != nil {
				fmt.Println("comment_id格式错误", err)
				httpx.ErrorCtx(r.Context(), w, errors.New("comment_id格式错误"))
			}
		}

		fmt.Println(req.Token, req.Video_id, req.Action_type, req.Comment_text, req.Comment_id)

		l := douyinComment.NewDouyinCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.DouyinCommentAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

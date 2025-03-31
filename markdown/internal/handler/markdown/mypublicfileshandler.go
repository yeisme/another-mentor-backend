package markdown

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"markdown/internal/logic/markdown"
	"markdown/internal/svc"
	"markdown/internal/types"
)

// 获取当前用户公开文件列表
func MyPublicFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := markdown.NewMyPublicFilesLogic(r.Context(), svcCtx)
		resp, err := l.MyPublicFiles(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

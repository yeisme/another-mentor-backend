package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"markdown/internal/logic/public"
	"markdown/internal/svc"
	"markdown/internal/types"
)

// 搜索公开文件
func SearchPublicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewSearchPublicLogic(r.Context(), svcCtx)
		resp, err := l.SearchPublic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

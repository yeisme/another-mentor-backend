package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"markdown/internal/logic/public"
	"markdown/internal/svc"
	"markdown/internal/types"
)

// 访问分享链接
func AccessSharedFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareAccessRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewAccessSharedFileLogic(r.Context(), svcCtx)
		resp, err := l.AccessSharedFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

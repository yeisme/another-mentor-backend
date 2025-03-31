package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"markdown/internal/logic/public"
	"markdown/internal/svc"
)

// 热门标签
func PopularTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewPopularTagsLogic(r.Context(), svcCtx)
		resp, err := l.PopularTags()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

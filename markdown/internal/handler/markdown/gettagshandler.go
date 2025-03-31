package markdown

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"markdown/internal/logic/markdown"
	"markdown/internal/svc"
)

// 获取用户所有标签
func GetTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := markdown.NewGetTagsLogic(r.Context(), svcCtx)
		resp, err := l.GetTags()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

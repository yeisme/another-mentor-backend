package public

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPublicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 搜索公开文件
func NewSearchPublicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPublicLogic {
	return &SearchPublicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchPublicLogic) SearchPublic(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

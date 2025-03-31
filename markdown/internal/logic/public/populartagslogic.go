package public

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PopularTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 热门标签
func NewPopularTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PopularTagsLogic {
	return &PopularTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PopularTagsLogic) PopularTags() (resp *types.TagsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

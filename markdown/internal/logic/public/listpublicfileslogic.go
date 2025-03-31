package public

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPublicFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 浏览公开文件列表
func NewListPublicFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPublicFilesLogic {
	return &ListPublicFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPublicFilesLogic) ListPublicFiles(req *types.ListRequest) (resp *types.ListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

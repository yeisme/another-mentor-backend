package logic

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkdownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarkdownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkdownLogic {
	return &MarkdownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarkdownLogic) Markdown(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

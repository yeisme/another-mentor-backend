package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyPublicFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户公开文件列表
func NewMyPublicFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyPublicFilesLogic {
	return &MyPublicFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyPublicFilesLogic) MyPublicFiles(req *types.ListRequest) (resp *types.ListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

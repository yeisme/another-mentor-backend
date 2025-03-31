package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户的文件列表
func NewMyFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyFilesLogic {
	return &MyFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyFilesLogic) MyFiles(req *types.ListRequest) (resp *types.ListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

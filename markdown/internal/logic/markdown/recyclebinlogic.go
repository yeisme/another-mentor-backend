package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecycleBinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户回收站文件
func NewRecycleBinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecycleBinLogic {
	return &RecycleBinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecycleBinLogic) RecycleBin(req *types.ListRequest) (resp *types.ListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

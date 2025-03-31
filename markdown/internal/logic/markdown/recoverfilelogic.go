package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 恢复已删除文件
func NewRecoverFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverFileLogic {
	return &RecoverFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverFileLogic) RecoverFile(req *types.RecoverFileRequest) (resp *types.RecoverFileResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

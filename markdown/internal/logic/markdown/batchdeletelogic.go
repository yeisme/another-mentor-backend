package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除文件
func NewBatchDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteLogic {
	return &BatchDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteLogic) BatchDelete(req *types.BatchDeleteRequest) (resp *types.BatchDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

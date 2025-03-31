package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyTrashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 清空回收站
func NewEmptyTrashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyTrashLogic {
	return &EmptyTrashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmptyTrashLogic) EmptyTrash() (resp *types.DeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package public

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取公开文件
func NewGetPublicFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicFileLogic {
	return &GetPublicFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicFileLogic) GetPublicFile(req *types.DownloadResponse) error {
	// todo: add your logic here and delete this line

	return nil
}

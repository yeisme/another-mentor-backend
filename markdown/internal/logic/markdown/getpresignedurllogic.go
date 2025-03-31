package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPresignedUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取预签名URL
func NewGetPresignedUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPresignedUrlLogic {
	return &GetPresignedUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPresignedUrlLogic) GetPresignedUrl(req *types.PresignedUrlRequest) (resp *types.PresignedUrlResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

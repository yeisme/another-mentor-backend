package public

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessSharedFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 访问分享链接
func NewAccessSharedFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessSharedFileLogic {
	return &AccessSharedFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccessSharedFileLogic) AccessSharedFile(req *types.ShareAccessRequest) (resp *types.DownloadResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

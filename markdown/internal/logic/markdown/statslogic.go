package markdown

import (
	"context"

	"markdown/internal/svc"
	"markdown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文件统计信息
func NewStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatsLogic {
	return &StatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatsLogic) Stats() (resp *types.StatsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

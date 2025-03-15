package user

import (
	"context"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerifyCodeLogic {
	return &SendVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TODO  将生成的验证码保存到缓存中，然后发送到用户邮箱，对比用户输入的验证码是否正确（登录 logic 中）
func (l *SendVerifyCodeLogic) SendVerifyCode(req *types.SendVerifyCodeRequest) (resp *types.SendVerifyCodeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

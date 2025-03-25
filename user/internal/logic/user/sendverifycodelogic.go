package user

import (
	"context"
	"net/http"
	"time"

	"user/internal/svc"
	"user/internal/types"

	"github.com/mojocn/base64Captcha"
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

// TODO  将生成的验证码保存到缓存中
func (l *SendVerifyCodeLogic) SendVerifyCode(req *types.SendVerifyCodeRequest) (resp *types.SendVerifyCodeResponse, err error) {
	// 配置验证码参数：高度80，宽度240，5位数字，背景复杂度0.7，最大干扰线80
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	// 使用默认内存存储生成验证码（临时保存验证码答案）
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	captchaId, captchaB64, _, err := c.Generate()
	if err != nil {
		logx.Errorf("存储验证码到redis失败: %v", err)
		return nil, err
	}

	// 从默认内存存储中获取验证码答案
	code := base64Captcha.DefaultMemStore.Get(captchaId, false)

	// 将验证码保存到redis中（5分钟过期）
	// 请确保 svcCtx.Redis 已经初始化为 *redis.Client
	err = l.svcCtx.Redis.Set(l.ctx, captchaId, code, 5*time.Minute).Err()
	if err != nil {
		logx.Errorf("存储验证码到redis失败: %v", err)
		return nil, err
	}

	resp = &types.SendVerifyCodeResponse{
		BaseResponse: types.BaseResponse{
			Code: http.StatusOK,
			Msg:  "发送验证码成功",
		},
		CaptchaId:    captchaId,
		CaptchaImage: captchaB64,
	}
	return resp, nil
}

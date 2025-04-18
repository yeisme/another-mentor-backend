package user

import (
	"context"
	"net/http"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {

	// 参数验证
	if req.Username == "" || req.Password == "" || req.Email == "" || req.Code == "" {
		return &types.RegisterResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusBadRequest,
				Msg:  "参数错误",
			},
		}, nil
	}

	// 检查验证码， 对比redis中的验证码
    storedCode, err := l.svcCtx.Redis.Get(l.ctx, req.CaptchaId).Result()

	logx.Infof("storedCode: %v", storedCode)

	if err != nil {
        logx.Errorf("获取验证码失败: %v", err)
        return &types.RegisterResponse{
            BaseResponse: types.BaseResponse{
                Code: http.StatusInternalServerError,
                Msg:  "验证码验证失败，可能缺少 CaptchaId",
            },
        }, nil
    }

	if storedCode != req.Code {
		logx.Errorf("验证码不匹配")
		return &types.RegisterResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusBadRequest,
				Msg:  "验证码不匹配",
			},
		}, nil
	}

	user, err := l.svcCtx.UserRepo.Register(req.Username, req.Password, req.Email)

	if err != nil {
		// 错误处理
		return &types.RegisterResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusInternalServerError,
				Msg:  "注册失败",
			},
		}, nil
	}

	// 构造响应
	return &types.RegisterResponse{
		BaseResponse: types.BaseResponse{
			Code: http.StatusOK,
			Msg:  "注册成功",
		},
		UserId: user.ID,
	}, nil
}

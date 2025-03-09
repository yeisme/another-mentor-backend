package user

import (
	"context"

	"user/internal/data/repository"
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
				Code: 400,
				Msg:  "参数错误",
			},
		}, nil
	}

	// 调用repository进行注册
	userRepo := repository.NewUserRepository()
	user, err := userRepo.Register(req.Username, req.Password, req.Email)
	if err != nil {
		// 错误处理
		return &types.RegisterResponse{
			BaseResponse: types.BaseResponse{
				Code: 400,
				Msg:  "注册失败",
			},
		}, nil
	}

	// 构造响应
	return &types.RegisterResponse{
		BaseResponse: types.BaseResponse{
			Code: 200,
			Msg:  "注册成功",
		},
		UserId: user.ID,
	}, nil
}

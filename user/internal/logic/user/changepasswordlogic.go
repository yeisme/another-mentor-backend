package user

import (
	"context"
	"net/http"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordRequest) (resp *types.ChangePasswordResponse, err error) {

	if req.OldPassword == "" || req.NewPassword == "" {
		return &types.ChangePasswordResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusBadRequest,
				Msg:  "密码不能为空",
			},
		}, err
	}

	// UserId 通过 token 获取
	userId, _ := l.ctx.Value("userId").(int64)
	err = l.svcCtx.UserRepo.UpdatePassword(userId, req.OldPassword, req.NewPassword)

	if err != nil {
		return &types.ChangePasswordResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusInternalServerError,
				Msg:  "密码修改失败: " + err.Error(),
			},
		}, err
	}

	return &types.ChangePasswordResponse{
		BaseResponse: types.BaseResponse{
			Code: http.StatusOK,
			Msg:  "密码修改成功",
		},
	}, nil
}

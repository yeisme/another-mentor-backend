package user

import (
	"context"
	"net/http"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoRequest) (resp *types.UpdateUserInfoResponse, err error) {

	userId, _ := l.ctx.Value("userId").(int64)

	// 调用Repository更新用户信息
	err = l.svcCtx.UserRepo.UpdateProfile(userId, req.Nickname, req.Avatar, req.Introduction)

	if err != nil {
		return &types.UpdateUserInfoResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusInternalServerError,
				Msg:  "更新用户信息失败: " + err.Error(),
			},
		}, nil
	}

	return &types.UpdateUserInfoResponse{
		BaseResponse: types.BaseResponse{
			Code: http.StatusOK,
			Msg:  "用户信息更新成功",
		},
	}, nil
}

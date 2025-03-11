package user

import (
	"context"
	"net/http"

	"user/internal/data/repository"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {

	// 确定要查询的用户ID
	var queryUserId int64

	// 如果请求中指定了用户ID，则查询该用户
	if req.UserId > 0 {
		queryUserId = req.UserId
	} else {
		// 否则获取当前登录用户ID
		userId, _ := l.ctx.Value("userId").(int64)

		queryUserId = userId
	}

	// 获取用户信息
	userRepo := repository.NewUserRepository()
	userModel, err := userRepo.GetUserByID(queryUserId)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return &types.GetUserInfoResponse{
				BaseResponse: types.BaseResponse{
					Code: http.StatusNotFound,
					Msg:  "用户不存在",
				},
			}, nil
		} else {
			return &types.GetUserInfoResponse{
				BaseResponse: types.BaseResponse{
					Code: http.StatusInternalServerError,
					Msg:  "获取用户信息失败: " + err.Error(),
				},
			}, nil
		}

	}

	user := &types.User{
		Id:           userModel.ID,
		Username:     userModel.Username,
		Email:        userModel.Email,
		Phone:        userModel.Phone,
		Avatar:       userModel.Avatar,
		Nickname:     userModel.Nickname,
		Introduction: userModel.Introduction,
		CreateTime:   userModel.CreatedAt.Unix(),
		UpdateTime:   userModel.UpdatedAt.Unix(),
		Status:       userModel.Status,
	}

	return &types.GetUserInfoResponse{
		BaseResponse: types.BaseResponse{
			Code: http.StatusOK,
			Msg:  "获取用户信息成功",
		},
		User: *user,
	}, nil
}

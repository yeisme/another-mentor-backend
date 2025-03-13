package admin

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详情
func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {

	if req.ID == "" {
		return &types.UserDetailResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusBadRequest,
				Msg:  "用户ID不能为空",
			}, User: types.User{},
		}, nil
	}

	// 转换ID为整数
	userID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		return &types.UserDetailResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusBadRequest,
				Msg:  fmt.Sprintf("无效的用户ID格式: %v", err),
			}, User: types.User{},
		}, nil
	}

	// 获取用户信息
	user, err := l.svcCtx.AdminRepo.GetUserByID(userID)
	if err != nil {
		return &types.UserDetailResponse{
			BaseResponse: types.BaseResponse{
				Code: http.StatusInternalServerError,
				Msg:  "查询用户信息失败，可能用户不存在",
			}, User: types.User{},
		}, nil
	}

	return &types.UserDetailResponse{
		BaseResponse: types.BaseResponse{
			Code: http.StatusOK,
			Msg:  "查询用户信息成功",
		}, User: types.User{
			Id:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			Phone:        user.Phone,
			Avatar:       user.Avatar,
			Nickname:     user.Nickname,
			Introduction: user.Introduction,
			CreateTime:   user.CreatedAt.Unix(),
			UpdateTime:   user.UpdatedAt.Unix(),
			Status:       user.Status,
		},
	}, nil
}

package admin

import (
	"context"
	"net/http"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (resp *types.UserListResponse, err error) {

	resp = &types.UserListResponse{
		BaseResponse: types.BaseResponse{
			Code: http.StatusOK,
			Msg:  "获取用户列表成功",
		},
	}

	userList, total, err := l.svcCtx.AdminRepo.GetUserList(req.Page, req.PageSize, req.Keyword, req.Status)

	resp.Total = total
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Msg = "获取用户列表失败"
		return resp, nil
	}

	// 转换 []*model.User 到 []types.User
	typesUserList := make([]types.User, len(userList))
	for i, user := range userList {
		// 根据字段进行映射，以下是示例
		typesUserList[i] = types.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			Avatar:   user.Avatar,
			Nickname: user.Nickname,
			Introduction: user.Introduction,
			CreateTime:   user.CreatedAt.Unix(),
			UpdateTime:   user.UpdatedAt.Unix(),
			Status:   user.Status,

		}
	}
	resp.List = typesUserList

	return resp, nil
}

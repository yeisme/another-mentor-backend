package user

import (
	"context"
	"net/http"
	"time"

	"user/internal/svc"
	"user/internal/types"
	"user/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
    return &LoginLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
    // 参数验证
    if req.Account == "" || req.Password == "" {
        return &types.LoginResponse{
            BaseResponse: types.BaseResponse{
                Code: http.StatusBadRequest,
                Msg:  "参数错误",
            },
        }, nil
    }

	// TODO 对比 redis 缓存的验证码与接收到的验证码是否一致

    // 调用repository进行登录
	user, err := l.svcCtx.UserRepo.Login(req.Account, req.Password)

    if err != nil {
        return &types.LoginResponse{
            BaseResponse: types.BaseResponse{
                Code: http.StatusInternalServerError,
                Msg:  "登录失败: " + err.Error(),
            },
        }, nil
    }

    // 创建JWT实例
    jwt := utils.NewJWT(l.svcCtx.Config.Auth.AccessSecret)
    // 设置过期时间
    duration := time.Duration(l.svcCtx.Config.Auth.AccessExpire) * time.Second
    jwt.WithExpireDuration(duration)

    // 创建JWT声明
    claims := utils.JWTClaims{
        UserID:   user.ID,
        Username: user.Username,
        Role:     user.Role, // 将角色转换为字符串
    }

    // 生成令牌
    token, expiresAt, err := jwt.CreateToken(claims)
    if err != nil {
        return &types.LoginResponse{
            BaseResponse: types.BaseResponse{
                Code: http.StatusInternalServerError,
                Msg:  "生成token失败",
            },
        }, nil
    }

    // 构造用户信息返回
    userInfo := types.User{
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
    }

    return &types.LoginResponse{
        BaseResponse: types.BaseResponse{
            Code: http.StatusOK,
            Msg:  "登录成功",
        },
        Token:     token,
        ExpiresIn: expiresAt, // token有效期
        User:      userInfo,
    }, nil
}

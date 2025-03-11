package middleware

import (
	"context"
	"net/http"
	"strings"
	"user/internal/config"
	"user/internal/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 定义上下文key名称
const (
	UserIdKey   = "userId"
	UsernameKey = "username"
	RoleKey     = "role"
)

type UserAuthMiddleware struct {
	Config config.Config
}

func NewUserAuthMiddleware(c config.Config) *UserAuthMiddleware {
	return &UserAuthMiddleware{
		Config: c,
	}
}

func (m *UserAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 1. 从请求头获取Authorization token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]any{
				"code": http.StatusUnauthorized,
				"msg":  "请先登录",
			})
			return
		}

		// 2. 提取 token 值 (Bearer token)
		parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            httpx.WriteJson(w, http.StatusUnauthorized, map[string]any{
                "code": http.StatusUnauthorized,
                "msg":  "认证失败：token格式不正确",
            })
            return
        }
        token := parts[1]

		// 3. 验证 token
		jwt := utils.NewJWT(m.Config.Auth.AccessSecret)
		claims, err := jwt.ParseToken(token)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]any{
				"code": http.StatusUnauthorized,
				"msg":  "认证失败" + err.Error(),
			})
			return
		}

		// 4. 验证身份为 user 或者 admin
		if claims.Role != "user" && claims.Role != "admin" {
			httpx.WriteJson(w, http.StatusForbidden, map[string]any{
				"code": http.StatusForbidden,
				"msg":  "无权限访问",
			})
			return
		}

		// 5. 将用户信息存入请求上下文
		ctx := context.WithValue(r.Context(), UserIdKey, claims.UserID)
		ctx = context.WithValue(ctx, UsernameKey, claims.Username)
		ctx = context.WithValue(ctx, RoleKey, claims.Role)

		next(w, r.WithContext(ctx))
	}
}

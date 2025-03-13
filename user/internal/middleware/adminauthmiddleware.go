package middleware

import (
	"context"
	"net/http"
	"strings"
	"user/internal/config"
	"user/internal/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AdminAuthMiddleware struct {
	Config config.Config
}

func NewAdminAuthMiddleware(c config.Config) *AdminAuthMiddleware {
	return &AdminAuthMiddleware{
		Config: c,
	}
}

func (m *AdminAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
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
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  "认证格式错误",
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

		// 4. 验证身份为 admin
		if claims.Role != "admin" {
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

package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired     = errors.New("token已过期")
	ErrTokenNotValidYet = errors.New("token尚未生效")
	ErrTokenMalformed   = errors.New("token格式错误")
	ErrTokenInvalid     = errors.New("非法token")
)

type JWTClaims struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// JWTConfig JWT配置选项
type JWTConfig struct {
	SigningKey          []byte
	SigningMethod       jwt.SigningMethod
	TokenExpireDuration time.Duration
	Issuer              string
}

type JWT struct {
	config JWTConfig
}

// NewJWT 创建JWT实例
func NewJWT(signingKey string) *JWT {
	return &JWT{
		config: JWTConfig{
			SigningKey:          []byte(signingKey),
			SigningMethod:       jwt.SigningMethodHS256,
			TokenExpireDuration: time.Hour * 24 * 7, // 默认7天有效期
			Issuer:              "another-mentor",
		},
	}
}

// WithExpireDuration 设置token过期时间
func (j *JWT) WithExpireDuration(duration time.Duration) *JWT {
	j.config.TokenExpireDuration = duration
	return j
}

// WithIssuer 设置token发行人
func (j *JWT) WithIssuer(issuer string) *JWT {
	j.config.Issuer = issuer
	return j
}

// CreateToken 创建JWT Token
func (j *JWT) CreateToken(claims JWTClaims) (string, int64, error) {
	// 设置token有效期
	expiresAt := time.Now().Add(j.config.TokenExpireDuration)
	claims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.Issuer = j.config.Issuer

	// 可选：设置NotBefore，提高安全性
	claims.NotBefore = jwt.NewNumericDate(time.Now())

	// 创建token
	token := jwt.NewWithClaims(j.config.SigningMethod, claims)
	tokenString, err := token.SignedString(j.config.SigningKey)

	return tokenString, expiresAt.Unix(), err
}

// ParseToken 解析JWT Token
func (j *JWT) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return j.config.SigningKey, nil
	})

	if err != nil {
		// 使用 jwt/v5 的错误处理机制
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, ErrTokenNotValidYet
		} else {
			return nil, ErrTokenInvalid
		}
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// ValidateTokenWithoutExpiration 验证令牌但忽略过期时间（用于特殊场景如刷新令牌）
func (j *JWT) ValidateTokenWithoutExpiration(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return j.config.SigningKey, nil
	}, jwt.WithoutClaimsValidation())

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 刷新Token
func (j *JWT) RefreshToken(tokenString string) (string, int64, error) {
	// 先验证旧token，忽略过期时间
	claims, err := j.ValidateTokenWithoutExpiration(tokenString)
	if err != nil {
		return "", 0, err
	}

	// 创建新token
	return j.CreateToken(*claims)
}

/*

jwt := NewJWT("your-secret-key").WithExpireDuration(time.Hour * 24).WithIssuer("api.example.com")

claims := JWTClaims{
	UserID:   123,
	Username: "user",
	Role:     "admin",
}

token, expires, err := jwt.CreateToken(claims)

// 解析token
claims, err := jwt.ParseToken(token)

// 刷新token
newToken, newExpires, err := jwt.RefreshToken(oldToken)
*/

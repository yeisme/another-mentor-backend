package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 测试常量
const (
	testSigningKey = "test-signing-key-for-jwt"
	testUsername   = "testuser"
)

// 测试创建和解析 Token
func TestCreateAndParseToken(t *testing.T) {
	// 创建 JWT 实例
	j := NewJWT(testSigningKey)

	// 创建测试用户声明
	claims := JWTClaims{
		UserID:   123456,
		Username: testUsername,
		Role:     "user",
	}

	// 创建 Token
	token, expires, err := j.CreateToken(claims)

	// 验证创建结果
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Greater(t, expires, int64(0))

	// 解析 Token
	parsedClaims, err := j.ParseToken(token)

	// 验证解析结果
	assert.NoError(t, err)
	assert.NotNil(t, parsedClaims)
	assert.Equal(t, claims.UserID, parsedClaims.UserID)
	assert.Equal(t, claims.Username, parsedClaims.Username)
	assert.Equal(t, claims.Role, parsedClaims.Role)
}

// 测试配置选项
func TestJWTOptions(t *testing.T) {
	// 创建带有自定义选项的 JWT 实例
	customExpire := time.Hour * 2
	customIssuer := "test-issuer"

	j := NewJWT(testSigningKey).
		WithExpireDuration(customExpire).
		WithIssuer(customIssuer)

	// 验证配置已应用
	assert.Equal(t, customExpire, j.config.TokenExpireDuration)
	assert.Equal(t, customIssuer, j.config.Issuer)

	// 创建并解析 Token 验证配置生效
	claims := JWTClaims{
		UserID:   123,
		Username: testUsername,
		Role:     "user",
	}

	token, _, err := j.CreateToken(claims)
	assert.NoError(t, err)

	parsedClaims, err := j.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, customIssuer, parsedClaims.Issuer)
}

// 测试无效 Token
func TestInvalidToken(t *testing.T) {
	j := NewJWT(testSigningKey)

	// 1. 格式错误的 Token
	_, err := j.ParseToken("invalid-token-format")
	assert.Error(t, err)
	assert.Equal(t, ErrTokenMalformed, err)

	// 2. 篡改的 Token
	claims := JWTClaims{
		UserID:   1,
		Username: testUsername,
	}
	token, _, _ := j.CreateToken(claims)
	tamperedToken := token + "tampered"
	_, err = j.ParseToken(tamperedToken)
	assert.Error(t, err)
	assert.Equal(t, ErrTokenInvalid, err)

	// 3. 使用不同密钥签名的 Token
	j2 := NewJWT("different-signing-key")
	token2, _, _ := j2.CreateToken(claims)
	_, err = j.ParseToken(token2)
	assert.Error(t, err)
	assert.Equal(t, ErrTokenInvalid, err)
}

// 测试 Token 过期
func TestExpiredToken(t *testing.T) {
	// 创建一个超短过期时间的 JWT
	j := NewJWT(testSigningKey).WithExpireDuration(time.Millisecond * 100)

	claims := JWTClaims{
		UserID:   1,
		Username: testUsername,
	}

	token, _, err := j.CreateToken(claims)
	assert.NoError(t, err)

	// 等待 Token 过期
	time.Sleep(time.Millisecond * 200)

	// 验证 Token 已过期
	_, err = j.ParseToken(token)
	assert.Error(t, err)
	assert.Equal(t, ErrTokenExpired, err)
}

// 测试 Token 刷新
// 测试刷新无效Token
func TestRefreshInvalidToken(t *testing.T) {
	j := NewJWT(testSigningKey)

	// 测试刷新格式错误的Token
	_, _, err := j.RefreshToken("invalid-token-format")
	assert.Error(t, err)

	// 测试刷新被篡改的Token
	claims := JWTClaims{
		UserID:   100,
		Username: testUsername,
		Role:     "user",
	}
	token, _, _ := j.CreateToken(claims)
	tamperedToken := token + "tampered"
	_, _, err = j.RefreshToken(tamperedToken)
	assert.Error(t, err)

	// 测试使用不同密钥签发的Token
	j2 := NewJWT("different-signing-key")
	token2, _, _ := j2.CreateToken(claims)
	_, _, err = j.RefreshToken(token2)
	assert.Error(t, err)
}

// 测试多次刷新Token
func TestMultipleRefresh(t *testing.T) {
	j := NewJWT(testSigningKey).WithExpireDuration(time.Second * 1)

	claims := JWTClaims{
		UserID:   200,
		Username: testUsername,
		Role:     "user",
	}

	originalToken, _, err := j.CreateToken(claims)
	assert.NoError(t, err)

	// 等待原Token过期
	time.Sleep(time.Second * 2)

	secondRefresh, _, err := j.RefreshToken(originalToken)
	assert.NoError(t, err)
	assert.NotEqual(t, originalToken, secondRefresh)

	// 解析最终Token，确保所有声明都正确保留
	parsedClaims, err := j.ParseToken(secondRefresh)
	assert.NoError(t, err)
	assert.Equal(t, claims.UserID, parsedClaims.UserID)
	assert.Equal(t, claims.Username, parsedClaims.Username)
	assert.Equal(t, claims.Role, parsedClaims.Role)
}

// 测试极端过期时间
func TestExtremeDurations(t *testing.T) {
	// 超短过期时间
	jShort := NewJWT(testSigningKey).WithExpireDuration(time.Millisecond * 50)

	// 超长过期时间
	jLong := NewJWT(testSigningKey).WithExpireDuration(time.Hour * 8760) // ~1 year

	claims := JWTClaims{
		UserID:   300,
		Username: testUsername,
		Role:     "guest",
	}

	// 测试超短过期
	shortToken, _, err := jShort.CreateToken(claims)
	assert.NoError(t, err)
	time.Sleep(time.Millisecond * 100) // 确保过期
	_, _, err = jShort.RefreshToken(shortToken)
	assert.NoError(t, err) // 刷新应该仍然有效

	// 测试超长过期
	longToken, expires, err := jLong.CreateToken(claims)
	_ = longToken

	assert.NoError(t, err)
	// 检查过期时间是否在将来一年左右
	assert.Greater(t, expires, time.Now().Unix()+365*24*3600-3600) // 允许1小时误差
}

// 测试组合配置选项
func TestCombinedOptions(t *testing.T) {
	customDuration := time.Hour * 12
	customIssuer := "combined-test-issuer"

	j := NewJWT(testSigningKey).
		WithExpireDuration(customDuration).
		WithIssuer(customIssuer)

	claims := JWTClaims{
		UserID:   400,
		Username: testUsername,
		Role:     "tester",
	}

	token, _, err := j.CreateToken(claims)
	assert.NoError(t, err)

	refreshed, _, err := j.RefreshToken(token)
	assert.NoError(t, err)

	parsedClaims, err := j.ParseToken(refreshed)
	assert.NoError(t, err)
	assert.Equal(t, customIssuer, parsedClaims.Issuer)
	assert.Equal(t, claims.UserID, parsedClaims.UserID)
	assert.Equal(t, claims.Role, parsedClaims.Role)
}

// 测试 ValidateTokenWithoutExpiration
func TestValidateTokenWithoutExpiration(t *testing.T) {
	j := NewJWT(testSigningKey).WithExpireDuration(time.Millisecond * 100)

	claims := JWTClaims{
		UserID:   999,
		Username: testUsername,
		Role:     "user",
	}

	token, _, _ := j.CreateToken(claims)

	// 等待 Token 过期
	time.Sleep(time.Millisecond * 200)

	// 常规解析应该失败
	_, err := j.ParseToken(token)
	assert.Error(t, err)
	assert.Equal(t, ErrTokenExpired, err)

	// 但忽略过期验证应该成功
	parsedClaims, err := j.ValidateTokenWithoutExpiration(token)
	assert.NoError(t, err)
	assert.Equal(t, claims.UserID, parsedClaims.UserID)
	assert.Equal(t, claims.Username, parsedClaims.Username)
	assert.Equal(t, claims.Role, parsedClaims.Role)
}

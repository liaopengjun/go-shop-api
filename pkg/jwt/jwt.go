package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"go-shop-api/global"
	"time"
)

//签名key
type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	AuthorityId string
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GA_CONFIG.JwtConfig.SigningKey),
	}
}

//定义常量
var (
	TokenExpired     error = errors.New("令牌已过期")
	TokenNotValidYet error = errors.New("令牌尚未激活")
	TokenMalformed   error = errors.New("令牌格式不对")
	TokenInvalid     error = errors.New("无法处理此令牌")
)

// CreateToken 创建Token
func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(global.GA_CONFIG.JwtConfig.BufferTime), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(global.GA_CONFIG.JwtConfig.ExpiresTime), // 过期时间 7天  配置文件
			Issuer:    global.GA_CONFIG.JwtConfig.Issuer,                                 // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

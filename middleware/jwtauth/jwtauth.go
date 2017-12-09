package jwtauth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mkgo/common"
	"mkgo/mklog"
	"net/http"
	"time"
)

const tokenExpireTime = 30 * time.Minute

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "" {
			j := NewJWT()
			claims, err := j.ParseToken(token)
			c.Set("claims", claims)

			if err != nil {
				if err == TokenExpired {
					c.JSON(http.StatusOK, common.NewTokenExpiredResponse())
					c.Abort()
				} else {
					c.JSON(http.StatusUnauthorized, common.NewAuthFailedResponse("Auth failed"))
					c.Abort()
				}
				return
			}
		}
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("token is expired")
	TokenNotValidYet error  = errors.New("token not active yet")
	TokenMalformed   error  = errors.New("that's not even a token")
	TokenInvalid     error  = errors.New("couldn't handle this token")
	SignKey          string = "MKApplication"
)

type CustomClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}
func GetSignKey() string {
	return SignKey
}
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

func GetToken(c *gin.Context, user string) string {
	claims := CustomClaims{}
	claims.User = user
	j := NewJWT()
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(tokenExpireTime).Unix()
	token, err := j.CreateToken(claims)
	if err != nil {
		c.Header("token", token)
	}
	mklog.Logger.Info(token)
	return token
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return claims, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return claims, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return claims, TokenNotValidYet
			} else {
				return claims, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return claims, TokenInvalid
}
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(tokenExpireTime).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
)

const (
	userIDKey    = "id"
	userEmailKey = "email"
)

func AuthMiddleware(key string, authenticator func(c *gin.Context) (interface{}, error)) *jwt.GinJWTMiddleware {
	mw, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(key),
		Timeout:     1000 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: userIDKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("Payload func", data)
			if v, ok := data.(*serializers.UserResp); ok {
				fmt.Println("parse token", v, ok)
				return jwt.MapClaims{
					userIDKey:    v.ID,
					userEmailKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user, err := authenticator(c)
			fmt.Println("is err authenticator", err)
			fmt.Println("user", user)
			if err != nil {
				// return nil, jwt.ErrFailedAuthentication
				return nil, err
			}
			return user, nil
		},
		HTTPStatusMessageFunc: func(err error, c *gin.Context) string {
			return err.Error()
		},
		Unauthorized: func(c *gin.Context, _ int, _ string) {
			c.JSON(http.StatusUnauthorized, serializers.Resp{
				Result: nil,
				Error:  service.ErrInvalidCredentials,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, serializers.Resp{
				Result: serializers.UserLoginResp{
					Token:   token,
					Expired: expire.Format(time.RFC3339),
				},
				Error: nil,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, serializers.Resp{
				Result: serializers.UserLoginResp{
					Token:   token,
					Expired: expire.Format(time.RFC3339),
				},
				Error: nil,
			})
		},
	})
	return mw
}

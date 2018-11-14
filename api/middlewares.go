package api

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/service"
)

const (
	userIDKey    = "id"
	userEmailKey = "email"
	userKindKey  = "kind"
)

func AuthMiddleware(key string, authenticator func(c *gin.Context) (interface{}, error)) *jwt.GinJWTMiddleware {
	mw, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(key),
		Timeout:     1000 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: userIDKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					userIDKey:    v.ID,
					userEmailKey: v.Email,
					userKindKey:  v.Type,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user, err := authenticator(c)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return user, nil
		},
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		//         if v, ok := data.(*models.User); ok && v.Email == "admin" {
		//                 return true
		//         }
		//
		//         return false
		// },
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Data":  nil,
				"Error": service.ErrInvalidCredentials,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"Data": struct {
					Token  string
					Expire string
				}{
					Token:  token,
					Expire: expire.Format(time.RFC3339),
				},
				"Error": nil,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"Data": struct {
					Token  string
					Expire string
				}{
					Token:  token,
					Expire: expire.Format(time.RFC3339),
				},
				"Result": nil,
			})
		},
	})
	return mw
}

package authx

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"

	"github.com/bezhai/multi-bot-task/biz/utils/respx"
	"github.com/bezhai/multi-bot-task/utils/env_utils"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const identityKey = "id"

type User struct {
	UserName string
}

var AuthMiddleware *jwt.HertzJWTMiddleware

func InitAuthMiddleware(ctx context.Context) {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "chiwei bot",
		Key:         []byte(env_utils.Value("HTTP_JWT_KEY")),
		Timeout:     24 * time.Hour,
		MaxRefresh:  6 * 24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals login
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if CheckUserPassword(ctx, userID, password) == nil {
				return &User{
					UserName: userID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*User); ok && v.UserName != "" {
				err = CheckUser(ctx, v.UserName)
				if err != nil {
					return false
				}
				return true
			}

			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			switch message {
			case "Token is expired", "token is expired":
				respx.FailWithCode(c, code, respx.TokenExpire, message)
			default:
				respx.FailWithCode(c, code, respx.TokenFail, message)
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			respx.SuccessWith(c, map[string]interface{}{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			respx.SuccessWith(c, map[string]interface{}{
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := AuthMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("AuthMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}

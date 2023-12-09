package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-admin/pkg/token"
	"go-admin/pkg/user"
	"net/http"
	"strings"
	"time"
)

const (
	UserKey             = "user"
	AuthorizationHeader = "Authorization"
	ExpireDuration      = 24 * time.Hour
)

type AuthMiddleware = func(c *gin.Context)

func AuthHandler(tokenizer token.Tokenizer, userStorage user.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenBearer := c.GetHeader(AuthorizationHeader)
		splitToken := strings.Split(tokenBearer, "Bearer ")
		var tokenStr string
		if len(splitToken) == 2 {
			tokenStr = splitToken[1]
		}
		claim, err := tokenizer.Parse(tokenStr)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if time.Unix(claim.IssuedAt, 0).Add(ExpireDuration).Before(time.Now()) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		u, err := userStorage.GetOne(c, int64(claim.UserID))
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !u.IsActive {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Set(UserKey, u)
		ctx := context.WithValue(c.Request.Context(), UserKey, u)
		c.Request.WithContext(ctx)
		c.Next()
	}
}

func ExtractUserFromCtx(c *gin.Context) (result user.User, existed bool) {
	u, ok := c.Get(UserKey)
	if !ok {
		existed = ok
		return
	}
	result, existed = u.(user.User)
	return
}

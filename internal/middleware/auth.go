package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrzalr/cookshare-go/internal/models"
	"github.com/mrzalr/cookshare-go/pkg/utils/jwt"
)

func (mw *MiddlewareManager) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				models.StatusUnauthorized([]string{"access denied! You need to login first"}),
			)

			return
		}

		authorization := strings.Split(authHeader, " ")
		if authorization[0] != "Bearer" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				models.StatusUnauthorized([]string{"access denied! Invalid token"}),
			)

			return
		}

		claims, err := jwt.VerifyToken(authorization[1])
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				models.StatusUnauthorized([]string{"access denied! Invalid token"}),
			)

			return
		}

		ctx.Set("id", (*claims)["id"])
		ctx.Next()
	}
}

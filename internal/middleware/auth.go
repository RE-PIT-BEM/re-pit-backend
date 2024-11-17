package middleware

import (
	"net/http"
	"strings"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/lib"
	"github.com/gin-gonic/gin"
)

func RequireAuth(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You're not login yet!",
		})
		return
	}

	token := strings.Split(bearerToken, " ")[1]

	claims, err := lib.ParseJWTToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Set("userId", claims["sub"])
	ctx.Next()
}

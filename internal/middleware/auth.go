package middleware

import (
	"be-assignment/domain"
	"be-assignment/dto"
	"be-assignment/internal/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 {
			ctx.AbortWithStatusJSON(401, &dto.Response{
				Message: domain.ErrInvalidToken.Error(),
				Data:    []string{},
				Status:  401,
			})
			return
		}

		tokenString = parts[1]

		if err := util.ValidateToken(tokenString); err != nil {
			ctx.AbortWithStatusJSON(401, &dto.Response{
				Message: err.Error(),
				Data:    []string{},
				Status:  401,
			})
			return
		}

		if claims, err := util.GetClaims(tokenString); err == nil {
			ctx.Set("x-user", claims)
		}

		ctx.Next()
	}
}

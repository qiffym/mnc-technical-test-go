package middleware

import (
	"mncPaymentAPI/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToke(ctx)
		_ = verifyToken

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"messege": err.Error(),
			})
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}

func ValidatorAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Verifikasi token JWT dan ambil claims
		claims, err := helper.VerifyToke(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		role, roleExists := claims["role"].(string)
		if !roleExists || role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not authorized to access this resource",
			})
			return
		}

		// Set user data di context jika authorized
		ctx.Set("userData", claims)
		ctx.Next()
	}
}

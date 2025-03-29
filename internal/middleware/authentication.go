package middleware

import (
	"github.com/FJericho/test-mnc/internal/helper"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		verifiedToken, err := helper.VerifyToken(context)

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Set("userData", verifiedToken)
		context.Next()
	}
}
package middleware

import (
	"gin-model/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Log.Print(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status": -1,
					"msg":    err,
					"data":   nil,
				})
			}
		}()

		ctx.Next()
	}
}

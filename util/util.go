package util

import "github.com/gin-gonic/gin"

func GetReqID(ctx *gin.Context) string {
	value, exists := ctx.Get("X-Request-Id")
	if !exists {
		return ""
	}
	if requestId, ok := value.(string); ok {
		return requestId
	}
	return ""
}

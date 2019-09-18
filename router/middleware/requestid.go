package middleware

import "github.com/gin-gonic/gin"
import "github.com/satori/go.uuid"

const HeaderRequestId = "X-Request-Id"

func RequestId(ctx *gin.Context) {
	requestId := ctx.Request.Header.Get(HeaderRequestId)
	if requestId == "" {
		requestId = uuid.NewV4().String()
	}
	ctx.Set(HeaderRequestId, requestId)
	ctx.Writer.Header().Set(HeaderRequestId, requestId)
	ctx.Next()
}

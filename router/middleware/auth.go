package middleware

import (
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/pkg/token"
	"github.com/fibbery/go-web/router/handler"
	"github.com/gin-gonic/gin"
)

func Token(ctx *gin.Context) {
	if _, e := token.ParseRequest(ctx); e != nil {
		handler.SendResponse(ctx, errno.ErrTokenInvalid, nil)
		ctx.Abort()
		return
	}
	ctx.Next()
}

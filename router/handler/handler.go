package handler

import (
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeError(err)
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

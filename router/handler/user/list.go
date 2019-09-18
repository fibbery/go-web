package user

import (
	"github.com/fibbery/go-web/model"
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/router/handler"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	var r ListRequest
	if err := ctx.Bind(&r); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}
	users, err := model.ListUsers(r.Offset, r.Limit)

	if err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(ctx, errno.OK, users)
}

package user

import (
	"github.com/fibbery/go-web/model"
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/router/handler"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(ctx, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(ctx, errno.OK, user)
}

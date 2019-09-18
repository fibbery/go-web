package user

import (
	"github.com/fibbery/go-web/model"
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/router/handler"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := model.Delete(id); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(ctx, errno.OK, nil)
}

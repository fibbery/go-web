package user

import (
	"github.com/fibbery/go-web/model"
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/router/handler"
	"github.com/fibbery/go-web/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(ctx *gin.Context) {
	log.Info("User Update function called.", lager.Data{"X-Request-Id": util.GetReqID(ctx)})
	userId, _ := strconv.Atoi(ctx.Param("id"))
	var u model.UserModel
	if err := ctx.Bind(&u); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}
	u.Id = userId

	if err := u.Validate(); err != nil {
		handler.SendResponse(ctx, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.ErrEncrypt, nil)
		return
	}

	if err := model.Update(&u); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}

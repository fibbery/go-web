package user

import (
	"github.com/fibbery/go-web/model"
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/router/handler"
	"github.com/fibbery/go-web/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(ctx *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(ctx)})
	var r CreateRequest
	if err := ctx.Bind(&r); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.UserName,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		handler.SendResponse(ctx, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(ctx, nil, CreateResponse{UserName: r.UserName})
}

package user

import (
	"fmt"
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/router/handler"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(ctx *gin.Context) {
	var r CreateRequest
	if err := ctx.Bind(&r); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}
	if r.UserName == "" {
		handler.SendResponse(ctx, errno.New(errno.ErrUserCreateFail, fmt.Errorf("username can't be blank")), nil)
		return
	}
	if r.Password == "" {
		handler.SendResponse(ctx, errno.New(errno.ErrUserCreateFail, fmt.Errorf("password can't be blank")), nil)
		return
	}

	log.Infof("Content-Type : %s", ctx.ContentType())
	log.Infof("[create user] username is %s, password is %s", r.UserName, r.Password)

	handler.SendResponse(ctx, nil, CreateResponse{UserName: r.UserName,})
}

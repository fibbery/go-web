package user

import (
	"github.com/fibbery/go-web/model"
	"github.com/fibbery/go-web/pkg/auth"
	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/pkg/token"
	"github.com/fibbery/go-web/router/handler"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var u model.UserModel
	if err := ctx.Bind(&u); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(ctx, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		handler.SendResponse(ctx, errno.ErrPasswordIncorrect, nil)
		return
	}

	tokenString, err := token.Sign(&token.Context{Id: d.Id, Username: d.Username}, "")
	if err != nil {
		handler.SendResponse(ctx, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(ctx, nil, model.Token{Token: tokenString})
}

package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

type Context struct {
	Id       int
	Username string
}

func Sign(c *Context, secrect string) (tokenString string, err error) {
	if secrect == "" {
		secrect = viper.GetString("jwt_secrect")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.Id,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err = token.SignedString([]byte(secrect))

	_, err = Parse(tokenString, secrect)
	if err != nil {
		log.Errorf(err, "jwt error")
	}
	return
}

func ParseRequest(ctx *gin.Context) (*Context, error) {
	header := ctx.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}
	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, viper.GetString("jwt_secrect"))
}

func Parse(tokenString string, secrect string) (*Context, error) {
	ctx := &Context{}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (i interface{}, e error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secrect), nil
	})

	if err != nil {
		return ctx, err
	}

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Id = (int)(claim["id"].(float64))
		ctx.Username = claim["username"].(string)
	}

	return ctx, err
}

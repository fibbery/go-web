package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/fibbery/go-web/pkg/errno"
	"github.com/fibbery/go-web/router/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r BodyLogWriter) Write(data []byte) (int, error) {
	r.body.Write(data)
	return r.ResponseWriter.Write(data)
}

func Log(ctx *gin.Context) {
	start := time.Now().UTC()
	path := ctx.Request.URL.Path

	//排除系统检查的
	if path == "/sd/health" || path == "/sd/disk" || path == "/sd/cpu" || path == "/sd/ram" {
		return
	}

	var body []byte
	if ctx.Request.Body != nil {
		body, _ = ioutil.ReadAll(ctx.Request.Body)
		//restore body data to request
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	}

	ip := ctx.ClientIP()
	method := ctx.Request.Method

	//reset ctx.writer to recrod response
	blw := &BodyLogWriter{
		ResponseWriter: ctx.Writer,
		body:           bytes.NewBufferString(""),
	}
	ctx.Writer = blw

	//execute method
	ctx.Next()

	cost := time.Now().UTC().Sub(start)
	code, message := -1, ""

	var response handler.Response
	if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
		log.Errorf(err, "unmarshal response body error, this body is %s", string(blw.body.Bytes()))
		code = errno.InternalServerError.Code
		message = err.Error()
	} else {
		code = response.Code
		message = response.Message
	}

	log.Infof("%-13s | %-12s | %s %s | request : %s | response : {code: %d, message: %s}", cost, ip, pad.Right(method, 5, ""), path, string(body), code, message)
}

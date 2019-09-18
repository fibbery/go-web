package main

import (
	"errors"
	"github.com/fibbery/go-web/config"
	"github.com/fibbery/go-web/model"
	"github.com/fibbery/go-web/router"
	"github.com/fibbery/go-web/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	c = pflag.StringP("config", "c", "", "config file absolute path")
)

func main() {
	pflag.Parse()
	if err := config.Init(*c); err != nil {
		panic(err)
	}

	// inint db
	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	mdw := []gin.HandlerFunc{middleware.RequestId, middleware.Log}

	router.Load(g, mdw...)
	go func() {
		if err := pingServer(); err != nil {
			panic(err)
		}
		log.Info("server healthcheck successfully!!!")
	}()

	log.Infof("server start at address %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_retry_count"); i++ {
		resp, err := http.Get(viper.GetString("uri") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		delay := 1 << uint(i)
		log.Infof("will retry in %d seconds to healthcheck", delay)
		time.Sleep(time.Duration(delay) * time.Second)
	}
	return errors.New("health check error, please check the server")
}

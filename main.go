package main

import (
	"errors"
	"github.com/fibbery/go-web/config"
	"github.com/fibbery/go-web/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
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

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	mdw := []gin.HandlerFunc{}

	router.Load(g, mdw...)
	go func() {
		if err := pingServer(); err != nil {
			panic(err)
		}
		log.Println("server healthcheck successfully!!!")
	}()

	log.Printf("server start at address %s", viper.GetString("addr"))
	log.Fatal(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_retry_count"); i++ {
		resp, err := http.Get(viper.GetString("uri")+ "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		delay := 1 << uint(i)
		log.Printf("will retry in %d seconds to healthcheck", delay)
		time.Sleep(time.Duration(delay) * time.Second)
	}
	return errors.New("health check error, please check the server")
}

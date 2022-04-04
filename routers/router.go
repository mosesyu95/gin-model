package routers

import (
	"gin-model/config"
	l "gin-model/log"
	"gin-model/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

func CollectRoute(router *gin.Engine) *gin.Engine {

	// if nor find router, will redirect to /crocodile/
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	return router
}

func Init() {
	r := gin.Default()
	gin.SetMode(viper.GetString("server.env")) // 设置运行环境debug  release  test
	gin.DefaultWriter = io.MultiWriter(l.Log.Out, os.Stdout)
	r.Use(middleware.LoggerWithFormatter(), middleware.RecoveryMiddleware(), middleware.CORSMiddleware())
	r = CollectRoute(r)
	port := config.Config.Port
	if port != 0 {
		panic(r.Run(fmt.Sprintf("0.0.0.0:%d", port)))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

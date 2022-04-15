package routers

import (
	"fmt"
	"gin-model/config"
	_ "gin-model/docs"
	l "gin-model/log"
	"gin-model/middleware"
	"gin-model/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"io"
	"net/http"
	"os"
)

func CollectRoute(router *gin.Engine) *gin.Engine {

	router.Use(middleware.LoggerWithFormatter(), middleware.RecoveryMiddleware(), middleware.CORSMiddleware())
	router.GET("/info", api.Info)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/info")
	})
	return router
}

func Init() {
	r := gin.Default()
	gin.SetMode(viper.GetString("server.env")) // 设置运行环境debug  release  test
	gin.DefaultWriter = io.MultiWriter(l.Log.Out, os.Stdout)
	r = CollectRoute(r)
	port := config.Config.Port
	if port != 0 {
		panic(r.Run(fmt.Sprintf("0.0.0.0:%d", port)))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

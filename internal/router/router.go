package router

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-gpt/config"
	"go-gpt/internal/middleware"
)

func SetRouters() *gin.Engine {
	var r *gin.Engine

	if strings.ToUpper(config.GLOBAL_CONFIG.ENV) == "PRODUCTION" {
		// 生产模式
		r = ReleaseRouter()
		r.Use(
			middleware.RequestCostHandler(),
			middleware.CustomLogger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	} else {
		// 开发调试模式
		r = gin.New()
		r.Use(
			middleware.RequestCostHandler(),
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	}
	// set up trusted agents
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	// ping
	r.Any("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	// 设置 API 路由
	initRoute(r)

	return r
}

// ReleaseRouter 生产模式使用官方建议设置为 release 模式
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式
	gin.SetMode(gin.ReleaseMode)
	// 禁用 gin 输出接口访问日志
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()

	return engine
}

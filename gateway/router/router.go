package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkyian78/gateway/controller"
	"github.com/rizkyian78/gateway/utils"
)

func Router(h *controller.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(utils.LoggerRequest))
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("notfound")

	router.GET("/ping", h.Ping)
	router.POST("/auth/token", h.OauthToken)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "notfound.html", gin.H{})
	})
	return router
}

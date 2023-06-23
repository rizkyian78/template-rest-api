package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rizkyian78/gateway/service"
	"github.com/rizkyian78/gateway/utils"
)

func Router(h *controller.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(utils.LoggerRequest))
	router.Use(gin.Recovery())
	router.Static("/css", "../view/style.css")
	router.LoadHTMLGlob("view/*")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "notfound.html", gin.H{})
	})

	router.GET("/ping", h.Ping)
	router.POST("/auth/token", h.OauthToken)

	return router
}

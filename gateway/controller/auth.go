package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
)

func (h *Handler) OauthToken(c *gin.Context) {

	trace, _ := opentracing.StartSpanFromContext(c.Request.Context(), "Handling POST /auth/token")

	defer trace.Finish()
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi, from token",
	})

}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/rizkyian78/gateway/utils"
)

func (h *Handler) OauthToken(c *gin.Context) {

	trace, _ := opentracing.StartSpanFromContext(c.Request.Context(), "Handling POST /auth/token")
	defer trace.Finish()

	token, duration, err := utils.GenerateToken(&utils.JWTClaims{
		Username: "rizkyian78",
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "succes",
		"token":      token,
		"type":       "Bearer",
		"scope":      "auth_token",
		"expired_in": duration,
	})

}

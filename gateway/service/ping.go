package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/rizkyian78/gateway/queue/rabbitmq"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	producer rabbitmq.Producer
}

func NewHandler(producer *rabbitmq.Producer) (*Handler, error) {
	handler := &Handler{
		producer: *producer,
	}
	return handler, nil
}

func (h *Handler) Ping(c *gin.Context) {
	log := logrus.WithFields(logrus.Fields{
		"service": "Gateway Service",
	})
	trace, _ := opentracing.StartSpanFromContext(c.Request.Context(), "Handling Get /ping")
	log.WithTime(time.Now()).Info("Gateway is processing")

	defer trace.Finish()

	c.JSON(http.StatusOK, gin.H{
		"message": "Hi, from ping",
	})
}

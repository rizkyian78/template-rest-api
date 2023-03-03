package controller

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/rizkyian78/deployment/queue/rabbitmq"
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

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"service": "Gateway Service",
	})
	trace, ctx := opentracing.StartSpanFromContext(r.Context(), "Handling Get /ping")
	fmt.Println(ctx)
	log.WithTime(time.Now()).Info("Gateway is processing")

	rByte, err := json.Marshal(map[string]string{
		"asdasd": "asds",
	})

	if err != nil {
		log.Error(err)
		return
	}

	rString := string(rByte)

	if err := h.producer.PublishMessage(rString); err != nil {
		log.WithError(err).Error("error in publish message")
	}

	defer trace.Finish()

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

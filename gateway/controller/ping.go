package controller

import (
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"service": "Gateway Service",
	})
	trace, ctx := opentracing.StartSpanFromContext(r.Context(), "Handling Get /ping")
	fmt.Println(ctx)
	log.WithTime(time.Now()).Info("Gateway is processing")

	defer trace.Finish()

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

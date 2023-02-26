package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-lib/metrics"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

const (
	serviceName = "Gateway Service"
)

func main() {

	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,

		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	fmt.Println(err)
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()

		span := tracer.StartSpan("say-hello")
		println("adasdas")
		span.Finish()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on localhost:8181")

	log.Fatal(http.ListenAndServe(":8181", nil))
}

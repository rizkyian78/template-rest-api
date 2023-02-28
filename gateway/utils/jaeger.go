package utils

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func JaegerTracing(serviceName string, useLogging bool, log *logrus.Entry) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: useLogging,
		},
	}

	closer, err := cfg.InitGlobalTracer(
		serviceName,
	)

	if err != nil {
		return nil, err
	}

	return closer, nil
}

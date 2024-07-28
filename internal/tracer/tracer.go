package tracer

import (
	"context"
	"sync"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	traceconfig "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	logger "gitlab.ozon.dev/berkinv/homework/internal/handlers/log"
)

func MustSetup(ctx context.Context, serviceName string, swap bool, brokers []string, topic string) {
	cfg := traceconfig.Configuration{
		ServiceName: serviceName,
		Sampler: &traceconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &traceconfig.ReporterConfig{},
	}
	log := logger.NewLogger(swap, brokers, topic)
	tracer, closer, err := cfg.NewTracer(traceconfig.Logger(jaeger.StdLogger), traceconfig.Metrics(prometheus.New()))
	if err != nil {
		log.Input("ERROR: cannot init Jaeger %s", nil)
	}

	go func() {
		onceCloser := sync.OnceFunc(func() {
			log.Input("closing tracer", nil)
			if err = closer.Close(); err != nil {
				log.Input("error closing tracer", nil)
			}
		})

		for {
			<-ctx.Done()
			onceCloser()
		}
	}()

	opentracing.SetGlobalTracer(tracer)
}

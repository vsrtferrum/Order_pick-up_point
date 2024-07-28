package http

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.ozon.dev/berkinv/homework/internal/logger"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

type (
	positionUpserter interface {
		UpsertPositions(ctx context.Context, positions []*models.DataUnitJson) error
	}
)

func MustRun(
	ctx context.Context,
	shutdownDur time.Duration,
	addr string,
	positionUpserter positionUpserter,
) {
	handler := &Handler{
		positionUpserter: positionUpserter,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/positions", handler.IssueCnt)
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{Addr: addr, Handler: mux}

	go func() {
		<-ctx.Done()

		logger.Errorf(ctx, "Shutting down server with duration %0.3fs", shutdownDur.Seconds())
		<-time.After(shutdownDur)

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Errorf(ctx, "HTTP handler Shutdown: %s", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		logger.Errorf(ctx, "HTTP server ListenAndServe: %s", err)
	}

}

type Handler struct {
	positionUpserter positionUpserter
}

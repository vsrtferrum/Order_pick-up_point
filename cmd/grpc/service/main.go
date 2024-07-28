package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/berkinv/homework/internal/api"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/cli"
	http2 "gitlab.ozon.dev/berkinv/homework/internal/http"
	"gitlab.ozon.dev/berkinv/homework/internal/imdb"
	"gitlab.ozon.dev/berkinv/homework/internal/middleware"
	"gitlab.ozon.dev/berkinv/homework/internal/tracer"
	"gitlab.ozon.dev/berkinv/homework/pkg/api/proto/pvz/v1/pvz/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	fulldbreq string
	mode      string
	brokers   = []string{
		"127.0.0.1:9091",
	}
	topic            string = "Log"
	grpcPort                = 50051
	httpPort                = ":63342"
	applicationPort         = ":8099"
	shutdownDuration        = 5 * time.Second
	limit                   = 5000
)

func main() {
	flag.StringVar(&fulldbreq, "fulldbreq", "user=vsrtf dbname=postgres  sslmode=disable", "fulldbreq")
	flag.StringVar(&mode, "mode", "kafka", "mode")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	client := cli.NewCLI(brokers, topic, handlers.IsKafkaMode(mode))

	err := client.Keepnthrow.Storage.SetFullDatabaseReq(fulldbreq, limit)
	if err != nil {
		panic(err)
	}

	pvzService := &api.PvzService{}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.Logging))
	pvz.RegisterPvzServer(grpcServer, pvzService.PvzServ)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = pvz.RegisterPvzHandlerFromEndpoint(ctx, mux, fulldbreq, opts)
	if err != nil {
		panic(err)
		return
	}

	go func() {
		gwServer := &http.Server{ // Создаем HTTP gateway сервер
			Addr:    httpPort,
			Handler: middleware.WithHTTPLoggingMiddleware(mux), // middleware
		}

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		errHttp := gwServer.ListenAndServe()
		if errHttp != nil {
			log.Println(errHttp)
			return
		}

	}()

	positionIMDB := imdb.NewRepository()

	http2.MustRun(ctx, shutdownDuration, applicationPort, positionIMDB)
	if err = grpcServer.Serve(lis); err != nil { // запускаем grpc сервер
		log.Fatalf("failed to serve: %v", err)
	}
	tracer.MustSetup(ctx, "Tracer", handlers.IsKafkaMode(mode), brokers, topic)
	log.Println("Done")
}

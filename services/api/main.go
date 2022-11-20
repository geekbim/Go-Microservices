package main

import (
	"context"
	"go-grpc-example/pkg/logger"
	"go-grpc-example/services/api/internal/config"
	"go-grpc-example/services/api/internal/delivery/grpc_server"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "X-User-Id":
		return key, true
	case "X-User-Role":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

var appLogger = logger.NewApiLogger(nil)

func main() {
	cfg, multierr := config.Api()
	if multierr != nil {
		log.Fatalln(multierr)
	}

	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	appLogger.Info("Serving gRPC on " + cfg.GrpcPort)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gw := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpc_server.NewAuthHandler(context.Background(), gw, opts, cfg)
	grpc_server.NewUserHandler(context.Background(), gw, opts, cfg)

	gwServer := &http.Server{
		Addr:    cfg.AppPort,
		Handler: gw,
	}

	appLogger.Info("Serving gRPC-Gateway on " + cfg.AppPort)
	log.Fatalln(gwServer.ListenAndServe())
}

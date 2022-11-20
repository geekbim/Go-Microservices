package main

import (
	"context"
	"fmt"
	"go-grpc-example/pkg/logger"
	"go-grpc-example/services/auth/internal/config"
	"go-grpc-example/services/auth/internal/delivery/grpc_server/grpc_auth"
	auth_repository "go-grpc-example/services/auth/internal/repository/psql/auth"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

var (
	appLogger = logger.NewApiLogger(nil)
	cfg       = config.Auth()
	db        = config.InitDatabase()
	authRepo  = auth_repository.NewAuthRepository(db)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	grpc_auth.HandlerAuthService(s, &cfg, authRepo)

	go func() {
		appLogger.Info(fmt.Sprintf("Server is listening on port: %v", cfg.Port))
		appLogger.Fatal(s.Serve(lis))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
		cancel()
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
		cancel()
	}

	s.GracefulStop()
	appLogger.Info("Server Exited Properly")
}

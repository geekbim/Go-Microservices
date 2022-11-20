package main

import (
	"context"
	"fmt"
	"go-grpc-example/services/user/internal/config"
	"go-grpc-example/services/user/internal/delivery/events"
	"go-grpc-example/services/user/internal/delivery/grpc_client/grpc_auth"
	"go-grpc-example/services/user/internal/delivery/grpc_server/grpc_user"
	psql_user "go-grpc-example/services/user/internal/repository/psql/user"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go-grpc-example/pkg/logger"

	"google.golang.org/grpc"
)

var (
	cfg            = config.User()
	ctx, cancel    = context.WithCancel(context.Background())
	appLogger      = logger.NewApiLogger(nil)
	db             = config.InitDatabase()
	userRepository = psql_user.NewUserRepository(db)
)

func main() {
	connAuth, err := grpc.Dial(cfg.GrpcClient.Auth, grpc.WithInsecure())
	if err != nil {
		cancel()
		appLogger.Fatal(err)
	}

	authGrpcClient, err := grpc_auth.NewAuthHandler(ctx, connAuth)
	if err != nil {
		cancel()
		appLogger.Fatal(err)
	}
	appLogger.Info("grpc auth started")

	events.EventHandler(authGrpcClient, appLogger, db)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		appLogger.Fatal(err)
	}

	s := grpc.NewServer()

	grpc_user.HandlerUserService(s, appLogger, cfg, userRepository)

	go func() {
		appLogger.Info(fmt.Sprintf("Server is listening on port: %v", cfg.Port))
		log.Fatal(s.Serve(lis))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
	}

	s.GracefulStop()
	appLogger.Info("Server Exited Properly")
}

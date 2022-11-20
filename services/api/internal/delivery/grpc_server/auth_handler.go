package grpc_server

import (
	"context"
	"fmt"
	authPb "go-grpc-example/proto/_generated/auth"
	"go-grpc-example/services/api/internal/config"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewAuthHandler(ctx context.Context, gw *runtime.ServeMux, opts []grpc.DialOption, cfg *config.ApiConfig) {
	if err := authPb.RegisterAuthServiceHandlerFromEndpoint(ctx, gw, cfg.ServicePort.Auth, opts); err != nil {
		log.Fatalln(fmt.Sprintf("Failed to register service %s : ", "auth"), err)
	}
}

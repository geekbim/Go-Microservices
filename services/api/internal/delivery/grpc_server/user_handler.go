package grpc_server

import (
	"context"
	"fmt"
	userPb "go-grpc-example/proto/_generated/user"
	"go-grpc-example/services/api/internal/config"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewUserHandler(ctx context.Context, gw *runtime.ServeMux, opts []grpc.DialOption, cfg *config.ApiConfig) {
	if err := userPb.RegisterUserServiceHandlerFromEndpoint(ctx, gw, cfg.ServicePort.User, opts); err != nil {
		log.Fatalln(fmt.Sprintf("Failed to register service %s : ", "user"), err)
	}
}

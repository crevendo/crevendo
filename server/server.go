package server

import (
	"fmt"
	"github.com/crevendo/crevendo/hook"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

type HooksServerOption func(*HooksServer)
type HooksServer struct {
	GrpcServer  *grpc.Server
	TcpListener net.Listener
}

func NewHooksServer(opts ...HooksServerOption) {
	var server HooksServer
	var err error

	_ = godotenv.Load()
	server.TcpListener, err = net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	unaryInterceptor := grpc.UnaryInterceptor(grpc_recovery.UnaryServerInterceptor())
	server.GrpcServer = grpc.NewServer(unaryInterceptor)

	if os.Getenv("DEV") == "true" {
		fmt.Println("Server is in DEV mode")
		server.GrpcServer = grpc.NewServer()
	}

	for _, opt := range opts {
		opt(&server)
	}
	reflection.Register(server.GrpcServer)
	fmt.Println("Hooks Server Running On Port " + os.Getenv("PORT") + "...")
	err = server.GrpcServer.Serve(server.TcpListener)
	if err != nil {
		panic("error grave")
	}
}

func WithGatewayHooks(hooks hook.GatewayHooksServer) HooksServerOption {
	return func(server *HooksServer) {
		hook.RegisterGatewayHooksServer(server.GrpcServer, hooks)
	}
}

func WithPaymentMethodsHooks(hooks hook.PaymentHooksServer) HooksServerOption {
	return func(server *HooksServer) {
		hook.RegisterPaymentHooksServer(server.GrpcServer, hooks)
	}
}

func WithProductHooks(hooks hook.ProductHooksServer) HooksServerOption {
	return func(server *HooksServer) {
		hook.RegisterProductHooksServer(server.GrpcServer, hooks)
	}
}

func WithCategoryHooks(hooks hook.CategoryHooksServer) HooksServerOption {
	return func(server *HooksServer) {
		hook.RegisterCategoryHooksServer(server.GrpcServer, hooks)
	}
}

func WithCartHooks(hooks hook.CartHooksServer) HooksServerOption {
	return func(server *HooksServer) {
		hook.RegisterCartHooksServer(server.GrpcServer, hooks)
	}
}

func WithOrderHooks(hooks hook.OrderHooksServer) HooksServerOption {
	return func(server *HooksServer) {
		hook.RegisterOrderHooksServer(server.GrpcServer, hooks)
	}
}

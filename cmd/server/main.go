package main

import (
	"fmt"
	"log"
	"net"

	internalGrpc "github.com/pepeunlimited/go-grpc-starter-kit/internal/server/grpc"
	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/api/v1/services"
	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/env"
	"google.golang.org/grpc"

	"go.temporal.io/sdk/client"
)

const (
	Version string = "0.0.1"
	Port    int    = 8080
	PortEnv string = "GO_GRPC_STARTER_KIT_PORT"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("🚀 Started the GoGrpcStarterKit, version=[%s]", Version)
	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Panicf("❌ temporal.client.NewClient occurred issue: %v", err)
	}
	// Todo & Temporal gRPC services
	todoServer := internalGrpc.NewTodoServer()
	temporalServer := internalGrpc.NewTemporalServer(temporalClient)

	// tcp & gRPC server
	goGrpcStarterKitPort := env.IntEnv(PortEnv, Port)
	conn, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", goGrpcStarterKitPort))
	if err != nil {
		log.Fatalln("❌ net.Listen occurred issue: ", err)
	}
	grpcServer := grpc.NewServer()

	// register gRPC services
	services.RegisterTodoServiceServer(grpcServer, todoServer)
	services.RegisterTemporalServiceServer(grpcServer, temporalServer)

	// serve
	if grpcServer.Serve(conn); err != nil {
		log.Fatalln("❌ grpcServer.Serve occurred issue: ", err)
	}

}

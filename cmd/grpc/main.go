package main

import (
	"fmt"
	"github.com/marmota-alpina/orders-service/internal/config"
	"github.com/marmota-alpina/orders-service/internal/db"
	"github.com/marmota-alpina/orders-service/internal/order"
	"github.com/marmota-alpina/orders-service/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// Carregar configuração
	cfg := config.LoadConfig()

	// Obter banco de dados
	database, err := db.GetDatabase()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Criar repositório e serviço
	repo := order.NewRepository(database)
	orderService := order.NewOrderService(repo)

	// Criar servidor gRPC
	grpcServer := grpc.NewServer()

	// Registrar o serviço gRPC
	proto.RegisterOrderServiceServer(grpcServer, orderService)

	// Registrar reflexão (para depuração)
	reflection.Register(grpcServer)

	// Definir o endereço de escuta
	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Iniciar servidor gRPC
	log.Printf("Starting gRPC server on %s...", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

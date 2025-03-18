package server

import (
	"fmt"
	"github.com/marmota-alpina/orders-service/internal/config"
	"github.com/marmota-alpina/orders-service/internal/order"
	"github.com/marmota-alpina/orders-service/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// Server estrutura que representa o servidor gRPC
type Server struct {
	grpcServer *grpc.Server
}

// NewServer cria e retorna uma nova instância do servidor gRPC
func NewServer(cfg *config.Config, repo *order.Repository, orderService proto.OrderServiceServer) *Server {
	// Criação do servidor gRPC
	grpcServer := grpc.NewServer()

	// Registra o serviço OrderService no servidor gRPC
	proto.RegisterOrderServiceServer(grpcServer, orderService)

	// Adiciona reflexão (opcional, útil para debugging e ferramentas como grpcurl)
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
	}
}

// Start inicia o servidor gRPC
func (s *Server) Start(cfg *config.Config) {
	// Define o endereço e porta do servidor
	address := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)

	// Cria o listener para o servidor gRPC
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", address, err)
	}

	// Inicia o servidor gRPC
	log.Printf("Starting gRPC server on %s...", address)
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

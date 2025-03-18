package order

import (
	"context"
	"github.com/marmota-alpina/orders-service/internal/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

const grpcServerAddr = "localhost:5001"

// Estabelece a conexão com o servidor gRPC real e cria o cliente
func newGrpcClient(t *testing.T) proto.OrderServiceClient {
	// Estabelecendo conexão com gRPC sem segurança (insecure) para testes locais
	conn, err := grpc.DialContext(
		context.Background(),
		grpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Para conexões sem TLS (para testes)
		grpc.WithBlock(), // Garantir que a conexão esteja pronta antes de continuar
	)
	if err != nil {
		t.Fatalf("could not connect to gRPC server: %v", err)
	}
	t.Cleanup(func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close gRPC connection: %v", err)
		}
	})

	return proto.NewOrderServiceClient(conn)
}

func TestListOrdersIntegration(t *testing.T) {
	client := newGrpcClient(t)

	// Chama o método ListOrders
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.ListOrders(ctx, &proto.Empty{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Orders), 0, "Expected at least one order")
}

func TestGetOrderByIdIntegration(t *testing.T) {
	client := newGrpcClient(t)

	// Suponha que o ID do pedido existente seja 1
	resp, err := client.GetOrderById(context.Background(), &proto.OrderRequest{Id: 3})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int32(3), resp.Id, "Expected order ID to be 1")
	assert.NotEmpty(t, resp.CustomerName, "Expected customer name to be non-empty")
	assert.NotEmpty(t, resp.CreatedAt, "Expected order creation date to be non-empty")
}

func TestCreateOrderIntegration(t *testing.T) {
	client := newGrpcClient(t)

	// Cria um novo pedido
	newOrder := &proto.CreateOrderRequest{
		CustomerName: "John Doe",
		TotalAmount:  123.45,
	}

	resp, err := client.CreateOrder(context.Background(), newOrder)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John Doe", resp.CustomerName, "Expected customer name to be 'John Doe'")
	assert.Equal(t, 123.45, resp.TotalAmount, "Expected total amount to be 123.45")
}

func TestUpdateOrderIntegration(t *testing.T) {
	client := newGrpcClient(t)

	// Atualiza um pedido existente
	updateOrder := &proto.UpdateOrderRequest{
		Id:           3, // Supondo que o ID do pedido seja 1
		CustomerName: "John Doe Updated",
		TotalAmount:  200.75,
	}

	resp, err := client.UpdateOrder(context.Background(), updateOrder)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John Doe Updated", resp.CustomerName, "Expected updated customer name")
	assert.Equal(t, 200.75, resp.TotalAmount, "Expected updated total amount")
}

func TestDeleteOrderIntegration(t *testing.T) {
	client := newGrpcClient(t)

	// Deleta um pedido existente
	resp, err := client.DeleteOrder(context.Background(), &proto.DeleteOrderRequest{Id: 4}) // Supondo que o ID do pedido seja 1
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

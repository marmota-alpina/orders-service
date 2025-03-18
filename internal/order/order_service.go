package order

import (
	"context"
	"database/sql"
	"errors"
	"github.com/marmota-alpina/orders-service/internal/proto"
	"time"
)

type Service struct {
	proto.UnimplementedOrderServiceServer
	repo *Repository
}

func NewOrderService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListOrders(ctx context.Context, req *proto.Empty) (*proto.OrderList, error) {
	orders, err := s.repo.ListOrders()
	if err != nil {
		return nil, err
	}

	var orderList []*proto.Order
	for _, o := range orders {
		orderList = append(orderList, &proto.Order{
			Id:           int32(o.ID),
			CustomerName: o.CustomerName,
			TotalAmount:  o.TotalAmount,
			CreatedAt:    o.CreatedAt.Format(time.RFC3339),
		})
	}

	return &proto.OrderList{Orders: orderList}, nil
}

func (s *Service) GetOrderById(ctx context.Context, req *proto.OrderRequest) (*proto.Order, error) {
	order, err := s.repo.GetOrderById(int(req.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return &proto.Order{
		Id:           int32(order.ID),
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
		CreatedAt:    order.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *Service) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.Order, error) {
	order := &Order{
		CustomerName: req.CustomerName,
		TotalAmount:  req.TotalAmount,
	}

	err := s.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return &proto.Order{
		Id:           int32(order.ID),
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
		CreatedAt:    order.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *proto.UpdateOrderRequest) (*proto.Order, error) {
	order := &Order{
		ID:           int(req.Id),
		CustomerName: req.CustomerName,
		TotalAmount:  req.TotalAmount,
	}

	err := s.repo.UpdateOrder(order)
	if err != nil {
		return nil, err
	}

	return &proto.Order{
		Id:           int32(order.ID),
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
	}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *proto.DeleteOrderRequest) (*proto.Empty, error) {
	err := s.repo.DeleteOrder(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

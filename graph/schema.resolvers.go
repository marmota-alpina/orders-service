package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.68

import (
	"context"
	"fmt"
	"github.com/marmota-alpina/orders-service/internal/order"
	"time"

	"github.com/marmota-alpina/orders-service/graph/model"
)

// ListOrders is the resolver for the ListOrders field.
func (r *queryResolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	repo := order.NewRepository(r.DB)
	orders, err := repo.ListOrders()
	if err != nil {
		return nil, err
	}

	// Convertendo a lista de pedidos do repositório para o formato GraphQL
	var result []*model.Order
	for _, o := range orders {
		result = append(result, &model.Order{
			ID:           fmt.Sprintf("%d", o.ID), // Convertendo int para string
			CustomerName: o.CustomerName,
			TotalAmount:  o.TotalAmount,
			CreatedAt:    o.CreatedAt.Format(time.RFC3339), // Formatando data para string ISO 8601
		})
	}

	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

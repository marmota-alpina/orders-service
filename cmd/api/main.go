package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/marmota-alpina/orders-service/internal/db"
	"github.com/marmota-alpina/orders-service/internal/order"
	"log"
)

func main() {
	database, err := db.GetDatabase()

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database)

	repo := order.NewRepository(database)
	handler := order.NewHandler(repo)

	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:     0,
		MinLength: 2048,
	}))
	e.GET("/orders", handler.ListOrders)
	e.GET("/orders/:id", handler.GetOrderById)
	e.POST("/orders", handler.CreateOrder)
	e.PUT("/orders/:id", handler.UpdateOrder)
	e.DELETE("/orders/:id", handler.DeleteOrder)

	log.Println("Server running on :8080")
	err = e.Start(":8080")
	if err != nil {
		panic(err)
	}
}

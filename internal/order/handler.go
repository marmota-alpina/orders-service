package order

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Handler representa o manipulador de requisições HTTP para orders.
type Handler struct {
	repo *Repository
}

// NewHandler cria um novo Handler para orders.
func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

// ListOrders lista todos os pedidos.
func (h *Handler) ListOrders(c echo.Context) error {
	orders, err := h.repo.ListOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, orders)
}

// GetOrderById busca um pedido pelo ID.
func (h *Handler) GetOrderById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order ID"})
	}

	order, err := h.repo.GetOrderById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, order)
}

// CreateOrder cria um novo pedido.
func (h *Handler) CreateOrder(c echo.Context) error {
	var order Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	err := h.repo.CreateOrder(&order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, order)
}

// UpdateOrder atualiza um pedido existente.
func (h *Handler) UpdateOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order ID"})
	}

	var order Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	order.ID = id
	err = h.repo.UpdateOrder(&order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "order updated"})
}

// DeleteOrder remove um pedido pelo ID.
func (h *Handler) DeleteOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order ID"})
	}

	err = h.repo.DeleteOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "order deleted"})
}

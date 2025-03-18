package order

import (
	"database/sql"
	"errors"
	"log"
)

// Repository representa o repositório de pedidos (Orders)
type Repository struct {
	db *sql.DB
}

// NewRepository cria uma nova instância do repositório de pedidos
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ListOrders retorna todos os pedidos
func (r *Repository) ListOrders() ([]Order, error) {
	rows, err := r.db.Query("SELECT id, customer_name, total_amount, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.TotalAmount, &o.CreatedAt); err != nil {
			log.Println("Error:", err)
			continue
		}
		orders = append(orders, o)
	}
	return orders, nil
}

// GetOrderById retorna um pedido pelo ID
func (r *Repository) GetOrderById(id int) (*Order, error) {
	var o Order
	err := r.db.QueryRow("SELECT id, customer_name, total_amount, created_at FROM orders WHERE id = $1", id).
		Scan(&o.ID, &o.CustomerName, &o.TotalAmount, &o.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &o, nil
}

// CreateOrder insere um novo pedido no banco de dados
func (r *Repository) CreateOrder(o *Order) error {
	query := "INSERT INTO orders (customer_name, total_amount) VALUES ($1, $2) RETURNING id, created_at"
	err := r.db.QueryRow(query, o.CustomerName, o.TotalAmount).Scan(&o.ID, &o.CreatedAt)
	return err
}

// UpdateOrder atualiza os dados de um pedido existente
func (r *Repository) UpdateOrder(o *Order) error {
	query := "UPDATE orders SET customer_name = $1, total_amount = $2 WHERE id = $3"
	res, err := r.db.Exec(query, o.CustomerName, o.TotalAmount, o.ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

// DeleteOrder remove um pedido pelo ID
func (r *Repository) DeleteOrder(id int) error {
	query := "DELETE FROM orders WHERE id = $1"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

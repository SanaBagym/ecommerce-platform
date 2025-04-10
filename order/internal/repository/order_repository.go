package repository

import (
	"fmt"
	"order/internal/domain"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	DB *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(order *domain.Order) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	query := `INSERT INTO orders (customer_name, status, created_at, is_deleted) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(query, order.CustomerName, order.Status, order.Created_at, false).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.OrderItems {
		_, err := tx.Exec(`INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`,
			order.ID, item.ProductID, item.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetAll() ([]domain.Order, error) {
	var orders []domain.Order
	// Ensure you're checking if the query returns any rows and handle possible errors
	err := r.DB.Select(&orders, `SELECT * FROM orders WHERE is_deleted = false`)
	if err != nil {
		// Log the error or return a more detailed message
		return nil, fmt.Errorf("error fetching orders: %w", err)
	}

	// Now, fetch associated order items for each order
	for i := range orders {
		var items []domain.OrderItem
		err = r.DB.Select(&items, `SELECT * FROM order_items WHERE order_id = $1`, orders[i].ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching order items for order ID %d: %w", orders[i].ID, err)
		}
		orders[i].OrderItems = items
	}

	return orders, nil
}

func (r *OrderRepository) GetByID(id int) (*domain.Order, error) {
	var order domain.Order
	err := r.DB.Get(&order, `SELECT * FROM orders WHERE id=$1 AND is_deleted = false`, id)
	if err != nil {
		return nil, err
	}

	var items []domain.OrderItem
	err = r.DB.Select(&items, `SELECT id, order_id, product_id, quantity FROM order_items WHERE order_id = $1`, id)
	if err != nil {
		return nil, err
	}

	// Assign the fetched items to the Order
	order.OrderItems = items

	return &order, nil
}

func (r *OrderRepository) Patch(id int, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	query := "UPDATE orders SET "
	params := []interface{}{}
	i := 1

	for k, v := range fields {
		query += fmt.Sprintf("%s = $%d, ", k, i)
		params = append(params, v)
		i++
	}
	query = query[:len(query)-2] // remove last comma
	query += fmt.Sprintf(" WHERE id = $%d", i)
	params = append(params, id)

	_, err := r.DB.Exec(query, params...)
	return err
}

func (r *OrderRepository) Delete(id int) error {
	_, err := r.DB.Exec(`UPDATE orders SET is_deleted = true WHERE id = $1`, id)
	return err
}

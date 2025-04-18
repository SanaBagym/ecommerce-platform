	package domain

	import "time"

	type Order struct {
		ID           int64       `json:"id" db:"id"`
		CustomerName string      `json:"customerName" db:"customer_name"`
		OrderItems   []OrderItem `json:"orderItems" db:"-"` // Keep this as `-` because it's populated separately
		Status       string      `json:"status" db:"status"`
		Created_at   time.Time   `json:"createdAt" db:"created_at"`
		IsDeleted    bool        `json:"isDeleted" db:"is_deleted"`
	}

	type OrderItem struct {
		ID        int64 `json:"id" db:"id"` // Add db tag for 'id' column from order_items table
		OrderID   int64 `json:"orderID,omitempty" db:"order_id"`
		ProductID int64 `json:"productID" db:"product_id"`
		Quantity  int64 `json:"quantity" db:"quantity"`
	}

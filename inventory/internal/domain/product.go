package domain

type Product struct {
	ID       int     `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Category string  `db:"category" json:"category"`
	Price    float64 `db:"price" json:"price"`
	Stock    int     `db:"stock" json:"stock"`
}

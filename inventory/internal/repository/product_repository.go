package repository

import (
	"fmt"
	"inventory/internal/domain"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	query := `INSERT INTO products (name, category, price, stock) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, product.Name, product.Category, product.Price, product.Stock)
	return err
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.DB.Select(&products, "SELECT * FROM products")
	return products, err
}
func (r *ProductRepository) GetByID(id int) (*domain.Product, error) {
	var product domain.Product
	err := r.DB.Get(&product, "SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Patch(id int, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	query := "UPDATE products SET "
	args := []interface{}{}
	i := 1
	for k, v := range fields {
		query += fmt.Sprintf("%s = $%d,", k, i)
		args = append(args, v)
		i++
	}
	query = query[:len(query)-1] // remove trailing comma
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	result, err := r.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no product with id %d", id)
	}

	return nil
}

func (r *ProductRepository) Delete(id int) error {
	result, err := r.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no product with id %d", id)
	}

	return nil
}

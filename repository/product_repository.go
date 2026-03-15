package repository

import (
	"go-roadmap/models"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ProductRepository interface {
	FindAllProduct() []models.Product
	SaveProduct(product models.Product)
	UpdateProduct(id int, product models.Product) error
	DeleteProduct(id int) error
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) FindAllProduct() []models.Product {
	var products []models.Product
	 err := r.db.Select(&products,"SELECT id, name_product, item, type FROM products")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return products
}

func (r *productRepo) SaveProduct(product models.Product) {
	_, err := r.db.NamedExec(
		`INSERT INTO products(name_product, item, type) VALUES(:name_product,:item,type)`, 
	product,
)
	if err != nil {
		log.Println("gagal Menambahkan Product", err)
	}
}

func (r *productRepo) UpdateProduct(id int, product models.Product) error {
	_, err := r.db.Exec("UPDATE products SET name_product=$1, item=$2, type=$3 WHERE id=$4", product.NameProduct, product.Item, product.Type, id)
	return err
}

func (r *productRepo) DeleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}

package repo

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"imgUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `INSERT INTO products (
		title, 
		description, 
		price, 
		img_url
	)
	VALUES (
	 	$1, $2, $3, $4
	)
	RETURNING id`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL)
	err := row.Scan(&p.ID)
	if err != nil {
		log.Printf("error while creating product: %v", err)
		return nil, err
	}

	return &p, nil
}
func (r *productRepo) Get(id int) (*Product, error) {
	var p Product

	query := `SELECT id, title, description, price, img_url FROM products WHERE id = $1`
	err := r.db.Get(&p, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("product not found: %v", err)
			return nil, nil
		}
		log.Printf("error while getting product: %v", err)
		return nil, err
	}

	return &p, nil
}
func (r *productRepo) List() ([]*Product, error) {
	var productList []*Product

	query := `SELECT id, title, description, price, img_url FROM products`
	err := r.db.Select(&productList, query)
	if err != nil {
		log.Printf("error while listing products: %v", err)
		return nil, err
	}

	return productList, nil
}
func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("error while deleting product: %v", err)
		return err
	}

	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	query := `UPDATE products SET
		title = $1,
		description = $2,
		price = $3,
		img_url = $4
	WHERE id = $5`
	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImgURL, product.ID)
	err := row.Err()
	if err != nil {
		log.Printf("error while updating product: %v", err)
		return nil, err
	}

	return &product, nil
}

// func generateInitialProducts(r *productRepo) {
// 	product1 := Product{
// 		ID:          1,
// 		Title:       "Wireless Headphones",
// 		Description: "Noise-cancelling over-ear headphones with 20 hours of battery life.",
// 		Price:       99.99,
// 		ImgURL:      "https://example.com/images/headphones.jpg",
// 	}

// 	product2 := Product{
// 		ID:          2,
// 		Title:       "Mechanical Keyboard",
// 		Description: "RGB backlit mechanical keyboard with blue switches.",
// 		Price:       79.50,
// 		ImgURL:      "https://example.com/images/keyboard.jpg",
// 	}

// 	product3 := Product{
// 		ID:          3,
// 		Title:       "Smart Watch",
// 		Description: "Fitness tracking smartwatch with heart rate monitor and GPS.",
// 		Price:       129.00,
// 		ImgURL:      "https://example.com/images/smartwatch.jpg",
// 	}

// 	product4 := Product{
// 		ID:          4,
// 		Title:       "Gaming Mouse",
// 		Description: "Ergonomic gaming mouse with adjustable DPI and 6 programmable buttons.",
// 		Price:       45.75,
// 		ImgURL:      "https://example.com/images/mouse.jpg",
// 	}

// 	product5 := Product{
// 		ID:          5,
// 		Title:       "4K Monitor",
// 		Description: "27-inch 4K UHD monitor with HDR and slim bezels.",
// 		Price:       299.99,
// 		ImgURL:      "https://example.com/images/monitor.jpg",
// 	}

// }

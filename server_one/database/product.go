package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Wireless Headphones",
		Description: "Noise-cancelling over-ear headphones with 20 hours of battery life.",
		Price:       99.99,
		ImgURL:      "https://example.com/images/headphones.jpg",
	}

	product2 := Product{
		ID:          2,
		Title:       "Mechanical Keyboard",
		Description: "RGB backlit mechanical keyboard with blue switches.",
		Price:       79.50,
		ImgURL:      "https://example.com/images/keyboard.jpg",
	}

	product3 := Product{
		ID:          3,
		Title:       "Smart Watch",
		Description: "Fitness tracking smartwatch with heart rate monitor and GPS.",
		Price:       129.00,
		ImgURL:      "https://example.com/images/smartwatch.jpg",
	}

	product4 := Product{
		ID:          4,
		Title:       "Gaming Mouse",
		Description: "Ergonomic gaming mouse with adjustable DPI and 6 programmable buttons.",
		Price:       45.75,
		ImgURL:      "https://example.com/images/mouse.jpg",
	}

	product5 := Product{
		ID:          5,
		Title:       "4K Monitor",
		Description: "27-inch 4K UHD monitor with HDR and slim bezels.",
		Price:       299.99,
		ImgURL:      "https://example.com/images/monitor.jpg",
	}

	ProductList = append(ProductList, product1, product2, product3, product4, product5)
}

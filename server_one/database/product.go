package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func GetSingle(productId int) *Product {
	for _, product := range productList {
		if product.ID == productId {
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for i, p := range productList {
		if p.ID == product.ID {
			productList[i] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}

	productList = tempList
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

	productList = append(productList, product1, product2, product3, product4, product5)
}

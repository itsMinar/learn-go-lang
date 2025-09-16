package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productId {
			return product, nil
		}
	}

	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(productId int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList

	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for i, p := range r.productList {
		if p.ID == product.ID {
			r.productList[i] = &product
		}
	}

	return &product, nil
}

func generateInitialProducts(r *productRepo) {
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

	r.productList = append(r.productList, &product1, &product2, &product3, &product4, &product5)
}

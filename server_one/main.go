package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreFlightReq(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed, need GET request", http.StatusMethodNotAllowed)
		return
	}

	sendData(w, productList, http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreFlightReq(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed, need POST request", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, http.StatusCreated)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Minar")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreFlightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProducts(w, r)
		case http.MethodPost:
			createProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
	})

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server stopped")
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

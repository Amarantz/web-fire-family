package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductID           int     `json:"productid,omitempty"`
	ProductName         string  `json:"productname,omitempty"`
	InventoryScanningID int     `json:"inventoryscanningid,omitempty"`
	Color               string  `json:"color,omitempty"`
	Size                string  `json:"size,omitempty"`
	Price               float32 `json:"price,omitempty"`
	Dimensions          string  `json:"dimensions,omitempty"`
	SKU                 int     `json:"sku,omitempty"`
}

var Products []Product

// InitRoutes creates the web API routes and sets their event handler functions
func InitRoutes() http.Handler {
	router := mux.NewRouter()
	//This should bring a list of all the Products

	Products = append(Products, Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	Products = append(Products, Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	Products = append(Products, Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.HandleFunc("/product", getProducts).Methods("GET")
	// This should bring back a specific Product
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	//This creates a new product using a Json String
	router.HandleFunc("/product/create", createProduct).Methods("POST")
	//This updates a product using a Json String
	router.HandleFunc("/product/update/{id}", updateProduct).Methods("PUT")
	//This sets the product to inactive in the database
	router.HandleFunc("/product/delete/{id}", deleteProduct).Methods("DELETE")
	//This allows us to set the quantity value of a product.
	router.HandleFunc("/inventory/update/{id}/{quantity}", updateInventory).Methods("PUT")
	return router
}

// Returns all of the products stored in the database in JSON format
func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

// Returns a specific product from the database in JSON format
func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	index := -1
	for i, value := range Products {
		productID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
			json.NewEncoder(w).Encode("")
		}

		if productID < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product ID."))
			return
		}

		if value.ProductID == productID {
			index = i
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Products[index])
			return
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}
}

// Creates a Product object from the passed in JSON Product and stores it in the database
func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	Products = append(Products, product)
}

// Deletes the specified product from the database
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, value := range Products {
		productID, err := strconv.Atoi(id)
		if err != nil || productID < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product ID."))
			return
		}

		if value.ProductID == productID {
			Products[i] = Products[len(Products)-1]
			Products = Products[:len(Products)-1]
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Invalid product ID."))
}

// Updates the product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	params := mux.Vars(r)
	id := params["id"]

	for i, value := range Products {
		productID, err := strconv.Atoi(id)
		if err != nil || productID < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product ID."))
			return
		}

		if value.ProductID == productID {
			Products[i] = product
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Invalid product ID."))
}

// Updates the inventory value for the inventory item
func updateInventory(w http.ResponseWriter, r *http.Request) {

}
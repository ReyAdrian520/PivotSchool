package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var products []product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func returnProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	}

	for _, product := range products {
		if int64(product.ID) == id {
			if err = json.NewEncoder(w).Encode(product); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

func createNewProductHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item product
	if err := json.Unmarshal(reqBody, &item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	products = append(products, item)

	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	}

	var updatedProduct product
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(reqBody, &updatedProduct); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	for i, product := range products {
		if int64(product.ID) == id {
			product.ID = updatedProduct.ID
			product.Name = updatedProduct.Name
			product.Description = updatedProduct.Description
			product.Price = updatedProduct.Price
			products[i] = product
			if err = json.NewEncoder(w).Encode(product); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	}

	for index, product := range products {
		if int64(product.ID) == id {
			products = append(products[:index], products[index+1:]...)
		}
	}
}

func initProducts() {
	bs, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bs, &products); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initProducts()
	r := mux.NewRouter()

	r.HandleFunc("/product", createNewProductHandler).Methods("POST")
	r.HandleFunc("/product/{id}", updateProductHandler).Methods("PUT")
	r.HandleFunc("/product/{id}", deleteProductHandler).Methods("DELETE")
	r.HandleFunc("/product/{id}", returnProductHandler).Methods("GET")
	r.HandleFunc("/products", getProductsHandler).Methods("GET")

	log.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package handler

import (
	"encoding/json"
	"module/model"
	"net/http"
	"strconv"
	"strings"
)

func (p *Products) ProductCreate(w http.ResponseWriter, r *http.Request) {
	product := model.Products{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = p.Product.CreateProduct(product)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created successfully"))
}

func (p *Products) ProductRead(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/ReadUsers/"))
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    products, err := p.Product.GetProducts()
    if err != nil {
        http.Error(w, "Failed to get products", http.StatusInternalServerError)
        return
    }

    var product model.Products
    for _, p := range products {
        if p.Id == id {
            product = p
            break
        }
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(product)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}



func (p *Products) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	product := model.Products{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = p.Product.UpdateProduct(product)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product updated successfully"))
}

func (p *Products) ProductDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/DeleteUsers/"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = p.Product.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product deleted successfully"))
}

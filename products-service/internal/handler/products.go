package handler

import (
	"net/http"
	"strconv"

	"github.com/DHBW-Malte/devops-lecture-project/pkg/httpx"
	"github.com/DHBW-Malte/devops-lecture-project/products-service/internal/service"
)

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpx.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	httpx.JSON(w, http.StatusOK, service.ListProducts())
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpx.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Product ID has wrong format")
		return
	}

	product, ok := service.GetProductByID(id)
	if !ok {
		httpx.Error(w, http.StatusNotFound, "Product not found")
		return
	}

	httpx.JSON(w, http.StatusOK, product)
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/DHBW-Malte/devops-lecture-project/pkg/httpx"
	"github.com/DHBW-Malte/devops-lecture-project/products-service/internal/model"
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
		httpx.Error(w, http.StatusNotFound, "Sorry, product not found")
		return
	}

	httpx.JSON(w, http.StatusOK, product)
}

func FilterProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.Error(w, http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}

	category := r.FormValue("category")
	filter := r.FormValue("filter")
	fValue := r.FormValue("value")

	if category != "name" && category != "price" {
		httpx.Error(w, http.StatusBadRequest, "Invalid Category")
		return
	}

	if filter != ">" && filter != "<" && filter != "=" {
		httpx.Error(w, http.StatusBadRequest, "Invalid filter")
		return
	}

	filtered := service.ProductsType{}
	if category == "name" {
		filtered = service.Products.Filter(func(product model.Product) bool { return product.Name == fValue })
	}

	if category == "price" {
		value, err := strconv.ParseFloat(fValue, 64)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "Invalid price value")
			return
		}
		switch filter {
		case "<":
			filtered = service.Products.Filter(func(product model.Product) bool { return product.Price < value })
		case ">":
			filtered = service.Products.Filter(func(product model.Product) bool { return product.Price > value })
		case "=":
			filtered = service.Products.Filter(func(product model.Product) bool { return product.Price == value })
		}
	}

	httpx.JSON(w, http.StatusOK, filtered)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	name := r.FormValue("name")
	price := r.FormValue("price")

	if name == "" || price == "" {
		httpx.Error(w, http.StatusBadRequest, "Please fill out the form completely")
		return
	}

	priceF, err := strconv.ParseFloat(price, 64)

	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid Price")
		return
	}

	newProduct := model.Product{
		ID:    len(service.Products) + 1,
		Name:  name,
		Price: priceF,
	}

	service.AddProduct(newProduct)
	httpx.JSON(w, http.StatusOK, "Product was successfully added")
}

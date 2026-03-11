package service

import (
	"github.com/DHBW-Malte/devops-lecture-project/products-service/internal/model"
)

type ProductsType []model.Product

var products = ProductsType{
	{ID: 1, Name: "Office PC", Price: 450.00},
	{ID: 2, Name: "Gaming PC", Price: 900.00},
	{ID: 3, Name: "Workstation", Price: 1500.00},
	{ID: 4, Name: "Server", Price: 2250.00},
	{ID: 5, Name: "Supe Computer", Price: 15000.00},
}

func ListProducts() []model.Product {
	return products
}

func GetProductByID(id int) (*model.Product, bool) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], true
		}
	}
	return nil, false
}

func (pr ProductsType) Filter(fn func(model.Product) bool) ProductsType {
	result := ProductsType{}
	for _, product := range pr {
		if fn(product) {
			result = append(result, product)
		}
	}
	return result
}

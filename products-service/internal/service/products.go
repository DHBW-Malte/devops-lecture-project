package service

import (
	"github.com/DHBW-Malte/devops-lecture-project/products-service/internal/model"
)

type ProductsType []model.Product

var Products = ProductsType{
	{ID: 1, Name: "Office PC", Price: 450.00},
	{ID: 2, Name: "Gaming PC", Price: 900.00},
	{ID: 3, Name: "Workstation", Price: 1500.00},
	{ID: 4, Name: "Server", Price: 2250.00},
	{ID: 5, Name: "Supe Computer", Price: 15000.00},
}

func ListProducts() []model.Product {
	return Products
}

func GetProductByID(id int) (*model.Product, bool) {
	for i := range Products {
		if Products[i].ID == id {
			return &Products[i], true
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

func AddProduct(product model.Product) bool {
	Products = append(Products, product)
	return true
}

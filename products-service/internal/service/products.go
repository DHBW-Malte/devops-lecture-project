package service

import "github.com/DHBW-Malte/devops-lecture-project/products-service/internal/model"

var products = []model.Product{
	{ID: 1, Name: "Office PC", Price: 450.00},
	{ID: 2, Name: "Gaming PC", Price: 900.00},
	{ID: 3, Name: "Workstation", Price: 1500.00},
	{ID: 4, Name: "Server", Price: 2500.00},
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

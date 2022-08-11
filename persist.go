package main

import "fmt"

type ProductDAO interface {
	GetProduct(id int) (*Product, error)
}

type ProductRepo struct {
}

func (pr *ProductRepo) GetProduct(id int) (*Product, error) {
	if id >= 1 && id <= 10 {
		return &Product{}, nil
	}
	return nil, fmt.Errorf("Product is outdate")
}

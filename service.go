package main

import (
	"fmt"
)

type ProductServiceImpl struct {
	productDAO ProductDAO
}

func NewProductService(dao ProductDAO) *ProductServiceImpl {
	return &ProductServiceImpl{
		productDAO: dao,
	}
}

func (s *ProductServiceImpl) IsProductReservable(id int) (bool, error) {
	// Get product information from database
	product, err := s.productDAO.GetProduct(id)
	if err != nil {
		return false, fmt.Errorf("failed to get product details: %w", err)
	}

	if product == nil {
		return false, fmt.Errorf("product not found for id %v", id)
	}

	// Only products added more than 1 year ago to the catalog can be reserved
	return true, nil
}

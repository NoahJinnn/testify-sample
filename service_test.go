package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductDAOMock struct {
	mock.Mock
}

func (pdM *ProductDAOMock) GetProduct(id int) (*Product, error) {
	args := pdM.Called(id)
	return args.Get(0).(*Product), args.Error(1)
}

func TestNewProductServiceImpl(t *testing.T) {
	mockProduct := ProductDAOMock{}
	productSvc := NewProductService(&mockProduct)
	assert.NotNil(t, productSvc)
}

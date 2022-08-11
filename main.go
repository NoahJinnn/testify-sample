package main

type ProductService interface {
	IsProductReservable(id int) (bool, error)
}

// Test cases:
// 1. The service implementation needs to respect the service definition
// 2. Products added to the catalog more than 1 year ago are reservable
// 3. Other products are not reservable
// 4. Products not in the catalog should cause a product not found error

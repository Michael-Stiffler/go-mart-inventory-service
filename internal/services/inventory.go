package services

import (
	"errors"
)

type InventoryService struct{}

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Inventory struct {
	Products []Product `json:"products"`
}

func NewInventoryService() *InventoryService {
	return &InventoryService{}
}

func (s *InventoryService) GetInventoryByStoreId(id string) (Inventory, error) {
	if id != "123" {
		return Inventory{}, errors.New("Store not found with id: " + id)
	}

	inventory := Inventory{
		Products: []Product{
			{
				ID:       1,
				Name:     "Product A",
				Price:    10.99,
				Quantity: 100,
			},
			{
				ID:       2,
				Name:     "Product B",
				Price:    20.50,
				Quantity: 50,
			},
		},
	}
	return inventory, nil
}

package models

// swagger:model 
type SalePlacementRequest struct {
	ProductType	string	`json:"product_type"`
	Price		int64	`json:"price"`
}

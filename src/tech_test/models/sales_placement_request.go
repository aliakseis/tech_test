package models

// swagger:model 
type SalesPlacementRequest struct {
	ProductType	string	`json:"product_type"`
	Price		int64	`json:"price"`
	Quantity		int64	`json:"quantity"`
}

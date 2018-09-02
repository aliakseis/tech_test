package models

// swagger:model 
type SalesAdjustmentRequest struct {
	ProductType	string	`json:"product_type"`
	Multiplier		int64	`json:"multiplier"`
	Divider		int64	`json:"divider"`
	Addend		int64	`json:"addend"`
}

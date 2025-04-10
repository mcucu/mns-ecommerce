package models

type Package struct {
	Items       []Product `json:"items"`
	TotalWeight float64   `json:"total_weight"`
	TotalPrice  float64   `json:"total_price"`
	CourierCost float64   `json:"courier_cost"`
}

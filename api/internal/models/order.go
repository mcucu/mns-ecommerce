package models

type Order struct {
	Items       []Product `json:"items"`
	TotalPrice  float64   `json:"total_price"`
	TotalWeight float64   `json:"total_weight"`
}

func (o *Order) CalculateTotal() {
	totalPrice := 0.0
	totalWeight := 0.0
	for _, item := range o.Items {
		totalPrice += item.Price
		totalWeight += item.Weight
	}
	o.TotalPrice = totalPrice
	o.TotalWeight = totalWeight
}

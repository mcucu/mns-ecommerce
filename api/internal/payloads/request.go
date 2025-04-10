package payloads

type OrderRequest struct {
	Items []ProductRequest `json:"items"`
}

type ProductRequest struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Weight float64 `json:"weight"`
}

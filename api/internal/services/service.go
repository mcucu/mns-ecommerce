package services

import (
	"github.com/mcucu/mns-ecommerce/internal/repositories"
	"github.com/mcucu/mns-ecommerce/internal/utils"
)

// Options anything any service object needed
type Options struct {
	utils.Options
	*repositories.Repositories
}

// Services all service object injected here
type Services struct {
	OrderService   IOrderService
	ProductService IProductService
}

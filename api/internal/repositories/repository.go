package repositories

import "github.com/mcucu/mns-ecommerce/internal/utils"

type Options struct {
	utils.Options
}

// Repositories all repository object injected here
type Repositories struct {
	Product IProduct
}

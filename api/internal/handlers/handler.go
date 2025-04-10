package handlers

import (
	"github.com/mcucu/mns-ecommerce/internal/services"
	"github.com/mcucu/mns-ecommerce/internal/utils"
)

type Options struct {
	utils.Options
	*services.Services
}

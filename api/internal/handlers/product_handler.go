package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcucu/mns-ecommerce/internal/payloads"
)

type ProductHandler struct {
	Options
}

func NewProductHandler(opt Options) *ProductHandler {
	return &ProductHandler{
		Options: opt,
	}
}

// GetProducts godoc
// @Summary GetProducts.
// @Description GetProducts.
// @Tags Product
// @Accept json
// @Produce json
// @Router /products [get]
// @Success 200 {object} payloads.Response
// @Security ApiStaticToken
func (h *ProductHandler) GetProducts(c echo.Context) error {
	// Logic to get all products
	products, err := h.ProductService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payloads.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Failed to fetch products",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Data:    products,
	})
}

package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcucu/mns-ecommerce/internal/payloads"
)

type OrderHandler struct {
	Options
}

func NewOrderHandler(opt Options) *OrderHandler {
	return &OrderHandler{
		Options: opt,
	}
}

// CreateOrder godoc
// @Summary Place Order.
// @Description Place Order.
// @Tags Order
// @Accept json
// @Produce json
// @Router /orders [post]
// @Param payload body payloads.OrderRequest true "Order Request"
// @Success 200 {object} payloads.Response
// @Security ApiStaticToken
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	req := new(payloads.OrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Failed to create order",
			Data:    nil,
		})
	}

	// Call the service to create an order
	packages, err := h.OrderService.CreateOrder(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payloads.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Failed to create order",
			Data:    nil,
		})
	}

	// Return the response
	return c.JSON(http.StatusOK, payloads.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Data:    packages,
	})
}

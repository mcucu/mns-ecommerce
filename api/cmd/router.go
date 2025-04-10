package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcucu/mns-ecommerce/internal/handlers"
	"github.com/mcucu/mns-ecommerce/internal/payloads"
)

func Router(e *echo.Echo, opt handlers.Options) {
	var (
		orderHandler   = handlers.NewOrderHandler(opt)
		productHandler = handlers.NewProductHandler(opt)
	)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, payloads.Response{
			Code:       http.StatusOK,
			Success:    true,
			Message:    "API is running",
			Data:       nil,
			Exceptions: nil,
		})
	})

	// Add product routes here
	productRoutes := e.Group("products")
	productRoutes.GET("", productHandler.GetProducts)

	// Add order routes here
	orderRoutes := e.Group("orders")
	orderRoutes.GET("", orderHandler.GetOrder)
	orderRoutes.POST("", orderHandler.CreateOrder)
}

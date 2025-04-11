package services

import (
	"github.com/mcucu/mns-ecommerce/internal/models"
	"github.com/mcucu/mns-ecommerce/internal/payloads"
)

type IOrderService interface {
	CreateOrder(order payloads.OrderRequest) ([]models.Package, error)
}

type OrderService struct {
	Options
}

func NewOrderService(opt Options) IOrderService {
	return &OrderService{
		Options: opt,
	}
}

func (s *OrderService) CreateOrder(order payloads.OrderRequest) ([]models.Package, error) {
	packages := splitIntoPackages(order.Items)
	return packages, nil
}

func splitIntoPackages(items []payloads.ProductRequest) []models.Package {
	var packages []models.Package
	var currentPackage models.Package

	for _, item := range items {
		// If adding the item exceeds $250, finalize the current package
		if currentPackage.TotalPrice+item.Price >= 250 {
			currentPackage.CourierCost = calculateCourierCost(currentPackage.TotalWeight)
			packages = append(packages, currentPackage)
			currentPackage = models.Package{}
		}

		// Create a new product instance
		product := models.Product{
			Name:   item.Name,
			Price:  item.Price,
			Weight: item.Weight,
		}

		// Add the item to the current package
		currentPackage.Items = append(currentPackage.Items, product)
		currentPackage.TotalWeight += item.Weight
		currentPackage.TotalPrice += item.Price
	}

	// Add the last package if it has items
	if len(currentPackage.Items) > 0 {
		currentPackage.CourierCost = calculateCourierCost(currentPackage.TotalWeight)
		packages = append(packages, currentPackage)
	}

	return packages
}

func calculateCourierCost(weight float64) float64 {
	if weight > 0 && weight <= 200 {
		return 5
	} else if weight > 200 && weight <= 500 {
		return 10
	} else if weight > 500 && weight <= 1000 {
		return 15
	} else if weight > 1000 && weight <= 2000 {
		return 20
	}

	return 0
}

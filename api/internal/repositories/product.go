package repositories

import (
	"fmt"

	"github.com/mcucu/mns-ecommerce/internal/models"
)

type IProduct interface {
	GetAll() (res []models.Product, err error)
}

type Product struct {
	Options
}

func NewProduct(opt Options) IProduct {
	return &Product{
		Options: opt,
	}
}

func (r *Product) GetAll() (res []models.Product, err error) {
	rows, err := r.DB.Query("SELECT id, name, price, weight FROM products") // Query to select all products
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Product
		if err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Weight); err != nil {
			return res, err
		}
		res = append(res, item)
	}
	return
}

package utils

import (
    "errors"
)

const maxPrice = 1000.0
const maxWeight = 50.0

func ValidateOrder(totalPrice float64, totalWeight float64) error {
    if totalPrice > maxPrice {
        return errors.New("total price exceeds the maximum limit")
    }
    if totalWeight > maxWeight {
        return errors.New("total weight exceeds the maximum limit")
    }
    return nil
}

func ValidateItems(selectedItems []string) error {
    if len(selectedItems) == 0 {
        return errors.New("at least one item must be selected")
    }
    return nil
}
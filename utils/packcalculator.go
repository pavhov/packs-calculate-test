package utils

import "github.com/pavhov/packs-calculate-test/models"

func CalculatePacks(order *models.Order) map[int]int {
	packCounts := make(map[int]int)

	for _, packSize := range order.PackSizes {
		if order.Items >= packSize {
			packCounts[packSize] = order.Items / packSize
			order.Items %= packSize
		}
	}

	return packCounts
}

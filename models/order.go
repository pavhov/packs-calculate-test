package models

type Order struct {
	Items     int   `json:"Items" binding:"required"`
	PackSizes []int `json:"PackSizes" binding:"required"`
}

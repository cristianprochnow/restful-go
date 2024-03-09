package models

type Car struct {
	Id    int     `json:"id"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

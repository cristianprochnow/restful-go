package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	Id    int     `json:"id"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

var cars = []car{
	{
		Id:    1,
		Brand: "Honda",
		Model: "Civic",
		Price: 164000.5,
	},
	{
		Id:    2,
		Brand: "Toyota",
		Model: "Corola",
		Price: 145000.8,
	},
	{
		Id:    3,
		Brand: "Kia",
		Model: "Cerato",
		Price: 123500.5,
	},
	{
		Id:    4,
		Brand: "Hyundai",
		Model: "Elantra",
		Price: 118560.5,
	},
}

func main() {
	router := gin.Default()

	router.GET("/cars", getCars)

	router.Run(":8080")
}

func getCars(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, cars)
}

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

type response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
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
	router.POST("/cars", postCars)

	router.Run(":8080")
}

func getCars(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, cars)
}

func postCars(request *gin.Context) {
	var newCar car
	requestError := request.BindJSON(&newCar)

	if requestError != nil {
		request.IndentedJSON(http.StatusBadRequest, response{
			Ok:      false,
			Message: "Formato dos dados enviados é inválido",
		})

		return
	}

	insertedCar := insertCar(newCar)

	request.IndentedJSON(http.StatusOK, insertedCar)
}

func insertCar(data car) car {
	newId := getLastId() + 1

	data.Id = newId
	cars = append(cars, data)

	return data
}

func getLastId() int {
	lastCar := getLastCar()
	lastId := lastCar.Id

	if lastId == 0 {
		lastId = 1
	}

	return lastId
}

func getLastCar() car {
	return cars[len(cars)-1]
}

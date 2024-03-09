package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		error(request, "É obrigatório o envio em formato JSON.")

		return
	}

	if !isValidCar(newCar) {
		error(request, "Formato de JSON inválido.")

		return
	}

	insertedCar := insertCar(newCar)

	request.IndentedJSON(http.StatusOK, insertedCar)
}

func isValidCar(dataSent any) bool {
	return fmt.Sprintf("%T", dataSent) == "car"
}

func insertCar(data car) car {
	newId := getLastId() + 1

	data.Id = newId
	cars = append(cars, data)

	return data
}

func error(request *gin.Context, message string) {
	request.IndentedJSON(http.StatusBadRequest, response{
		Ok:      false,
		Message: message,
	})
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

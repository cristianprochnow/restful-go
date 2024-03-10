package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	router.PUT("/cars/:id", updateCars)

	router.Run(":8080")
}

func getCars(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, listCars())
}

func postCars(request *gin.Context) {
	var newCar car
	requestError := request.BindJSON(&newCar)

	if requestError != nil {
		error(request, "Formato inválido de JSON enviado.")

		return
	}

	if !isValidCar(newCar) {
		error(request, "JSON enviado com dados obrigatórios faltando.")

		return
	}

	insertedCar := insertCar(newCar)

	request.IndentedJSON(http.StatusOK, insertedCar)
}

func updateCars(request *gin.Context) {
	var updateCar car
	requestError := request.BindJSON(&updateCar)
	carId := toInt(request.Param("id"))

	if carId == 0 {
		error(request, "ID do carro em formato inválido ou não enviado.")

		return
	}

	if requestError != nil {
		error(request, "Formato inválido de JSON enviado.")

		return
	}

	if !isValidCar(updateCar) {
		error(request, "JSON enviado com dados obrigatórios faltando.")

		return
	}

	updatedCar := refreshCar(carId, updateCar)

	if updatedCar.Id == 0 {
		error(request,
			fmt.Sprint("Carro não encontrado com o ID ", carId))

		return
	}

	request.IndentedJSON(http.StatusOK, updatedCar)
}

func toInt(text string) int {
	value, error := strconv.Atoi(text)

	if error != nil {
		value = 0
	}

	return value
}

func listCars() []car {
	return cars
}

func isValidCar(dataSent car) bool {
	isValid := true

	if len(dataSent.Brand) == 0 ||
		len(dataSent.Model) == 0 ||
		dataSent.Price == 0 {
		isValid = false
	}

	return isValid
}

func refreshCar(carId int, dataSent car) car {
	carItem := searchCar(carId)

	carItem.Model = dataSent.Model
	carItem.Brand = dataSent.Brand
	carItem.Price = dataSent.Price

	for index, item := range cars {
		if item.Id == carId {
			cars[index] = carItem
		}
	}

	return carItem
}

func searchCar(carId int) car {
	var carItem car

	for _, item := range listCars() {
		if item.Id == carId {
			carItem = item

			break
		}
	}

	return carItem
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

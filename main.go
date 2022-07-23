package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type address struct {
	Cep        string `json:"cep"`
	Neiborhood string `json:"neiborhood"`
	Complement string `json:"complement"`
	Country    string `json:"country"`
	State      string `json:"state"`
	City       string `json:"city"`
	Ibge       string `json:"ibge"`
	Street     string `json:"street"`
}

func fetchFromViaCep(cep int) []byte {
	var cepToString string = strconv.Itoa(cep)
	response, e := http.Get("http://viacep.com.br/ws/" + cepToString + "/json/")
	datasFromAPI, e := ioutil.ReadAll(response.Body)

	if e != nil {
		fmt.Println(e.Error())
	}
	return datasFromAPI
}

func fetchAddress(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, fetchFromViaCep(13040109))
}

func runApp() {
	router := gin.Default()
	router.GET("/address", fetchAddress)
	router.Run("localhost:8080")
}

func main() {
	runApp()
}

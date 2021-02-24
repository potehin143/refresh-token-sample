package controller

import (
	"../models"
	"../utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Authenticate = func(writer http.ResponseWriter, r *http.Request) {

	fmt.Println("Login called")

	credentials := &models.LoginCredentials{}
	err := json.NewDecoder(r.Body).Decode(credentials) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		utils.Respond(writer, utils.Message(utils.WrongParameter))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := models.Login(credentials.Id, credentials.Password)
	if err != nil {
		log.Print(err)
		writer.WriteHeader(http.StatusUnprocessableEntity)
	}
	utils.Respond(writer, resp)
}

var Refresh = func(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("Refresh called")
	refreshData := &models.RefreshData{}
	err := json.NewDecoder(r.Body).Decode(refreshData) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		utils.Respond(writer, utils.Message(utils.WrongParameter))
		return
	}
	resp := models.Refresh(refreshData.Id, refreshData.RefreshToken)
	utils.Respond(writer, resp)
}

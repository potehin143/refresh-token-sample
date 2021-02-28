package controller

import (
	"../helper"
	"encoding/json"
	"fmt"
	guid "github.com/google/uuid"
	"net/http"
)

//97451926-d9ad-43a1-89f7-dfdd7435e2f9
var Login = func(writer http.ResponseWriter, r *http.Request) {

	fmt.Println("Login called")

	credentials := &helper.LoginCredentials{}
	err := json.NewDecoder(r.Body).Decode(credentials) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		helper.Respond(writer, helper.Message(helper.WrongParameter))

		return
	}
	_, err2 := guid.Parse(credentials.Id)
	if err2 != nil {
		writer.WriteHeader(http.StatusBadRequest)
		helper.Respond(writer, helper.Message(helper.WrongParameter))
		return
	}

	resp, _ := helper.Login(credentials.Id, credentials.Password, writer)

	helper.Respond(writer, resp)
}

var Refresh = func(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("Refresh called")
	refreshData := &helper.RefreshData{}
	err := json.NewDecoder(r.Body).Decode(refreshData) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		helper.Respond(writer, helper.Message(helper.WrongParameter))
		return
	}
	resp, _ := helper.Refresh(refreshData.Id, refreshData.RefreshToken, writer)
	helper.Respond(writer, resp)
}

var Logout = func(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Logout called")
	resp, _ := helper.Logout(writer, request)
	helper.Respond(writer, resp)
}

var Register = func(writer http.ResponseWriter, r *http.Request) {

	fmt.Println("Register called")

	credentials := &helper.LoginCredentials{}
	err := json.NewDecoder(r.Body).Decode(credentials) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		helper.Respond(writer, helper.Message(helper.WrongParameter))

		return
	}
	_, err2 := guid.Parse(credentials.Id)
	if err2 != nil {
		writer.WriteHeader(http.StatusBadRequest)
		helper.Respond(writer, helper.Message(helper.WrongParameter))
		return
	}

	resp, _ := helper.Register(credentials.Id, credentials.Password, writer)

	helper.Respond(writer, resp)
}

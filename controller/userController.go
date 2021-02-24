package controller

import (
	"../models"
	"../utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var GetUser = func(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(utils.WrongParameter))
		return
	}
	if id == 10 {
		w.WriteHeader(http.StatusNotFound)
		utils.Respond(w, utils.Message(utils.ErrorNotFound))
	}

	data := models.GetContact(uint(id))
	resp := utils.Message("success")
	resp["data"] = data
	utils.Respond(w, resp)
}

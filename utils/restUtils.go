package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	SUCCESS        = "success"
	BadCredentials = "error.bad.credentials"
	ErrorNotFound  = "error.not.found"
	WrongParameter = "error.wrong.parameter"
)

func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	data["timestamp"] = time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln("Error {}", err)
	}
}

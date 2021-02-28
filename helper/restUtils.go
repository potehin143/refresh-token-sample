package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	SUCCESS             = "success"
	BadCredentials      = "error.bad.credentials"
	ErrorNotFound       = "error.not.found"
	WrongParameter      = "error.wrong.parameter"
	Unauthorized        = "error.unauthorized"
	Forbidden           = "error.forbidden"
	AlreadyExists       = "error.already.exists"
	InternalServerError = "error.internal.server.error"
)

func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

func Respond(writer http.ResponseWriter, data map[string]interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	data["timestamp"] = time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		log.Fatalln("Error {}", err)
	}
}

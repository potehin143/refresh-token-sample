package helper

import (
	"context"
	"net/http"
	"strings"
)

const (
	AnonymousUserId = "00000000-0000-0000-0000-000000000000"
	AuthHeaderName  = "Authorization"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		notAuth := []string{ //urlPaths where authentication is not required
			"/api/user/login",
			"/api/user/refresh",
			"/api/user/register",
		}
		requestPath := request.URL.Path

		authRequired := true
		authenticated := false

		for _, value := range notAuth {
			if value == requestPath {
				authRequired = false
				break
			}
		}

		response := make(map[string]interface{})
		tokenHeader := request.Header.Get(AuthHeaderName) //Получение токена

		userId := AnonymousUserId

		tokenString := ""
		if tokenHeader != "" {
			split := strings.Split(tokenHeader, " ")

			if len(split) == 2 {
				tokenString = split[1] // it might be `Bearer {token-body}`
			} else {
				tokenString = tokenHeader
			}

			tokenUserId, err := retrieveAccessToken(tokenString)

			if err == nil {
				authenticated = true
				userId = tokenUserId
			}
		}

		if authRequired && !authenticated {
			response = Message(Unauthorized)
			writer.WriteHeader(http.StatusUnauthorized)
			Respond(writer, response)
		}

		ctx := context.WithValue(
			context.WithValue(request.Context(), "userId", userId),
			"authenticated", authenticated)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request) // calling next handler
	})
}

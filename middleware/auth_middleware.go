package middleware

import (
	"learn-go-restful-api/helper"
	"learn-go-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleWare(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (authMiddleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "RAHASIA" {
		// teruskan handler
		authMiddleware.Handler.ServeHTTP(writer, request)

	} else {
		// gak ada api key
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   nil,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}

}

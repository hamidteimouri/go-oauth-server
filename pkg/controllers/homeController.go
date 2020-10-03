package controllers

import (
	"net/http"

	"github.com/hamidteimouri/go-oauth-server/pkg/utils"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	message := utils.Message(true, "You are visiting our service...")
	utils.JsonResponse(writer, http.StatusOK, message)
}

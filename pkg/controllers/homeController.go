package controllers

import (
	"github.com/hamidteimouri/go-oauth-server/pkg/utils"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	utils.JsonResponse(writer, http.StatusOK, "this is method")
}

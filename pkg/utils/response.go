package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status":  status,
	}
}

func JsonResponse(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)

	err := json.NewEncoder(writer).Encode(data)

	if err != nil {
		fmt.Printf("error : %s", err)
	}

}

func ErrorResponse(writer http.ResponseWriter) {

}

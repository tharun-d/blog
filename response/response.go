package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w http.ResponseWriter, status int, message string, data interface{}) {

	responseData := ResponseData{
		Status:  status,
		Message: message,
		Data:    data,
	}
	response, _ := json.Marshal(responseData)

	w.Header().Set("HTTP", fmt.Sprintf("%d", status))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

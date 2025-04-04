package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/utils"
)

func Generate(w http.ResponseWriter, uuid string, statusCode int, description string, data any) {
	validUUID := utils.ValidateUUID(uuid)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := &models.GenericResponse{
		BaseJSONResponse: models.BaseJSONResponse{
			StatusCode: statusCode,
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       validUUID,
		},
		Description: description,
		Data:        data,
	}
	jsonGenericResponse, err := json.Marshal(response)
	if err != nil {
		console.Log.Warn(validUUID, "Error marshalling the error response: %v", err)
		return
	}

	_, err = w.Write(jsonGenericResponse)
	if err != nil {
		console.Log.Warn(validUUID, "Error writing the json response: %v", err)
		return
	}
	console.Log.Success(validUUID, "Response generated: %v", string(jsonGenericResponse))
	return
}

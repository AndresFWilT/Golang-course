package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/utils"
)

func GenerateError(w http.ResponseWriter, uuid string, statusCode int, error string) {
	validUUID := utils.ValidateUUID(uuid)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := &models.ErrorResponse{
		BaseJSONResponse: models.BaseJSONResponse{
			StatusCode: statusCode,
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       validUUID,
		},
		Error: error,
	}
	jsonErrorResponse, err := json.Marshal(errorResponse)
	if err != nil {
		console.Log.Warn(validUUID, "Error marshalling the error response: %v", err)
		return
	}

	_, err = w.Write(jsonErrorResponse)
	if err != nil {
		console.Log.Warn(validUUID, "Error writing the json response: %v", err)
		return
	}
	console.Log.Error(validUUID, "Response error generated: %v", string(jsonErrorResponse))
	return
}

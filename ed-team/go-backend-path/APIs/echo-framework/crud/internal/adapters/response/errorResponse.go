package response

import (
	"time"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/utils"
)

func GenerateError(c echo.Context, uuid string, statusCode int, error string) error {
	validUUID := utils.ValidateUUID(uuid)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(statusCode)
	errorResponse := &models.ErrorResponse{
		BaseJSONResponse: models.BaseJSONResponse{
			StatusCode: statusCode,
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       validUUID,
		},
		Error: error,
	}
	console.Log.Error(validUUID, "Response error generated: %v", errorResponse)
	return c.JSON(statusCode, errorResponse)
}

package response

import (
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/shared/validation"
	"time"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/model"
)

func GenerateError(c echo.Context, uuid string, statusCode int, error string) error {
	validUUID := validation.ValidateUUID(uuid)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(statusCode)
	errorResponse := &model.ErrorResponse{
		BaseJSONResponse: model.BaseJSONResponse{
			StatusCode: statusCode,
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       validUUID,
		},
		Error: error,
	}
	console.Log.Error(validUUID, "Response error generated: %v", errorResponse)
	return c.JSON(statusCode, errorResponse)
}

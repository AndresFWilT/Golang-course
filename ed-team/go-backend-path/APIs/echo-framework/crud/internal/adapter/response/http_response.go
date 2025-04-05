package response

import (
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/shared/validation"
	"time"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/model"
)

func Generate(c echo.Context, uuid string, statusCode int, description string, data any) error {
	validUUID := validation.ValidateUUID(uuid)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(statusCode)
	response := &model.GenericResponse{
		BaseJSONResponse: model.BaseJSONResponse{
			StatusCode: statusCode,
			DateTime:   time.Now().Format(time.RFC3339),
			UUID:       validUUID,
		},
		Description: description,
		Data:        data,
	}
	console.Log.Success(validUUID, "Response generated: %v", response)
	return c.JSON(statusCode, response)
}

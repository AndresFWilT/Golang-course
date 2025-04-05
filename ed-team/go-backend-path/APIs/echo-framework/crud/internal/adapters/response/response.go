package response

import (
	"time"

	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/labstack/echo"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/utils"
)

func Generate(c echo.Context, uuid string, statusCode int, description string, data any) error {
	validUUID := utils.ValidateUUID(uuid)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(statusCode)
	response := &models.GenericResponse{
		BaseJSONResponse: models.BaseJSONResponse{
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

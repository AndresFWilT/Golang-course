package validation

import (
	"github.com/AndresFWilT/afwt-clean-go-logger/console"
	"github.com/google/uuid"
)

func ValidateUUID(requestUUID string) string {
	if err := uuid.Validate(requestUUID); err != nil {
		newUUID := uuid.New().String()
		console.Log.Warn(newUUID, "invalid UUID, generating new one: %v", err)
		return newUUID
	}
	return requestUUID
}

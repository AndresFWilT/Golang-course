package utils

import (
	"fmt"
	"github.com/google/uuid"
)

func ValidateUUID(requestUUID string) string {
	if err := uuid.Validate(requestUUID); err != nil {
		_ = fmt.Errorf("invalid UUID: %v", err)
		return uuid.New().String()
	}
	return requestUUID
}

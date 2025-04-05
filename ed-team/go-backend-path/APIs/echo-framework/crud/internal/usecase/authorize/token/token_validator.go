package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/model"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/certificate"
)

func Validate(tokenString string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, verifyTokenFunction)
	if err != nil {
		return model.Claim{}, fmt.Errorf("cannot parse token %w", err)
	}
	if !token.Valid {
		return model.Claim{}, fmt.Errorf("invalid token")
	}
	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, fmt.Errorf("cannot get claim from token")
	}
	return *claim, nil
}

func verifyTokenFunction(t *jwt.Token) (interface{}, error) {
	return certificate.VerifyKey, nil
}

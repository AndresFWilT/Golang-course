package tokens

import (
	"fmt"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/certificates"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/models"
)

func GenerateToken(data *models.Login) (string, error) {
	claim := models.Claim{
		Email: data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "AFWT",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(certificates.SignKey)
	if err != nil {
		return "", fmt.Errorf("cannot get signed token %w", err)
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (models.Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, verifyTokenFunction)
	if err != nil {
		return models.Claim{}, fmt.Errorf("cannot parse token %w", err)
	}
	if !token.Valid {
		return models.Claim{}, fmt.Errorf("invalid token")
	}
	claim, ok := token.Claims.(*models.Claim)
	if !ok {
		return models.Claim{}, fmt.Errorf("cannot get claim from token")
	}
	return *claim, nil
}

func verifyTokenFunction(t *jwt.Token) (interface{}, error) {
	return certificates.VerifyKey, nil
}

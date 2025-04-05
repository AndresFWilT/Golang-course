package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/domain/model"
	"github.com/AndresFWilT/afwt-clean-go-crud-echo/internal/usecase/authorize/certificate"
)

func Generate(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "AFWT",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(certificate.SignKey)
	if err != nil {
		return "", fmt.Errorf("cannot get signed token %w", err)
	}

	return signedToken, nil
}

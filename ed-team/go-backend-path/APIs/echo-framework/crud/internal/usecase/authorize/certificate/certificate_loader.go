package certificate

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"sync"
)

var (
	SignKey   *rsa.PrivateKey
	VerifyKey *rsa.PublicKey
	once      sync.Once
)

func Load(privateCertPath, publicCertPath string) error {
	var err error
	once.Do(func() {
		err = loadCertificatesFiles(privateCertPath, publicCertPath)
	})

	return err
}

func loadCertificatesFiles(privateCertPath, publicCertPath string) error {
	privateBytes, err := os.ReadFile(privateCertPath)
	if err != nil {
		return fmt.Errorf("could not read private key: %w", err)
	}

	publicBytes, err := os.ReadFile(publicCertPath)
	if err != nil {
		return fmt.Errorf("could not read public key: %w", err)
	}

	return parseRSACerts(privateBytes, publicBytes)
}

func parseRSACerts(privateBytes, publicBytes []byte) error {
	var err error

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return fmt.Errorf("could not parse private key: %w", err)
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return fmt.Errorf("could not parse public key: %w", err)
	}

	return nil
}

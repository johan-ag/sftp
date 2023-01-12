package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/mercadolibre/fury_hsm-lib/v4/environment"
	"github.com/mercadolibre/fury_hsm-lib/v4/pkg/hsm"
	"github.com/mercadolibre/fury_hsm-lib/v4/pkg/hsm/model"
)

func main() {
	if err := run(); err != nil {
		log.Println("[ERROR][%s]", err.Error())
	}
}
func run() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		return err
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}
	privateKeyPem := string(pem.EncodeToMemory(&privateKeyBlock))

	publicKey := privateKey.PublicKey
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	publicKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}
	publicKeyPem := string(pem.EncodeToMemory(&publicKeyBlock))

	fmt.Println(privateKeyPem)
	fmt.Println(publicKeyPem)


	return nil

}

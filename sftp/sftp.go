package sftp

import (
	"fmt"

	"github.com/mercadolibre/fury_go-core/pkg/crypto/publickey"
	"github.com/mercadolibre/fury_go-core/pkg/crypto/secretkey"
)

func RandomKey() error {
	// generates a pair of keys
	public, private := publickey.NewRandomPair()

	// encode public-key as PEM
	pem, err := public.MarshalPEM()
	if err != nil {
		return err
	}
	fmt.Printf("public key:\n%s\n", pem)
	
	// generates a secret key
	privateMasterKey := secretkey.NewRandom()
	
	// encode & encrypt private-key as PEM
	pem, err = private.MarshalEncryptedPEM(privateMasterKey)
	if err != nil {
		return err
	}
	fmt.Printf("private key:\n%s\n", pem)
	return nil
}

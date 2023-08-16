package jwsutils

import (
	"crypto/rsa"
	"fmt"

	"github.com/go-jose/go-jose/v3"
)

func SignWithGoJOSE(payload string, privateKey *rsa.PrivateKey) (string, error) {
	fmt.Println(privateKey)
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS512, Key: privateKey}, nil)
	if err != nil {
		return "", err
	}

	object, err := signer.Sign([]byte(payload))
	if err != nil {
		return "", err
	}

	serialized, err := object.CompactSerialize()
	return serialized, err
}

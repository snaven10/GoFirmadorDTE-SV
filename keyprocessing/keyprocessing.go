package keyprocessing

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
)

type KeyGenerator struct{}

func NewKeyGenerator() *KeyGenerator {
	return &KeyGenerator{}
}

func (kg *KeyGenerator) ByteToPrivateKey(bytes []byte) (*rsa.PrivateKey, error) {
	priv, err := x509.ParsePKCS8PrivateKey(bytes)
	if err != nil {
		return nil, err
	}

	rsaPriv, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("la clave no es de tipo RSA")
	}

	return rsaPriv, nil
}

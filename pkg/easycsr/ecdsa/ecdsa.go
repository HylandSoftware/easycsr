package ecdsa

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.hylandqa.net/do/easycsr/cmd/args"
)

func LoadOrGeneratePrivateKey(commonArgs *args.Common, ecdsaArgs *args.ECDSA) (*ecdsa.PrivateKey, error) {
	keyPath := commonArgs.Key
	if f, err := os.Stat(keyPath); err == nil && !f.IsDir() {
		raw, err := ioutil.ReadFile(keyPath)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Loading private key from %s\n", keyPath)
		decoded, _ := pem.Decode(raw)
		return x509.ParseECPrivateKey(decoded.Bytes)
	} else if f != nil && f.IsDir() {
		return nil, errors.New("key is a directory, not a file")
	}

	err := ecdsaArgs.DecodeCurve()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Generating new EC Private Key on curve %s", ecdsaArgs.Curve.Params().Name)

	key, err := ecdsa.GenerateKey(ecdsaArgs.Curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	b, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}

	keyBuff := new(bytes.Buffer)
	if err := pem.Encode(keyBuff, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}); err != nil {
		return nil, err
	}

	if _, err := os.Stat(keyPath); err != nil && os.IsNotExist(err) {
		if err := ioutil.WriteFile(keyPath, keyBuff.Bytes(), 0644); err != nil {
			return nil, err
		}
	}

	return key, nil
}

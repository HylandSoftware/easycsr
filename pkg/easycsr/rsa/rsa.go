package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.hylandqa.net/do/easycsr/cmd/args"
)

func LoadOrGeneratePrivateKey(common *args.Common, rsaArgs *args.RSA) (*rsa.PrivateKey, error) {
	keyPath := common.Key
	if f, err := os.Stat(keyPath); err == nil && !f.IsDir() {
		raw, err := ioutil.ReadFile(keyPath)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Loading private key from %s\n", keyPath)
		decoded, _ := pem.Decode(raw)
		return x509.ParsePKCS1PrivateKey(decoded.Bytes)
	} else if f != nil && f.IsDir() {
		return nil, errors.New("key is a directory, not a file")
	}

	fmt.Printf("Generating new private key of length %d\n", rsaArgs.KeyLength)

	key, err := rsa.GenerateKey(rand.Reader, rsaArgs.KeyLength)
	if err != nil {
		return nil, err
	}

	keyBuff := new(bytes.Buffer)
	if err := pem.Encode(keyBuff, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}); err != nil {
		return nil, err
	}

	if _, err := os.Stat(keyPath); err != nil && os.IsNotExist(err) {
		if err := ioutil.WriteFile(keyPath, keyBuff.Bytes(), 0644); err != nil {
			return nil, err
		}
	}

	return key, nil
}

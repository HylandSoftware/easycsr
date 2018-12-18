package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

func LoadOrGeneratePrivateKey(keyPath string) (*rsa.PrivateKey, error) {
	if f, err := os.Stat(keyPath); err == nil && !f.IsDir() {
		raw, err := ioutil.ReadFile(keyPath)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Loading private key from %s\n", keyPath)
		decoded, _ := pem.Decode(raw)
		return x509.ParsePKCS1PrivateKey(decoded.Bytes)
	}

	length := viper.GetInt("key-length")
	fmt.Printf("Generating new private key of length %d\n", length)

	key, err := rsa.GenerateKey(rand.Reader, length)
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

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
	"strings"

	"github.com/spf13/viper"
)

// DecodeSignatureAlgorithm decodes a string to an x509.SignatureAlgorithm
func DecodeSignatureAlgorithm() (x509.SignatureAlgorithm, error) {
	alg := viper.GetString("signature-algorithm")
	switch strings.ToLower(alg) {
	case "sha256":
		return x509.SHA256WithRSA, nil
	case "sha384":
		return x509.SHA384WithRSA, nil
	case "sha512":
		return x509.SHA512WithRSA, nil
	default:
		return x509.UnknownSignatureAlgorithm, fmt.Errorf("Unknown signature algorithm %s", alg)
	}
}

func LoadOrGeneratePrivateKey() (*rsa.PrivateKey, error) {
	keyPath := viper.GetString("key")
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

package ecdsa

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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
		return x509.ECDSAWithSHA256, nil
	case "sha384":
		return x509.ECDSAWithSHA384, nil
	case "sha512":
		return x509.ECDSAWithSHA512, nil
	default:
		return x509.UnknownSignatureAlgorithm, fmt.Errorf("Unknown signature algorithm %s", alg)
	}
}

func decodeCurve() (elliptic.Curve, error) {
	curve := viper.GetString("curve")
	switch strings.ToLower(curve) {
	case "p224":
		return elliptic.P224(), nil
	case "p256":
		return elliptic.P256(), nil
	case "p384":
		return elliptic.P384(), nil
	case "p521":
		return elliptic.P521(), nil
	default:
		return nil, errors.New("unknown curve")
	}
}

func LoadOrGeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	keyPath := viper.GetString("key")
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

	curve, err := decodeCurve()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Generating new EC Private Key on curve %s", curve.Params().Name)

	key, err := ecdsa.GenerateKey(curve, rand.Reader)
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

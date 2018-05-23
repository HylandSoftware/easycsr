package easycsr

import (
	"crypto/x509"
	"fmt"
	"strings"
)

// DecodeSignatureAlgorithm decodes a string to an x509.SignatureAlgorithm
func DecodeSignatureAlgorithm(alg string) (x509.SignatureAlgorithm, error) {
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

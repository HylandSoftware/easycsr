package args

import (
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type RSA struct {
	KeyLength int `mapstructure:"key-length"`
}

func (t *Common) DecodeRSASigAlg() error {
	switch strings.ToLower(t.SignatureAlgorithmRaw) {
	case "sha256":
		t.SignatureAlgorithm = x509.SHA256WithRSA
	case "sha384":
		t.SignatureAlgorithm = x509.SHA384WithRSA
	case "sha512":
		t.SignatureAlgorithm = x509.SHA512WithRSA
	default:
		return fmt.Errorf("unknown or unsupported signature algorithm %s", t.SignatureAlgorithmRaw)
	}

	return nil
}

func BindRSAArgs(flags *pflag.FlagSet) {
	flags.Int("key-length", 2048, "The length of the private key, in bits. Must be a power of two")

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

func RSAArgs() *RSA {
	result := &RSA{}

	if err := viper.Unmarshal(result); err != nil {
		panic(err)
	}

	return result
}

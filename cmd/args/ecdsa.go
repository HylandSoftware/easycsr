package args

import (
	"crypto/elliptic"
	"crypto/x509"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ECDSA struct {
	CurveRaw string         `mapstructure:"curve"`
	Curve    elliptic.Curve `mapstructure:"-"`
}

func (t *ECDSA) DecodeCurve() error {
	switch strings.ToLower(t.CurveRaw) {
	case "p224":
		t.Curve = elliptic.P224()
	case "p256":
		t.Curve = elliptic.P256()
	case "p384":
		t.Curve = elliptic.P384()
	case "p521":
		t.Curve = elliptic.P521()
	default:
		return errors.New("unknown curve")
	}

	return nil
}

func (t *Common) DecodeECDSASigAlg() error {
	switch strings.ToLower(t.SignatureAlgorithmRaw) {
	case "sha256":
		t.SignatureAlgorithm = x509.ECDSAWithSHA256
	case "sha384":
		t.SignatureAlgorithm = x509.ECDSAWithSHA384
	case "sha512":
		t.SignatureAlgorithm = x509.ECDSAWithSHA512
	default:
		return fmt.Errorf("unknown or unsupported signature algorithm %s", t.SignatureAlgorithmRaw)
	}

	return nil
}

func BindECDSAArgs(flags *pflag.FlagSet) {
	flags.String("curve", "P256", "The curve to use [P224, P256, P384, P521]")

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

func ECDSAArgs() *ECDSA {
	result := &ECDSA{}

	if err := viper.Unmarshal(result); err != nil {
		panic(err)
	}

	return result
}

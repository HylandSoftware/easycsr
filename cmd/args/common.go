package args

import (
	"crypto/x509"
	"crypto/x509/pkix"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Common houses args shared amongst all sub-commands
type Common struct {
	Key string `mapstructure:"key"`

	Country            string `mapstructure:"country"`
	State              string `mapstructure:"state"`
	Locality           string `mapstructure:"locality"`
	Organization       string `mapstructure:"org"`
	OrganizationalUnit string `mapstructure:"ou"`
	CommonName         string `mapstructure:"common-name"`

	Out                   string                  `mapstructure:"out"`
	SignatureAlgorithmRaw string                  `mapstructure:"signature-algorithm"`
	SignatureAlgorithm    x509.SignatureAlgorithm `mapstructure:"-"`

	SubjectAlternateNames []string `mapstructure:"san"`
}

func (t *Common) BuildSubject() *pkix.Name {
	return &pkix.Name{
		CommonName:         t.CommonName,
		Country:            []string{t.Country},
		Province:           []string{t.State},
		Locality:           []string{t.Locality},
		Organization:       []string{t.Organization},
		OrganizationalUnit: []string{t.OrganizationalUnit},
	}
}

func BindCommonArgs(flags *pflag.FlagSet) {
	flags.StringP("key", "k", "", "The key file to use. If it does not exist, it will be created with the specified key size")

	flags.String("country", "US", "Subject Country")
	flags.String("state", "Ohio", "Subject State")
	flags.String("locality", "Westlake", "Subject Locality")
	flags.String("org", "Hyland Software", "Subject Organization")
	flags.String("ou", "Research & Development", "Subject Organizational unit")
	flags.StringP("common-name", "n", "", "The common name (FQDN) for the certificate. Will be appended to the SAN list to conform with RFC2818 3.1")

	flags.String("out", "", "Where to save the CSR to. Printed to standard out if not specified")
	flags.String("signature-algorithm", "sha256", "The algorithm to sign the CSR with")

	flags.StringSlice("san", []string{}, "Subject Alternate Names (The subject will automatically be appended to this list)")

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

func CommonArgs() *Common {
	result := &Common{}

	if err := viper.Unmarshal(result); err != nil {
		panic(err)
	}

	return result
}

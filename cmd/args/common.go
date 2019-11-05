package args

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"os"
	"runtime"

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

var (
	shouldPersist *bool
	fileSettings  = viper.New()
)

func (t *Common) BuildSubject() *pkix.Name {
	name := &pkix.Name{CommonName: t.CommonName}

	if t.Country != "" {
		name.Country = []string{t.Country}
	}

	if t.State != "" {
		name.Province = []string{t.State}
	}

	if t.Locality != "" {
		name.Locality = []string{t.Locality}
	}

	if t.Organization != "" {
		name.Organization = []string{t.Organization}
	}

	if t.OrganizationalUnit != "" {
		name.OrganizationalUnit = []string{t.OrganizationalUnit}
	}

	return name
}

func BindCommonArgs(flags *pflag.FlagSet) {
	fileSettings.SetConfigName(".easycsr")
	fileSettings.AddConfigPath(".")
	fileSettings.AddConfigPath("$HOME")

	fileSettings.SetDefault("country", "US")
	fileSettings.SetDefault("signature-algorithm", "sha256")

	if err := fileSettings.ReadInConfig(); err != nil {
		if _, isNotFoundError := err.(viper.ConfigFileNotFoundError); !isNotFoundError {
			panic(fmt.Sprintf("failed to load config: %s", err))
		}
	}

	flags.String("country", fileSettings.GetString("country"), "Subject Country")
	flags.String("state", fileSettings.GetString("state"), "Subject State")
	flags.String("locality", fileSettings.GetString("locality"), "Subject Locality")
	flags.String("org", fileSettings.GetString("org"), "Subject Organization")
	flags.String("ou", fileSettings.GetString("ou"), "Subject Organizational unit")
	flags.String("signature-algorithm", fileSettings.GetString("signature-algorithm"), "The algorithm to sign the CSR with")

	if err := fileSettings.BindPFlags(flags); err != nil {
		panic(err)
	}

	flags.StringP("key", "k", "", "The key file to use. If it does not exist, it will be created with the specified key size")
	flags.StringP("common-name", "n", "", "The common name (FQDN) for the certificate. Will be appended to the SAN list to conform with RFC2818 3.1")

	flags.String("out", "", "Where to save the CSR to. Printed to standard out if not specified")

	flags.StringSlice("san", []string{}, "Subject Alternate Names (The subject will automatically be appended to this list)")

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}

	shouldPersist = flags.Bool("save", false, "Save common settings as defaults")
}

func CommonArgs() *Common {
	result := &Common{}

	if err := viper.Unmarshal(result); err != nil {
		panic(err)
	}

	return result
}

// Implementation copied from viper
// https://github.com/spf13/viper/blob/6d33b5a963d922d182c91e8a1c88d81fd150cfd4/util.go#L138-L147
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func TryPersistCommonArgs() {
	if !*shouldPersist {
		return
	}

	fmt.Println("Saving config")
	if fileSettings.ConfigFileUsed() == "" {
		if f, err := os.Create(fmt.Sprintf("%s/.easycsr.yaml", userHomeDir())); err != nil {
			panic("failed to create config file")
		} else {
			_ = f.Close()
		}
	}

	if err := fileSettings.WriteConfig(); err != nil {
		panic(fmt.Sprintf("failed to write config: %s", err))
	}
}

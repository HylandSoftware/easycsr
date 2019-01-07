package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version string
	rootCmd = &cobra.Command{
		Use:     "easycsr",
		Short:   "Certificate Requests, made easy",
		Long:    "easycsr simplifies the generation of Certificate Signing Requests by providing sane defaults and aiding in SAN generation",
		Args:    cobra.NoArgs,
		Version: version,
	}
)

func Init() {
	viper.SetEnvPrefix("easycsr")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	flags := rootCmd.PersistentFlags()
	flags.StringP("key", "k", "", "The key file to use. If it does not exist, it will be created with the specified key size")

	flags.String("country", "US", "Subject Country")
	flags.String("state", "Ohio", "Subject State")
	flags.String("locality", "Westlake", "Subject Locality")
	flags.String("org", "Hyland Software", "Subject Organization")
	flags.String("ou", "Research & Development", "Subject Organizational unit")
	flags.String("common-name", "", "The common name (FQDN) for the certificate. Will be appended to the SAN list to conform with RFC2818 3.1")

	flags.String("out", "", "Where to save the CSR to. Printed to standard out if not specified")
	flags.String("signature-algorithm", "sha256", "The algorithm to sign the CSR with")

	flags.StringSlice("san", []string{}, "Subject Alternate Names (The subject will automatically be appended to this list)")

	viper.BindPFlags(flags)

	rootCmd.AddCommand(RSA)
	initRSA()

	rootCmd.AddCommand(ECDSA)
	initECDSA()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

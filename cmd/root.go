package cmd

import (
	"os"
	"strings"

	"bitbucket.hylandqa.net/do/easycsr/cmd/args"
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
		PersistentPostRun: func(cmd *cobra.Command, _ []string) {
			args.TryPersistCommonArgs()
		},
	}
)

func init() {
	viper.SetEnvPrefix("easycsr")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	args.BindCommonArgs(rootCmd.PersistentFlags())

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

package cmd

import (
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr"
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr/rsa"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RSA = &cobra.Command{
	Use:   "rsa",
	Short: "Generate an RSA Certificate Signing Request",
	Run: func(_ *cobra.Command, _ []string) {
		key, err := rsa.LoadOrGeneratePrivateKey()
		if err != nil {
			panic(err)
		}

		sigAlg, err := rsa.DecodeSignatureAlgorithm()
		if err != nil {
			panic(err)
		}

		if err := easycsr.Build(key, sigAlg); err != nil {
			panic(err)
		}
	},
}

func initRSA() {
	flags := RSA.PersistentFlags()
	flags.Int("key-length", 2048, "The length of the private key, in bits. Must be a power of two")

	viper.BindPFlags(flags)
}

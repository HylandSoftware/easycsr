package cmd

import (
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr"
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr/ecdsa"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ECDSA = &cobra.Command{
	Use:   "ecdsa",
	Short: "Generate an ECDSA Certificate Signing Request",
	Run: func(_ *cobra.Command, _ []string) {
		key, err := ecdsa.LoadOrGeneratePrivateKey()
		if err != nil {
			panic(err)
		}

		sigAlg, err := ecdsa.DecodeSignatureAlgorithm()
		if err != nil {
			panic(err)
		}

		if err := easycsr.Build(key, sigAlg); err != nil {
			panic(err)
		}
	},
}

func initECDSA() {
	flags := ECDSA.PersistentFlags()
	flags.String("curve", "P256", "The curve to use [P224, P256, P384, P521]")

	viper.BindPFlags(flags)
}

package cmd

import (
	"bitbucket.hylandqa.net/do/easycsr/cmd/args"
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr"
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr/ecdsa"
	"github.com/spf13/cobra"
)

var ECDSA = &cobra.Command{
	Use:   "ecdsa",
	Short: "Generate an ECDSA Certificate Signing Request",
	Run: func(_ *cobra.Command, _ []string) {
		common, ecdsaArgs := args.CommonArgs(), args.ECDSAArgs()

		key, err := ecdsa.LoadOrGeneratePrivateKey(common, ecdsaArgs)
		if err != nil {
			panic(err)
		}

		if err := ecdsaArgs.DecodeCurve(); err != nil {
			panic(err)
		}

		if err := easycsr.Build(key, common); err != nil {
			panic(err)
		}
	},
}

func initECDSA() {
	args.BindECDSAArgs(ECDSA.PersistentFlags())
}

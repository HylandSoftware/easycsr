package cmd

import (
	"bitbucket.hylandqa.net/do/easycsr/cmd/args"
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr"
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr/rsa"
	"github.com/spf13/cobra"
)

var RSA = &cobra.Command{
	Use:   "rsa",
	Short: "Generate an RSA Certificate Signing Request",
	Run: func(_ *cobra.Command, _ []string) {
		common, rsaArgs := args.CommonArgs(), args.RSAArgs()

		key, err := rsa.LoadOrGeneratePrivateKey(common, rsaArgs)
		if err != nil {
			panic(err)
		}

		if err := common.DecodeRSASigAlg(); err != nil {
			panic(err)
		}

		if err := easycsr.Build(key, common); err != nil {
			panic(err)
		}
	},
}

func initRSA() {
	args.BindRSAArgs(RSA.PersistentFlags())
}

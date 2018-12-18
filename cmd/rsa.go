package cmd

import (
	"os"

	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr"
	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr/rsa"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RSA = &cobra.Command{
	Use:   "rsa",
	Short: "Generate an RSA Certificate Signing Request",
	Run: func(_ *cobra.Command, _ []string) {
		keyPath := viper.GetString("key")
		if f, _ := os.Stat(keyPath); f != nil && f.IsDir() {
			panic("Key is a directory, not a file!")
		}

		key, err := rsa.LoadOrGeneratePrivateKey(keyPath)
		if err != nil {
			panic(err)
		}

		if err := easycsr.Build(key); err != nil {
			panic(err)
		}
	},
}

func initRSA() {
	flags := RSA.PersistentFlags()
	flags.Int("key-length", 2048, "The length of the private key, in bits. Must be a power of two")

	viper.BindPFlags(flags)
}

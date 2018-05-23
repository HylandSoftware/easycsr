package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"bitbucket.hylandqa.net/do/easycsr/pkg/easycsr"
)

func main() {
	args := easycsr.DefaultArgs()

	flag.IntVar(&args.KeyLength, "keyLength", 2048, "RSA Key Length in Bits, must be a power of 2")
	flag.StringVar(&args.KeyPath, "key", "", "The key file to use. If it does not exist, it will be created with the specified key size")

	flag.StringVar(&args.Subject.Country, "country", "US", "Subject Country")
	flag.StringVar(&args.Subject.State, "st", "Ohio", "Subject State")
	flag.StringVar(&args.Subject.Locality, "locality", "Westlake", "Subject Locality")
	flag.StringVar(&args.Subject.Organization, "org", "Hyland Software", "Subject Organization")
	flag.StringVar(&args.Subject.OrganizationalUnit, "ou", "Research & Development", "Subject Organizational unit")
	flag.StringVar(&args.Subject.CommonName, "cn", "", "The common name (FQDN) for the certificate. Will be appended to the SAN list to conform with RFC2818 3.1")

	flag.StringVar(&args.OutPath, "out", "", "Where to save the CSR to. Printed to standard out if not specified")
	flag.StringVar(&args.SignatureAlgorithmRaw, "signatureAlgorithm", "sha256", "The algorithm to sign the CSR with")
	flag.Var(&args.AlternateNames, "san", "Subject Alternate Names")

	flag.Parse()

	if alg, err := easycsr.DecodeSignatureAlgorithm(args.SignatureAlgorithmRaw); err != nil {
		panic(err)
	} else {
		args.SignatureAlgorithm = alg
	}

	if !args.Validate() {
		fmt.Printf("Invalid syntax!\n\nUsage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if privateKey, csr, err := easycsr.Build(args); err != nil {
		panic(err)
	} else {
		if _, err := os.Stat(args.KeyPath); err != nil && os.IsNotExist(err) {
			if err := ioutil.WriteFile(args.KeyPath, []byte(privateKey), 0644); err != nil {
				panic(err)
			}
		}

		if strings.TrimSpace(args.OutPath) != "" {
			if err := ioutil.WriteFile(args.OutPath, []byte(csr), 0644); err != nil {
				panic(err)
			}
		}

		fmt.Printf("Generated CSR:\n%s\n", csr)
	}
}

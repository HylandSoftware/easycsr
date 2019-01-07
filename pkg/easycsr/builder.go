package easycsr

import (
	"bytes"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

func getSubject() pkix.Name {
	return pkix.Name{
		CommonName:         viper.GetString("common-name"),
		Country:            []string{viper.GetString("country")},
		Province:           []string{viper.GetString("state")},
		Locality:           []string{viper.GetString("locality")},
		Organization:       []string{viper.GetString("org")},
		OrganizationalUnit: []string{viper.GetString("ou")},
	}
}

func buildCSR(privateKey interface{}, sigAlg x509.SignatureAlgorithm) (string, error) {
	subject := getSubject()

	commonNameFound := false
	sanList := viper.GetStringSlice("subject-alternate-names")
	for _, san := range sanList {
		if strings.ToLower(san) == strings.ToLower(subject.CommonName) {
			commonNameFound = true
			break
		}
	}

	if !commonNameFound {
		sanList = append(sanList, subject.CommonName)
	}

	rawSubject, err := asn1.Marshal(subject.ToRDNSequence())
	if err != nil {
		return "", err
	}

	csrTemplate := x509.CertificateRequest{
		RawSubject:         rawSubject,
		SignatureAlgorithm: sigAlg,
		DNSNames:           sanList,
	}

	csr, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		return "", err
	}

	csrBuff := &bytes.Buffer{}
	if err := pem.Encode(csrBuff, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csr}); err != nil {
		return "", err
	}

	return csrBuff.String(), nil
}

func Build(privateKey interface{}, sigAlg x509.SignatureAlgorithm) error {
	if csr, err := buildCSR(privateKey, sigAlg); err != nil {
		return err
	} else {
		outPath := viper.GetString("out")
		if strings.TrimSpace(outPath) != "" {
			if err := ioutil.WriteFile(outPath, []byte(csr), 0644); err != nil {
				return err
			}
		}

		fmt.Printf("Generated CSR:\n%s\n", csr)
	}

	return nil
}

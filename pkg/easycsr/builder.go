package easycsr

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/hylandsoftware/easycsr/cmd/args"
)

func buildCSR(privateKey crypto.PrivateKey, common *args.Common) (string, error) {
	subject := common.BuildSubject()

	commonNameFound := false
	for _, san := range common.SubjectAlternateNames {
		if strings.ToLower(san) == strings.ToLower(subject.CommonName) {
			commonNameFound = true
			break
		}
	}

	if !commonNameFound {
		common.SubjectAlternateNames = append(common.SubjectAlternateNames, subject.CommonName)
	}

	rawSubject, err := asn1.Marshal(subject.ToRDNSequence())
	if err != nil {
		return "", err
	}

	csrTemplate := x509.CertificateRequest{
		RawSubject:         rawSubject,
		SignatureAlgorithm: common.SignatureAlgorithm,
		DNSNames:           common.SubjectAlternateNames,
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

func Build(privateKey interface{}, common *args.Common) error {
	if csr, err := buildCSR(privateKey, common); err != nil {
		return err
	} else {
		if strings.TrimSpace(common.Out) != "" {
			if err := ioutil.WriteFile(common.Out, []byte(csr), 0644); err != nil {
				return err
			}
		}

		fmt.Printf("Generated CSR:\n%s\n", csr)
	}

	return nil
}

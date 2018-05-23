package easycsr

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var dnsObjectIdentifier = asn1.ObjectIdentifier{2, 5, 29, 17, 2}

func (s Subject) toPkixName() pkix.Name {
	return pkix.Name{
		CommonName:         s.CommonName,
		Country:            []string{s.Country},
		Province:           []string{s.State},
		Locality:           []string{s.Locality},
		Organization:       []string{s.Organization},
		OrganizationalUnit: []string{s.OrganizationalUnit},
	}
}

func loadOrGeneratePrivateKey(args *Args) (*rsa.PrivateKey, error) {
	if f, err := os.Stat(args.KeyPath); err == nil && !f.IsDir() {
		raw, err := ioutil.ReadFile(args.KeyPath)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Loading private key from %s\n", args.KeyPath)
		decoded, _ := pem.Decode(raw)
		return x509.ParsePKCS1PrivateKey(decoded.Bytes)
	} else {
		fmt.Printf("Generating new private key of length %d\n", args.KeyLength)
		return rsa.GenerateKey(rand.Reader, args.KeyLength)
	}
}

// Build generates a private key and certificate signing request
// from the specified arguments. It returns the base64 encoded
// private key, the base64 encoded CSR, and any errors
func Build(args *Args) (string, string, error) {
	key, err := loadOrGeneratePrivateKey(args)
	if err != nil {
		return "", "", err
	}

	keyBuff := new(bytes.Buffer)
	if err := pem.Encode(keyBuff, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}); err != nil {
		return "", "", err
	}

	commonNameFound := false
	sanList := []string{}
	for _, san := range args.AlternateNames {
		if strings.ToLower(san) == strings.ToLower(args.Subject.CommonName) {
			commonNameFound = true
		}
		sanList = append(sanList, san)
	}

	if !commonNameFound {
		sanList = append(sanList, args.Subject.CommonName)
	}

	rawSubject, err := asn1.Marshal(args.Subject.toPkixName().ToRDNSequence())
	if err != nil {
		return "", "", err
	}

	csrTemplate := x509.CertificateRequest{
		RawSubject:         rawSubject,
		SignatureAlgorithm: args.SignatureAlgorithm,
		DNSNames:           sanList,
	}

	csr, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, key)
	if err != nil {
		return "", "", err
	}

	csrBuff := new(bytes.Buffer)
	if err := pem.Encode(csrBuff, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csr}); err != nil {
		return "", "", err
	}

	return keyBuff.String(), csrBuff.String(), nil
}

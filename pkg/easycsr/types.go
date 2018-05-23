package easycsr

import (
	"crypto/x509"
	"fmt"
	"strings"
)

// Subject represents the subject field of a certificate
type Subject struct {
	Country            string
	State              string
	Locality           string
	Organization       string
	OrganizationalUnit string
	CommonName         string
}

type SANList []string

// Args contains application arguments passed in via flags
type Args struct {
	KeyLength             int
	KeyPath               string
	Subject               Subject
	OutPath               string
	SignatureAlgorithmRaw string
	SignatureAlgorithm    x509.SignatureAlgorithm
	AlternateNames        SANList
}

func (s *SANList) String() string {
	return strings.Join(*s, " ")
}

func (s *SANList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// DefaultArgs creates a default Args object
func DefaultArgs() *Args {
	result := Args{
		KeyLength: 2048,
		Subject: struct {
			Country            string
			State              string
			Locality           string
			Organization       string
			OrganizationalUnit string
			CommonName         string
		}{
			Country:            "US",
			State:              "Ohio",
			Locality:           "Westlake",
			Organization:       "Hyland Software",
			OrganizationalUnit: "Research & Development",
			CommonName:         "",
		},
		OutPath:               "",
		SignatureAlgorithmRaw: "sha512",
		SignatureAlgorithm:    x509.UnknownSignatureAlgorithm,
		AlternateNames:        SANList{},
	}

	return &result
}

func (a *Args) Validate() bool {
	key := uint(a.KeyLength)
	if key&(key-1) != 0 {
		fmt.Println("Invalid Key Size: Not a power of 2")
		return false
	}

	if strings.TrimSpace(a.Subject.Country) == "" {
		fmt.Println("Missing required field: Country")
		return false
	} else if strings.TrimSpace(a.Subject.State) == "" {
		fmt.Println("Missing required field: State")
		return false
	} else if strings.TrimSpace(a.Subject.Locality) == "" {
		fmt.Println("Missing required field: Locality")
		return false
	} else if strings.TrimSpace(a.Subject.Organization) == "" {
		fmt.Println("Missing required field: Organization")
		return false
	} else if strings.TrimSpace(a.Subject.OrganizationalUnit) == "" {
		fmt.Println("Missing required field: OU")
		return false
	} else if strings.TrimSpace(a.Subject.CommonName) == "" {
		fmt.Println("Missing required field: Common Name")
		return false
	}

	if a.SignatureAlgorithm == x509.UnknownSignatureAlgorithm {
		fmt.Printf("Could not parse signature algorithm %s\n", a.SignatureAlgorithmRaw)
		return false
	}

	return true
}

package easycsr

// Subject represents the subject field of a certificate
type Subject struct {
	Country            string
	State              string
	Locality           string
	Organization       string
	OrganizationalUnit string
	CommonName         string
}

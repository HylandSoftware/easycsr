# EasyCSR

Starting with Chrome 58, you need to include the subject common name of a certificate in the SAN list as well.
This isn't exactly easy with OpenSSL, and this utility aims to make this as simple as possible.

## Building

The easiest way to build this is via `make`. If you do not have access to `make`, you can build the command directly:

```bash
go build ./cmd/easycsr -o easycsr -v
```

## Usage

```text
Usage of easycsr.exe:
  -cn string
        The common name (FQDN) for the certificate. Will be appended to the SAN list to conform with RFC2818 3.1
  -country string
        Subject Country (default "US")
  -key string
        The key file to use. If it does not exist, it will be created with the specified key size
  -keyLength int
        RSA Key Length in Bits, must be a power of 2 (default 2048)
  -locality string
        Subject Locality (default "Westlake")
  -org string
        Subject Organization (default "Hyland Software")
  -ou string
        Subject Organizational unit (default "Research & Development")
  -out string
        Where to save the CSR to. Printed to standard out if not specified
  -san value
        Subject Alternate Names
  -signatureAlgorithm string
        The algorithm to sign the CSR with (default "sha256")
  -st string
        Subject State (default "Ohio")
```

### Example

```bash
$ easycsr -key jenkins.hylandqa.net.key -cn jenkins.hylandqa.net -san '*.jenkins.hylandqa.net'
Generating new private key of length 2048
Generated CSR:
-----BEGIN CERTIFICATE REQUEST-----
MIIDGTCCAgECAQAwgYkxCzAJBgNVBAYTAlVTMQ0wCwYDVQQIEwRPaGlvMREwDwYD
VQQHEwhXZXN0bGFrZTEYMBYGA1UEChMPSHlsYW5kIFNvZnR3YXJlMR8wHQYDVQQL
DBZSZXNlYXJjaCAmIERldmVsb3BtZW50MR0wGwYDVQQDExRqZW5raW5zLmh5bGFu
ZHFhLm5ldDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL/m7Qc14C70
4VfCB5sD+TZIU2jCw4xRQb8Di0mO/ZJmDKcZfOBUFltpFGL+OrTIgHHmLg4lPIuv
gMFdKxoh6dVRW2FEhsOuGwns2xZtEOJVgCNn8aFqyJ4hkvw4Z6VbJDfXo74r4Qqs
uXVcCFhgtCA7b89aH3/5kFSbSaE7z/5NAZ6+tZS6B15WSFK1wdIbbW9gQgvStaRy
Zfb+j3qJgFzo+LWF9rdsJSkMhhvXKSxBnEhqMZJ7dvOZjH6fFfWEiBdt0pTXBtyY
lRh2poGaNcpaHuzGNveXe8mb8KeJsk2c8Lel7DvsYteHE4sowQvk6YJ1cNbUWR/w
sCd/VERdxP0CAwEAAaBKMEgGCSqGSIb3DQEJDjE7MDkwNwYDVR0RBDAwLoIWKi5q
ZW5raW5zLmh5bGFuZHFhLm5ldIIUamVua2lucy5oeWxhbmRxYS5uZXQwDQYJKoZI
hvcNAQELBQADggEBAARr5uGETstiAgiHqh3MNm+0Kxl/SpV8ptzx/6rO+oP+eJIW
loAVqJPIXhYasDBkkUO6EQxfT5ll+FcdrM0aCgzKGZ/7qsexlbtZv+WDGTC/S72/
jLFca5mwTV68mVPkqRyA26PFfQndvvFwMCtYi9ECmaZAo1B3YyMliZIoNmdbfEFG
e7r9NlFFKeJ7YoN4Zq7VLYnJkyXv0AGxA7QBpxt0cpkEXY2bfhZv/fCfFXgR1wCq
L2/UnsWvRJ6mITwsLu/XOqBUJK8W1x1FyBHOVGA8EYCiPKrqIBktmhiok1M1A/df
10BxP0I7+QsNzKAqgMx/Tn3ChQuIG142a3X39bs=
-----END CERTIFICATE REQUEST-----
```

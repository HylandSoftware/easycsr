# EasyCSR

![](https://github.com/hylandsoftware/easycsr/workflows/CI/badge.svg)  [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/hylandsoftware/easycsr)](https://goreportcard.com/report/github.com/hylandsoftware/easycsr)

Starting with Chrome 58, you need to include the subject common name of a certificate in the SAN list as well.
This isn't exactly easy with OpenSSL, and this utility aims to make this as simple as possible.

## Building

This project requires Go 1.11+ or vgo for Go Modules support. It can be built like any standard Go binary:

```bash
go build ./cmd/easycsr -o easycsr -v
```

## Usage

```text
easycsr simplifies the generation of Certificate Signing Requests by providing sane defaults and aiding in SAN generation

Usage:
  easycsr [command]

Available Commands:
  ecdsa       Generate an ECDSA Certificate Signing Request
  help        Help about any command
  rsa         Generate an RSA Certificate Signing Request

Flags:
  -n, --common-name string           The common name (FQDN) for the certificate. Will be appended to the SAN list to conform with RFC2818 3.1
      --country string               Subject Country
  -h, --help                         help for easycsr
  -k, --key string                   The key file to use. If it does not exist, it will be created with the specified key size
      --locality string              Subject Locality
      --org string                   Subject Organization
      --ou string                    Subject Organizational unit
      --out string                   Where to save the CSR to. Printed to standard out if not specified
      --san strings                  Subject Alternate Names (The subject will automatically be appended to this list)
      --save                         Save common settings as defaults
      --signature-algorithm string   The algorithm to sign the CSR with (default "sha256")
      --state string                 Subject State

Use "easycsr [command] --help" for more information about a command.
```

### Example

```bash
$ easycsr rsa -k jenkins.hylandqa.net.key --common-name jenkins.hylandqa.net --san '*.jenkins.hylandqa.net'
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

### Docker

If you do not have the `go` toolchain installed, you can use a pre-built docker
image instead. Just mount your working directory to `/csr`:

```bash
$ docker run -it --rm -v "$(pwd):/csr" hylandsoftware/easycsr rsa -k jenkins.hylandqa.net.key --common-name jenkins.hylandqa.net --san '*.jenkins.hylandqa.net'
Unable to find image 'hylandsoftware/easycsr:latest' locally
latest: Pulling from hylandsoftware/easycsr
Digest: sha256:1a176b600019e29a6856a66617179fe4be1f990052e47a5a16a0be2d32fe6dc0
Status: Downloaded newer image for hylandsoftware/easycsr:latest
Loading private key from jenkins.hylandqa.net.key
Generated CSR:
-----BEGIN CERTIFICATE REQUEST-----
MIIDGTCCAgECAQAwgYkxCzAJBgNVBAYTAlVTMQ0wCwYDVQQIEwRPaGlvMREwDwYD
VQQHEwhXZXN0bGFrZTEYMBYGA1UEChMPSHlsYW5kIFNvZnR3YXJlMR8wHQYDVQQL
DBZSZXNlYXJjaCAmIERldmVsb3BtZW50MR0wGwYDVQQDExRqZW5raW5zLmh5bGFu
ZHFhLm5ldDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALaKxCs1vGXS
wGWwh647x35amrrcX5poVEdFj3t3ncD3yYj8K8ICMuwMSSFbpBK/oQYBNf8cP3qk
ozH9kE/PP83IZVaExTgpAs+IAtihjl5NJ7SZX2/ZiLj/Oo4rS9H0oZi6S1LUwJk+
KJ7t58vHKOyuY1DDP0rdYVh+/VrB3SzwFdQF68I9mFsC/CLAMZL0Ueiawx30gl51
5qpkEWyQ41hGbtbhAdmeWm0Z1F5ys1ViMZfiRdLuS7vo25l+fxfXmfD3CcZNNrPv
jJRaPRwLZKis6mv7mkFQxupP9oVdDWY0iZzfFr35yfMhdKnA2eO1UD3kdyMM00vL
I2zVXgykRfsCAwEAAaBKMEgGCSqGSIb3DQEJDjE7MDkwNwYDVR0RBDAwLoIWKi5q
ZW5raW5zLmh5bGFuZHFhLm5ldIIUamVua2lucy5oeWxhbmRxYS5uZXQwDQYJKoZI
hvcNAQELBQADggEBAGjCx+/xtWB6CBjESMllH2sekwyfK8gJG1NeGS4jOs6V3VwW
uzo14ZT+Yc+KUiP7wa7e6El8vGshwa9SGnj7TpVG8yggwJDi/wqYuMSROJL3zp1d
ki8/DHhUAcGqdFdUHnbI92f5nqgg6CgsV7xVHPGr/2XEc9yfkRXfJZVw+3cXhp5a
8W/57v6V4CS8izBA6ElXX2xIm5AIfAWh6xWrJBRfAFCGAA5Acm13mtqLVNVQwFqi
3MkTa49FrFHkX9U7WD3dYcpHXH/XJnyQa8eVDbtxXBEE0ZQLER+ync0Ay4G/OVtH
6lt7t8sm+cuAPbCzaWaf6Z1hE3/AH/gXckJ8xyY=
-----END CERTIFICATE REQUEST-----
```

# IPGeo

IPGeo adalah CLI berbasis Go untuk melakukan lookup informasi IP/domain.

## Fitur

- IP geolocation
- Domain → IP resolution
- ISP dan Organization
- ASN
- Timezone
- Latitude / Longitude
- Reverse DNS
- Local cache
- Local history
- JSON output
- Request timeout
- Version information

## Build

```bash
gofmt -w *.go
go build -o ipgeo .
```
## Usage
```bash
./ipgeo 8.8.8.8
```
## Domain:
```
./ipgeo example.com
```
## Json:
```bash
./ipgeo --json 8.8.8.8
```
## history:
```bash
./ipgeo --history
```
## Clear Cache:
```bash
./ipgeo --clear-cache
```
## Version:
```bash
./ipgeo --version
```
## Help:
```bash
./ipgeo --help
```
## Disclaimer

Informasi geolocation berbasis IP bersifat perkiraan dan bukan lokasi fisik pasti seseorang.
Gunakan tool hanya terhadap IP/domain yang kamu miliki atau yang memang kamu punya izin untuk periksa. EOF

### Tambah automated test

Bikin:

```bash
cat > geo_test.go <<'EOF'
package main

import (
        "net"
        "testing"
)

func TestResolveIP(t *testing.T) {
        input := "8.8.8.8"

        got, err := resolveTarget(input)
        if err != nil {
                t.Fatalf("unexpected error: %v", err)
        }

        if net.ParseIP(got) == nil {
                t.Fatalf("expected valid IP, got %q", got)
        }
}

func TestResolveInvalidTarget(t *testing.T) {
        _, err := resolveTarget("this-domain-should-not-exist.invalid")

        if err == nil {
                t.Fatal("expected an error for invalid domain")
        }
}
EOF
```
## jalankan:
```bash
go test ./...
```
## kalau muncul:

PASS
ok      ipgeo

## berarti mantap joss

## build final
```bash
gofmt -w *.go
go test ./...
go build -o ipgeo .
cp ipgeo $PREFIX/bin/ipgeo
```
## cek:
```bash
ipgeo --version
ipgeo --help
```
## hanya untuk pembelajaran

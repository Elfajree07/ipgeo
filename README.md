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

## install

```bash
pkg update && pkg upgrade -y
pkg install git -y
pkg install golang
git clone https://github.com/Elfajree07/ipgeo.git
cd ipgeo/v4
```
## Build

```bash
gofmt -w *.go
go build -o ipgeo .
cp ipgeo $PREFIX/bin/ipgeo
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

## cek:
```bash
ipgeo --version
ipgeo --help
```
## hanya untuk pembelajaran

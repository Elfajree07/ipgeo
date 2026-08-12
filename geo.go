package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type GeoData struct {
	Status     string  `json:"status"`
	Message    string  `json:"message"`
	Query      string  `json:"query"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionName"`
	City       string  `json:"city"`
	ISP        string  `json:"isp"`
	Org        string  `json:"org"`
	AS         string  `json:"as"`
	Timezone   string  `json:"timezone"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lon"`
	ReverseDNS string  `json:"reverse_dns,omitempty"`
}

func resolveTarget(target string) (string, error) {
	target = strings.TrimSpace(target)

	if net.ParseIP(target) != nil {
		return target, nil
	}

	ips, err := net.LookupIP(target)
	if err != nil {
		return "", fmt.Errorf("resolve gagal: %w", err)
	}

	for _, ip := range ips {
		if v4 := ip.To4(); v4 != nil {
			return v4.String(), nil
		}
	}

	if len(ips) > 0 {
		return ips[0].String(), nil
	}

	return "", fmt.Errorf("IP tidak ditemukan")
}

func lookupIP(ip string, timeout time.Duration) (*GeoData, error) {
	client := &http.Client{Timeout: timeout}

	url := "http://ip-api.com/json/" + ip +
		"?fields=status,message,query,country,regionName,city,isp,org,as,timezone,lat,lon"

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	var data GeoData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Status != "success" {
		return nil, fmt.Errorf("%s", data.Message)
	}

	if names, err := net.LookupAddr(data.Query); err == nil && len(names) > 0 {
		data.ReverseDNS = strings.TrimSuffix(names[0], ".")
	}

	return &data, nil
}

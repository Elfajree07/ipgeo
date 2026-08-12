package main

import (
	"encoding/json"
	"fmt"
)

func banner() {
	fmt.Println("\033[91m╔══════════════════════════════════════╗\033[0m")
	fmt.Println("\033[91m║          IP GEOLOCATION v4.0        ║\033[0m")
	fmt.Println("\033[91m╚══════════════════════════════════════╝\033[0m")
}

func printTable(data *GeoData, cached bool) {
	fmt.Println()
	fmt.Println("┌──────────────────┬────────────────────────────┐")
	fmt.Printf("│ %-16s │ %-26s │\n", "FIELD", "VALUE")
	fmt.Println("├──────────────────┼────────────────────────────┤")
	fmt.Printf("│ %-16s │ %-26s │\n", "IP", data.Query)
	fmt.Printf("│ %-16s │ %-26s │\n", "Country", data.Country)
	fmt.Printf("│ %-16s │ %-26s │\n", "Region", data.RegionName)
	fmt.Printf("│ %-16s │ %-26s │\n", "City", data.City)
	fmt.Printf("│ %-16s │ %-26s │\n", "ISP", data.ISP)
	fmt.Printf("│ %-16s │ %-26s │\n", "Organization", data.Org)
	fmt.Printf("│ %-16s │ %-26s │\n", "AS", data.AS)
	fmt.Printf("│ %-16s │ %-26s │\n", "Timezone", data.Timezone)
	fmt.Printf("│ %-16s │ %-26.6f │\n", "Latitude", data.Latitude)
	fmt.Printf("│ %-16s │ %-26.6f │\n", "Longitude", data.Longitude)
	fmt.Printf("│ %-16s │ %-26s │\n", "Reverse DNS", data.ReverseDNS)
	fmt.Println("└──────────────────┴────────────────────────────┘")

	if cached {
		fmt.Println("⚡ CACHE HIT")
	} else {
		fmt.Println("✓ Fresh lookup")
	}
}

func printJSON(data *GeoData) {
	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	fmt.Println(string(raw))
}

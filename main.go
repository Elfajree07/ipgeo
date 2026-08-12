package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const version = "4.0.0"

func getData(target string, timeout time.Duration) (*GeoData, bool, error) {
	ip, err := resolveTarget(target)
	if err != nil {
		return nil, false, err
	}

	// Cache berlaku 24 jam.
	if data, ok := readCache(ip, 24*time.Hour); ok {
		return data, true, nil
	}

	data, err := lookupIP(ip, timeout)
	if err != nil {
		return nil, false, err
	}

	_ = writeCache(ip, data)
	_ = addHistory(data)

	return data, false, nil
}

func main() {
	jsonMode := flag.Bool("json", false, "output JSON")
	timeout := flag.Duration("timeout", 10*time.Second, "request timeout")
	history := flag.Bool("history", false, "tampilkan history")
	clear := flag.Bool("clear-cache", false, "hapus cache")
	showVersion := flag.Bool("version", false, "tampilkan versi")

        flag.Usage = func() {
	fmt.Println("IPGeo v" + version)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  ipgeo <IP|domain>")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ipgeo 8.8.8.8")
	fmt.Println("  ipgeo google.com")
	fmt.Println("  ipgeo --json 8.8.8.8")
	fmt.Println("  ipgeo --history")
	fmt.Println("  ipgeo --clear-cache")
	fmt.Println("  ipgeo --version")
}

	flag.Parse()

	if *showVersion {
		fmt.Println("IPGeo v" + version)
		fmt.Println("Go-based IP information tool")
		return
	}

	if *clear {
		if err := clearCache(); err != nil {
			fmt.Println("[!] Gagal menghapus cache:", err)
			return
		}

		fmt.Println("[+] Cache berhasil dihapus.")
		return
	}

	if *history {
		showHistory()
		return
	}

	if flag.NArg() < 1 {
		banner()

		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  ipgeo <IP/domain>")
		fmt.Println("  ipgeo --json <IP/domain>")
		fmt.Println("  ipgeo --history")
		fmt.Println("  ipgeo --clear-cache")
		fmt.Println("  ipgeo --version")
		return
	}

	target := flag.Arg(0)

	start := time.Now()

	data, cached, err := getData(target, *timeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[!] Error:", err)
		os.Exit(1)
	}

	if *jsonMode {
		printJSON(data)
		return
	}

	banner()
	printTable(data, cached)

	fmt.Println()
	fmt.Println("✓ Completed:", time.Since(start).Round(time.Millisecond))
}

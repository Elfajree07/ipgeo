package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type CacheEntry struct {
	Data      GeoData `json:"data"`
	CreatedAt string  `json:"created_at"`
}

func cacheDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".ipgeo-cache"
	}

	return filepath.Join(home, ".cache", "ipgeo")
}

func cachePath(ip string) string {
	return filepath.Join(cacheDir(), ip+".json")
}

func readCache(ip string, maxAge time.Duration) (*GeoData, bool) {
	path := cachePath(ip)

	info, err := os.Stat(path)
	if err != nil {
		return nil, false
	}

	if time.Since(info.ModTime()) > maxAge {
		return nil, false
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}

	var entry CacheEntry

	if json.Unmarshal(raw, &entry) != nil {
		return nil, false
	}

	return &entry.Data, true
}

func writeCache(ip string, data *GeoData) error {
	if err := os.MkdirAll(cacheDir(), 0700); err != nil {
		return err
	}

	entry := CacheEntry{
		Data:      *data,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	raw, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cachePath(ip), raw, 0600)
}

func clearCache() error {
	return os.RemoveAll(cacheDir())
}

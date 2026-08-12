package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type HistoryEntry struct {
	Time string  `json:"time"`
	Data GeoData `json:"data"`
}

func historyPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".ipgeo-history.json"
	}

	return filepath.Join(home, ".local", "share", "ipgeo", "history.json")
}

func addHistory(data *GeoData) error {
	path := historyPath()

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	var history []HistoryEntry

	if raw, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(raw, &history)
	}

	history = append(history, HistoryEntry{
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Data: *data,
	})

	// Simpan maksimal 100 entry.
	if len(history) > 100 {
		history = history[len(history)-100:]
	}

	raw, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, raw, 0600)
}

func showHistory() {
	raw, err := os.ReadFile(historyPath())
	if err != nil {
		println("Belum ada history.")
		return
	}

	var history []HistoryEntry

	if json.Unmarshal(raw, &history) != nil {
		println("History rusak atau tidak valid.")
		return
	}

	println()
	println("TIME                  IP              COUNTRY")
	println("------------------------------------------------")

	for _, item := range history {
		println(item.Time, "  ", item.Data.Query, "   ", item.Data.Country)
	}
}

package main

import "testing"

func TestLoadConfiguration(t *testing.T) {
	filePath := "./config.json"
	config := LoadConfiguration(filePath)
	reportsCount := len(config.Reports)
	if reportsCount == 0 {
		t.Errorf("LoadConfiguration was incorrect, got: %d, want: >0", reportsCount)
	}
}

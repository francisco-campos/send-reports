package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestLoadConfiguration(t *testing.T) {
	filePath := "./config.json"
	config := LoadConfiguration(filePath)
	reportsCount := len(config.Reports)
	if reportsCount == 0 {
		t.Errorf("LoadConfiguration was incorrect, got: %d, want: >0", reportsCount)
	}
}

package main

import (
	"testing"
	"time"
)

func TestIso2Unix(t *testing.T) {
	tests := []struct {
		iso      string
		expected int64
	}{
		{"2025-12-16T10:19:25UTC", 1765880365},
		{"2025-12-16 10:19:25 UTC", 1765880365},
		{"2025-12-16T10:19:25Z", 1765880365},
		{"2025-12-16 10:19:25 Z", 1765880365},
		{"2025-12-16T10:19:25", 1765880365}, // Assuming local timezone is UTC
		{"2025-12-16 10:19:25", 1765880365}, // Assuming local timezone is UTC
	}
	location, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatalf("Failed to load location: %v", err)
	}
	time.Local = location
	for _, test := range tests {
		t.Run(test.iso, func(t *testing.T) {
			result, err := Iso2Unix(test.iso)
			if err != nil {
				t.Errorf("Iso2Unix(%q) returned error: %v", test.iso, err)
			}
			if result != test.expected {
				t.Errorf("Iso2Unix(%q) = %d; want %d", test.iso, result, test.expected)
			}
		})
	}
}

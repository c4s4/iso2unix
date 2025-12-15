package main

import (
	"testing"
)

func TestIso2Unix(t *testing.T) {
	tests := []struct {
		iso     string
		human   bool
		want    int64
		wantErr bool
	}{
		{"2023-10-01T12:00:00Z", false, 1696161600, false},
		{"2023-10-01 12:00:00 UTC", true, 1696161600, false},
		{"invalid-date", false, 0, true},
		{"2023-10-01T12:00:00", false, 0, true},
	}
	for _, tt := range tests {
		got, err := Iso2Unix(tt.iso, tt.human)
		if (err != nil) != tt.wantErr {
			t.Errorf("Iso2Unix(%q, %v) error = %v, wantErr %v", tt.iso, tt.human, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Iso2Unix(%q, %v) = %v, want %v", tt.iso, tt.human, got, tt.want)
		}
	}
}

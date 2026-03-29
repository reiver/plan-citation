package main

import (
	"encoding/hex"
	"strings"
	"testing"
	"time"
)

func TestGenerateUUIDv7_Format(t *testing.T) {

	tests := []struct{
		Name string
		Time time.Time
	}{
		{
			Name: "unix-epoch",
			Time: time.Unix(0, 0),
		},
		{
			Name: "2024-01-01",
			Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Name: "now",
			Time: time.Now(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			result, err := generateUUIDv7(test.Time)
			if nil != err {
				t.Fatalf("unexpected error: %s", err)
			}

			// check overall length: 8-4-4-4-12 = 36 characters
			if expected, actual := 36, len(result); expected != actual {
				t.Errorf("expected length %d but got %d for %q", expected, actual, result)
				return
			}

			// check hyphen positions
			for _, pos := range []int{8, 13, 18, 23} {
				if '-' != result[pos] {
					t.Errorf("expected '-' at position %d but got %q in %q", pos, result[pos], result)
				}
			}

			// check all non-hyphen characters are valid hex
			stripped := strings.ReplaceAll(result, "-", "")
			_, err = hex.DecodeString(stripped)
			if nil != err {
				t.Errorf("UUID contains non-hex characters: %q", result)
			}
		})
	}
}

func TestGenerateUUIDv7_Version(t *testing.T) {

	tests := []struct{
		Name string
		Time time.Time
	}{
		{
			Name: "unix-epoch",
			Time: time.Unix(0, 0),
		},
		{
			Name: "2024-01-01",
			Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Name: "now",
			Time: time.Now(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			result, err := generateUUIDv7(test.Time)
			if nil != err {
				t.Fatalf("unexpected error: %s", err)
			}

			// version nibble is the first character of the third section
			// format: xxxxxxxx-xxxx-Vxxx-xxxx-xxxxxxxxxxxx
			if expected, actual := byte('7'), result[14]; expected != actual {
				t.Errorf("expected version nibble %q but got %q in %q", expected, actual, result)
			}
		})
	}
}

func TestGenerateUUIDv7_Variant(t *testing.T) {

	tests := []struct{
		Name string
		Time time.Time
	}{
		{
			Name: "unix-epoch",
			Time: time.Unix(0, 0),
		},
		{
			Name: "2024-01-01",
			Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Name: "now",
			Time: time.Now(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			result, err := generateUUIDv7(test.Time)
			if nil != err {
				t.Fatalf("unexpected error: %s", err)
			}

			// variant bits are the first character of the fourth section
			// format: xxxxxxxx-xxxx-xxxx-Vxxx-xxxxxxxxxxxx
			// must be 8, 9, a, or b (binary 10xx)
			variantChar := result[19]
			switch variantChar {
			case '8', '9', 'a', 'b':
				// good
			default:
				t.Errorf("expected variant character to be 8, 9, a, or b but got %q in %q", variantChar, result)
			}
		})
	}
}

func TestGenerateUUIDv7_Timestamp(t *testing.T) {

	tests := []struct{
		Name string
		Time time.Time
	}{
		{
			Name: "unix-epoch",
			Time: time.Unix(0, 0),
		},
		{
			Name: "2024-01-01",
			Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Name: "2038-problem",
			Time: time.Date(2038, 1, 19, 3, 14, 7, 0, time.UTC),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			result, err := generateUUIDv7(test.Time)
			if nil != err {
				t.Fatalf("unexpected error: %s", err)
			}

			// extract the timestamp: first 12 hex chars (bytes 0-5) = 48 bits
			stripped := strings.ReplaceAll(result, "-", "")
			tsHex := stripped[0:12]

			tsBytes, err := hex.DecodeString(tsHex)
			if nil != err {
				t.Fatalf("could not decode timestamp hex %q: %s", tsHex, err)
			}

			var ms uint64
			for _, b := range tsBytes {
				ms = (ms << 8) | uint64(b)
			}

			expected := uint64(test.Time.UnixMilli())
			if expected != ms {
				t.Errorf("expected timestamp %d but got %d in %q", expected, ms, result)
			}
		})
	}
}

func TestGenerateUUIDv7_Uniqueness(t *testing.T) {

	now := time.Now()

	seen := make(map[string]struct{}, 100)
	for range 100 {
		result, err := generateUUIDv7(now)
		if nil != err {
			t.Fatalf("unexpected error: %s", err)
		}
		if _, exists := seen[result]; exists {
			t.Fatalf("duplicate UUID generated: %q", result)
		}
		seen[result] = struct{}{}
	}
}

package base64

import (
	"testing"

	"github.com/Diya-Project/base-64/lib"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"A", "QQ=="},
		{"AB", "QUI="},
		{"ABC", "QUJD"},
		{"Hello", "SGVsbG8="},
		{"Base64", "QmFzZTY0"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := lib.Encode(tc.input)
		if result != tc.expected {
			t.Errorf("Encode(%q) = %q; want %q", tc.input, result, tc.expected)
		}
	}
}

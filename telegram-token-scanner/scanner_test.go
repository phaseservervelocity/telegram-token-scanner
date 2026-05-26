package main

import (
	"testing"
)

func TestAddressDetection(t *testing.T) {
	scanner := NewScanner()

	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "Check this token: 0x1234567890abcdef1234567890abcdef12345678",
			expected: 1,
		},
		{
			input:    "Multiple: 0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa and 0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			expected: 2,
		},
		{
			input:    "No token here",
			expected: 0,
		},
		{
			input:    "Invalid 0x123",
			expected: 0,
		},
	}

	for _, tt := range tests {
		matches := scanner.addressPattern.FindAllString(tt.input, -1)
		if len(matches) != tt.expected {
			t.Errorf("Expected %d matches, got %d for input: %s", tt.expected, len(matches), tt.input)
		}
	}
}

func TestFormatAddressInfo(t *testing.T) {
	scanner := NewScanner()
	addr := "0x1234567890abcdef1234567890abcdef12345678"
	result := scanner.formatAddressInfo(1, addr)

	if !contains(result, addr) {
		t.Errorf("Expected address %s in formatted output", addr)
	}
	if !contains(result, "Etherscan") {
		t.Error("Expected Etherscan link in formatted output")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			contains(s[1:], substr)))
}
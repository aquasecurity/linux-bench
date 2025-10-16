package main

import (
	"testing"
)

func TestGetPlatformVersion(t *testing.T) {
	tests := []struct {
		output   string
		platform string
		expected string
	}{
		{"version_id=7.6", "rhel", "7"},
		{"version_id=18.04", "ubuntu", "18"},
		{"version_id=2023", "amzn", "2023"},
		{"version_id=2", "amzn", "2"},
		{"version_id=foobar", "debian", ""},
		{"version_id=3.0.12", "azurelinux3", ""},
	}
	for _, tc := range tests {
		got := getPlatformVersion(tc.output, tc.platform)
		if got != tc.expected {
			t.Errorf("getPlatformVersion(%q, %q) = %q; want %q",
				tc.output, tc.platform, got, tc.expected)
		}
	}
}

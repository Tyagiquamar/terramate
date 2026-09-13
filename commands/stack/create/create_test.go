// Copyright 2024 Terramate GmbH
// SPDX-License-Identifier: MPL-2.0

package create

import (
	"testing"
)

func TestHasTFExt(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"main.tf", true},
		{"main.tf.json", true},
		{"main.tofu", true},
		{"main.tofu.json", true},
		{"README.md", false},
		{"config.json", false},
		{"main.tf.txt", false},
	}

	for _, tt := range tests {
		got := hasTFExt(tt.filename)
		if got != tt.expected {
			t.Errorf("hasTFExt(%q) = %v; want %v", tt.filename, got, tt.expected)
		}
	}
}

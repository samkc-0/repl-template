package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{
			input: " hello ",
			want:  []string{"hello"},
		},
		{
			input: "hello world",
			want:  []string{"hello", "world"},
		},
		{
			input: "hello world",
			want:  []string{"hello", "world"},
		},
		{
			input: "Hell0 wOrld",
			want:  []string{"hell0", "world"},
		},
	}
	for _, tt := range cases {
		got := cleanInput(tt.input)
		if !stringSlicesAreEqual(got, tt.want) {
			t.Errorf("cleanInput(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func stringSlicesAreEqual(a, b []string) bool {
	// true when both slices contain the same strings in the same order
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

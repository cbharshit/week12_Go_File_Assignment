package main

import "testing"

func TestConcat(t *testing.T) {
	testCases := []struct {
		a, b     string
		expected string
	}{
		{"Hello", "World", "HelloWorld"},
		{"Go", "Lang", "GoLang"},
		{"", "Empty", "Empty"},
		{"Pre", "", "Pre"},
		{"", "", ""},
		{"Hello", "World", "HelloWorld!"},
		{"Go", "Lang", "GoLangu"},
	}

	for _, tc := range testCases {
		result := Concat(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Concat(%q, %q) = %q; expected %q", tc.a, tc.b, result, tc.expected)
		}
	}
}

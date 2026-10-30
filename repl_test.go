package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "  ALL  IN  UPPERCASE   ",
			expected: []string{"all", "in", "uppercase"},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " Using RaNdOMcAsE oN wORdS",
			expected: []string{"using", "randomcase", "on", "words"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length: %d, got %d", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected string: %q, got %q", expectedWord, word)
			}
		}
	}
}

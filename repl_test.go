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
		// add more cases here
		{
			input:    "Go is Awesome",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    "   MULTIPLE    SPACES   ",
			expected: []string{"multiple", "spaces"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("For input %q, expected length %d but got %d", c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("For input %q, expected word %q but got %q", c.input, expectedWord, word)
			}
		}
	}
}

func TestGetCommands(t *testing.T) {
	commands := getCommands()

	expectedCommands := []string{"exit", "help"}

	for _, cmdName := range expectedCommands {
		if _, exists := commands[cmdName]; !exists {
			t.Errorf("Expected command %q to be present in commands map", cmdName)
		}
	}
}

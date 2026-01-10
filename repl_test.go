package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
		input: " hello world ",
		expected: []string{"hello", "world"},
		},{
		input: "        haha you suck         man",
		expected: []string{"haha", "you", "suck", "man"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(c.expected) != len(actual) {
		t.Errorf("The length of the actual slice does not match the length of the expected slice!\nInput: %s\nExpected: %d\nactual: %d", c.input, len(c.expected), len(actual))
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("The actual word does not match the expected word!\nActual: %s\nExpected: %s", word, expectedWord)
			}
		}	
	}
}

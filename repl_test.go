package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "        haha you suck         man",
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

func TestMapNavigatorReset(t *testing.T) {
	nav := NewMapNavigator()

	// Set some values
	nav.offset = 40
	nav.total = 1089

	// Reset
	nav.Reset()

	if nav.offset != 0 {
		t.Errorf("Expected offset to be 0 after reset, got %d", nav.offset)
	}

	if nav.total != 0 {
		t.Errorf("Expected total to be 0 after reset, got %d", nav.total)
	}

	if nav.cache == nil {
		t.Error("Expected cache to still exist after reset")
	}
}

func TestMapNavigatorInitialization(t *testing.T) {
	nav := NewMapNavigator()

	if nav == nil {
		t.Fatal("Expected NewMapNavigator to return non-nil navigator")
	}

	if nav.offset != 0 {
		t.Errorf("Expected initial offset to be 0, got %d", nav.offset)
	}

	if nav.total != 0 {
		t.Errorf("Expected initial total to be 0, got %d", nav.total)
	}

	if nav.cache == nil {
		t.Error("Expected cache to be initialized")
	}
}

func TestCommandListInitialized(t *testing.T) {
	requiredCommands := []string{"exit", "help", "map", "mapb"}

	for _, cmdName := range requiredCommands {
		cmd, exists := CommandsList[cmdName]
		if !exists {
			t.Errorf("Expected command '%s' to exist in CommandsList", cmdName)
			continue
		}

		if cmd.name != cmdName {
			t.Errorf("Expected command name to be '%s', got '%s'", cmdName, cmd.name)
		}

		if cmd.description == "" {
			t.Errorf("Expected command '%s' to have a description", cmdName)
		}

		if cmd.callback == nil {
			t.Errorf("Expected command '%s' to have a callback function", cmdName)
		}
	}

	// Verify global navigator is initialized
	if navigator == nil {
		t.Error("Expected global navigator to be initialized")
	}
}

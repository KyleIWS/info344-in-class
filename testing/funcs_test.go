package testing

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "empty string",
			input:          "",
			expectedOutput: "",
		},
		{
			name:           "single character",
			input:          "a",
			expectedOutput: "a",
		},
		{
			name:           "two-chars",
			input:          "ab",
			expectedOutput: "ba",
		},
		{
			name:           "palindrome",
			input:          "stressed",
			expectedOutput: "desserts",
		},
		{
			name:           "Unicode check",
			input:          "木曜日",
			expectedOutput: "日曜木",
		},
	}

	for _, c := range cases {
		if output := Reverse(c.input); output != c.expectedOutput {
			t.Errorf("input was %s got %s expected %s\n", c.input, output, c.expectedOutput)
		}
	}
}

func TestGetGreeting(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Empty string",
			input:          "",
			expectedOutput: "Hello, World!",
		},
		{
			name:           "Non-empty string",
			input:          "Kyle",
			expectedOutput: "Hello, Kyle!",
		},
	}

	for _, c := range cases {
		if output := GetGreeting(c.input); output != c.expectedOutput {
			t.Errorf("Case %s: Output: %s ExpectedOutput: %s\n", c.name, output, c.expectedOutput)
		}
	}
}

func TestParseSize(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput *Size
	}{
		{
			name:           "Empty",
			input:          "",
			expectedOutput: &Size{},
		},
	}

	for _, c := range cases {
		if output := ParseSize(c.input); output.Height != c.expectedOutput.Height || output.Width != c.expectedOutput.Width {
			t.Errorf("Case %s: %v", c.name, c.expectedOutput)
		}
	}
}

func TestLateDaysConsume(t *testing.T) {
	ld := NewLateDays()
	for i := 3; i > -10; i-- {
		expectedLateDays := i
		if expectedLateDays < 0 {
			expectedLateDays = 0
		}
		if output := ld.Consume("test"); output != expectedLateDays {
			t.Errorf("Iteration %d: got %d but expected %d", i, output, expectedLateDays)
		}
	}
}

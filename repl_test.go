package main

import "testing"

func testCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " the itsy bitsy spidEr",
			expected: []string{"the", "itsy", "bitsy", "spider"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Error("Expected", c.expected, "but got", c.input)
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Error("Expected", actual[i], "but got", c.expected[i])
			}
		}
	}
}

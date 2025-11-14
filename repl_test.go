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
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " spuds AND kumara FoR BrEaKfAst",
			expected: []string{"spuds", "and", "kumara", "for", "breakfast"},
		},
		{
			input: "450 tImEs 100 eqUalS 45000",
			expected: []string{"450", "times", "100", "equals", "45000"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of slice does not match expected output")
		} 
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word mismatch found!!")
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}

}

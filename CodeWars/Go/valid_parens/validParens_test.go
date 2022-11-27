package validParens

import "testing"

type testCase struct {
	given    string
	expected bool
}

var testCases = []testCase{
	testCase{
		given:    "()",
		expected: true,
	},
	testCase{
		given:    ")(()))",
		expected: false,
	},
	testCase{
		given:    "(",
		expected: false,
	},
	testCase{
		given:    "(())((()())())",
		expected: true,
	},
}

func TestValidParens(t *testing.T) {
	for _, e := range testCases {
		if res := Validate(e.given); res != e.expected {
			t.Fatalf(`Validate("%s") returned %t, wanted %t`, e.given, res, e.expected)
		}
	}
}

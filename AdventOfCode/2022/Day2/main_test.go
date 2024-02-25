package main

import "testing"

var test1Input = `a y
b x
c z`

var testCase = struct {
	given  string
	wanted int
}{
	test1Input,
	12,
}

func TestHelloWorld(t *testing.T) {
	if res := Getpart2(testCase.given); res != testCase.wanted {
		t.Fatalf("Oi ,we got %d, but wanted %d", res, testCase.wanted)
	}
}

// Package hashtaggen but this time its for testing
package hashtaggen

import (
	"errors"
	"testing"
)

type testCase struct {
	arg            string
	expectedResult string
	err            error
}

var testCases = []testCase{
	{
		arg:            "",
		expectedResult: "",
		err:            errors.New("Invalid string"),
	},
	{
		arg:            "   Hello there   ",
		expectedResult: "#HelloThere",
		err:            nil,
	},
	{
		arg:            "Hello there  ",
		expectedResult: "#HelloThere",
		err:            nil,
	},
	{
		arg:            "Hello there",
		expectedResult: "#HelloThere",
		err:            nil,
	},
}

func TestHashTagGen(t *testing.T) {
	for _, e := range testCases {
		err, res := HashTagGen(e.arg)
		if res != e.expectedResult {
			t.Fatalf(`HashTagGen("%s") failed. Got "%s", wanted "%s"`, e.arg, res, e.expectedResult)
		}
		if e.err != nil && err == nil {
			t.Fatalf(`Should have errored, wanted "%s", got "%s"`, e.err, err)
		}
	}
}

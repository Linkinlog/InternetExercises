package main

import "testing"

type testString struct {
	given  string
	wanted string
}

var tests = []testString{
	testString{
		given:  "AATCG",
		wanted: "TTAGC",
	},
	testString{
		given:  "AAAA",
		wanted: "TTTT",
	},
	testString{
		given:  "ATTGC",
		wanted: "TAACG",
	},
	testString{
		given:  "GTAT",
		wanted: "CATA",
	},
}

func TestDNAStrand(t *testing.T) {
	for _, e := range tests {
		tester(e.given, e.wanted, t)
	}
}

func tester(given string, wanted string, t *testing.T) error {
	ret := DNAStrand(given)
	if ret != wanted {
		t.Fatalf("Test 1: DNAStrand(%s) returned: %s. Wanted: %s", given, ret, wanted)
	}
	return nil
}

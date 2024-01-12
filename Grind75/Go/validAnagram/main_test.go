package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := map[string]struct {
		s              string
		e              string
		expectedResult bool
	}{
		"aacc": {
			s:              "aacc",
			e:              "ccac",
			expectedResult: false,
		},
		"anagram": {
			s: "anagram",
			e: "nagaram",
			expectedResult: true,
		},
		"rat": {
			s: "rat",
			e: "car",
			expectedResult: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if val := isAnagram(tt.s, tt.e); val != tt.expectedResult {
				t.Fatalf("expexted %t, got %t", tt.expectedResult, val)
			}
		})
	}
}

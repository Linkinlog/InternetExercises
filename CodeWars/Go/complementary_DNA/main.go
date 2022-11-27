// Package main includes utilitiez for DNA strand testing
package main

import (
	"fmt"
	"os"
)

func main() {
	var str = os.Args[1]
	fmt.Print(DNAStrand(str))
}

// DNAStrand returns the complementary DNA strand for the given strand
func DNAStrand(str string) string {
	var newStrand string
	m := map[rune]rune{
		'A': 'T',
		'T': 'A',
		'G': 'C',
		'C': 'G',
	}

	for _, e := range str {
		newStrand = newStrand + string(m[e])
	}

	return newStrand
}

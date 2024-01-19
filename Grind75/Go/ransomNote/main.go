package main

func canConstruct(ransomNote string, magazine string) bool {
	// put the letters of magazine into a hash
	// go through the letters of the ransomNote, for each character check if it exists, if so subtract one from the amount, else false
	// once all through, true
	magazineRuneOccurs := map[rune]int{}
	for _, r := range magazine {
		if _, ok := magazineRuneOccurs[r]; ok {
			magazineRuneOccurs[r] += 1
		} else {
			magazineRuneOccurs[r] = 1
		}
	}

	for _, r := range ransomNote {
		if v, ok := magazineRuneOccurs[r]; !ok || v < 1 {
			return false
		} else if ok {
			magazineRuneOccurs[r] -= 1
		}
	}

	return true
}

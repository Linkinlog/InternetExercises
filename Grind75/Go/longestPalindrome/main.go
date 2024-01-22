package main

func longestPalindrome(s string) int {
	ans := 0
	h := make(map[rune]int, 0)
	for _, c := range s {
		h[c] += 1
	}

	for _, f := range h {
		if f%2 == 0 {
			ans += f
		} else {
			ans += f - 1
		}
	}

	hasOddCount := false
	for _, v := range h {
		if v&1 == 1 {
			hasOddCount = true
			break
		}
	}

	if hasOddCount {
		return ans + 1
	}
	return ans
}

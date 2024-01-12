package main

func isValid(s string) bool {
	stack := []rune{}
	pMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	openSet := map[rune]bool{
		'(': true,
		'{': true,
		'[': true,
	}

	for _, ch := range s {
		if openSet[ch] {
			stack = append(stack, ch)
		} else if closing, exists := pMap[ch]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

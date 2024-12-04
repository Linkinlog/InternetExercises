package main

var chars []rune = []rune{'x', 'm', 'a', 's'}

type direction int

const (
	_ direction = iota
	up
	up_right
	up_left
	down
	down_right
	down_left
	left
	right
)

var directions = []direction{up, up_right, up_left, down, down_right, down_left, left, right}

func check(needles []byte, haystack []string, d direction, x, y int) bool {
	for _, needle := range needles {
		switch d {
		case up:
			y = y - 1
		case up_left:
			y = y - 1
			x = x - 1
		case up_right:
			y = y - 1
			x = x + 1
		case down:
			y = y + 1
		case down_right:
			y = y + 1
			x = x + 1
		case down_left:
			y = y + 1
			x = x - 1
		case left:
			x = x - 1
		case right:
			x = x + 1
		}

		if x < 0 || y < 0 || y >= len(haystack) || x >= len(haystack[y]) {
			return false
		}

		if haystack[y][x] != needle {
			return false
		}
	}

	return true
}

func part1(input []string) any {
	sum := 0
	for i, line := range input {
		for ii, char := range line {
			if char == 'X' {
				for _, dir := range directions {
					if check([]byte{'M', 'A', 'S'}, input, dir, ii, i) {
						sum += 1
					}
				}
			}
		}
	}

	return sum
}

func check2(haystack []string, x, y int) bool {
	if x < 1 || y < 1 || y >= len(haystack)-1 || x >= len(haystack[y])-1 {
		return false
	}

	if y-1 < 0 || x-1 < 0 || y+1 >= len(haystack) || x+1 >= len(haystack[y]) {
		return false
	}

	topLeft := haystack[y-1][x-1]
	topRight := haystack[y-1][x+1]
	bottomLeft := haystack[y+1][x-1]
	bottomRight := haystack[y+1][x+1]
	count := 0


	if topLeft == 'M' && bottomRight == 'S' {
		count += 1
	}

	if topRight == 'M' && bottomLeft == 'S' {
		count += 1
	}

	if bottomLeft == 'M' && topRight == 'S' {
		count += 1
	}

	if bottomRight == 'M' && topLeft == 'S' {
		count += 1
	}

	return count == 2
}

func part2(input []string) any {
	sum := 0
	for i, line := range input {
		for ii, char := range line {
			if char == 'A' {
				if check2(input, ii, i) {
					sum += 1
				}
			}
		}
	}

	return sum
}

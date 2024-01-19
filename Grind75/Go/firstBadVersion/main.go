package main

import (
	"fmt"
	"math"
)

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return version >= 2
}

func firstBadVersion(n int) int {
	sqrt := math.Floor(math.Sqrt(float64(n)))

	cursor := 0
	for ; cursor < n; cursor += int(sqrt) {
		if isBadVersion(cursor) {
			break
		}
	}

	for cursor -= int(sqrt); cursor <= n; cursor++ {
		if isBadVersion(cursor) {
			return cursor
		}
	}
	return 0
}

func main() {
	fmt.Println(firstBadVersion(2))
}

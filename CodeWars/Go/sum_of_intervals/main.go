// Package main provides utilities for summing an array of integers
package main

import "fmt"

func main() {
	fmt.Println(SumOfIntervals([][2]int{{1, 5}}))
	fmt.Println(SumOfIntervals([][2]int{{1, 5}, {6, 10}}))
	fmt.Println(SumOfIntervals([][2]int{{1, 5}, {1, 5}}))
	fmt.Println(SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}}))
	// {[][2]int{{1, 5}}, 4},
	// {[][2]int{{1, 5}, {6, 10}}, 8},
	// {[][2]int{{1, 5}, {1, 5}}, 4},
	// {[][2]int{{1, 4}, {7, 10}, {3, 5}}, 7},
}

// SumOfIntervals digests a slice of int arrays and returns the sum of the difference between said ints
func SumOfIntervals(intervals [][2]int) int {
	var total int
	var y int
	var x int
	var lastY int
	var lastX int
	if len(intervals) <= 1 {
		return intervals[0][1] - intervals[0][0]
	}
	for i := 0; i < len(intervals); i++ {
		x, y = intervals[i][0], intervals[i][1]
		if x <= lastX && lastX >= y || lastX == x {
			fmt.Println("BREAK")
			break
		}
		if lastY == 0 || lastY > y {
			lastY = y
		}
		if lastX == 0 || lastX > x {
			lastX = x
		}
		fmt.Printf(`i: %d, lastX: %d`, i, lastX)
		fmt.Println()
		fmt.Printf(`i: %d, lastY: %d`, i, lastY)
		fmt.Println()
		if x < lastY {
			y -= (lastY - x)
		}

		total += (y - x)
	}
	fmt.Println("Here")
	return total
}

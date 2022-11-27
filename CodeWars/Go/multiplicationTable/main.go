// Package main provides utilities for visualizing multiplication tables
package main

import "fmt"

func main() {
	fmt.Print(MultiplicationTable(3))
}

// MultiplicationTable creates a NxN multiplication table, of size size
func MultiplicationTable(size int) [][]int {
	res := make([][]int, size)
	f := 1
	for i := 0; f < size+1; {
		res[i] = append(res[i], (i+1)*f)
		i++
		if i == size {
			f++
			i = 0
		}
	}
	return res
}

func Comparison(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		for x := 1; x < size+1; x++ {
			res[i] = append(res[i], (i+1)*x)
		}
	}
	return res
}

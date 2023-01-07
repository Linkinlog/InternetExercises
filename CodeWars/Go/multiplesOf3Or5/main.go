package main

import "fmt"

func main() {
	fmt.Println(Multiple3And5(100))
}

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in. Additionally, if the number is negative, return 0 (for languages that do have them).

Note: If the number is a multiple of both 3 and 5, only count it once.
*/

func Multiple3And5(number int) int {
	sum := 0

	for i := 1; i < number; i++ {
		if (i%5 == 0) || (i%3 == 0) {
			sum += i
		}
	}

	return sum
}

func Multiple3And52(number int) int {
	n3 := (number - 1) / 3
	n5 := (number - 1) / 5
	n15 := (number - 1) / 15
	return (3*n3*(n3+1) + 5*n5*(n5+1) - 15*n15*(n15+1)) / 2
}

func Multiple3And53(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}

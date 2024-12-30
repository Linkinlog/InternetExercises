package main

import "fmt"

func main() {
	test := []int{3, 3}

	res := twoSum(test, 6)
	fmt.Println(res)
}

func twoSum(nums []int, target int) (output []int) {
	for i, num := range nums {
		for j, num2 := range nums {
			if i != j && num+num2 == target {
				return append(output, i, j)
			}
		}
	}
	return nil
}

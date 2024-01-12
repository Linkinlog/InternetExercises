// main does things
package main

func twoSum(nums []int, target int) []int {
	numbies := make(map[int]int, 0)
	for i, num := range nums {
		goalNumber := target - num
		if idx, ok := numbies[goalNumber]; ok {
			return []int{i, idx}
		}
		numbies[num] = i
	}
	return []int{}
}

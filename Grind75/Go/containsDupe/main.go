package main

func containsDuplicate(nums []int) bool {
	hashy := map[int]int{}
	for _, i := range nums {
		if _,ok := hashy[i]; ok {
			return true
		}
		hashy[i] = i
	}
	return false
}

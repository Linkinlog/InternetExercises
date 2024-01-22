package main


func majorityElement(nums []int) int {
	// make hash
	// fill hash
	// profit?
	hash := map[int]int{}

	for _, n := range nums {
		hash[n] += 1
	}
	majorE := 0
	max := 0
	for k, v := range hash {
		if v > max {
			max = v
			majorE = k
		}
	}
	return majorE
}

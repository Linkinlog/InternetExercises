package main

func maxSubArray(nums []int) int {
	maxSum := 0
	currentSum := 0
	for _, n := range nums {
		if n+currentSum > currentSum {
			currentSum = n + currentSum
		} else {
			if n > currentSum {
				currentSum = n
			}
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

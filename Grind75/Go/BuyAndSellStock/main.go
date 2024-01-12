package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
    
    for i := 1; i < len(prices); i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if (prices[i] - minPrice) > profit {
            profit = prices[i] - minPrice
        }
    }
    
    return profit
}

func main() {
	if val := maxProfit([]int{7,1,5,3,6,4}); val == 5{
		fmt.Println("passed")
	} else {
		fmt.Println(val)
	}
}

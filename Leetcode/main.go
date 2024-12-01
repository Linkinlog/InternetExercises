package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	w []int
}

func Constructor(w []int) Solution {
	return Solution{
		w: w,
	}
}

/*
You need to implement the function pickIndex(), which randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it. The probability of picking an index i is w[i] / sum(w).

For example, if w = [1, 3], the probability of picking index 0 is 1 / (1 + 3) = 0.25 (i.e., 25%), and the probability of picking index 1 is 3 / (1 + 3) = 0.75 (i.e., 75%).
*/

func (this *Solution) PickIndex() int {
    total := 0

    for _, v := range this.w {
        total += v
    }

    random := rand.Intn(total)
    if random == 0 { random = 1 }

    cursor := 0
    for i := range this.w {
        cursor += this.w[i]
        if cursor >= random {
            return i
        }
    }
    return 0
}

func main() {
	w := []int{1, 3}

	sol := Constructor(w)

	fmt.Println(sol.PickIndex())
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

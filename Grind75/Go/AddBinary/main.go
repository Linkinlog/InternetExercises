package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	res := ""
	remainder := 0

	if len(a) < len(b) {
		a,b = b,a
	}

	for i := 0; i < len(a); i++ {
		x := getDigit(a, len(a)-i-1)
		y := getDigit(b, len(b)-i-1)

		sum := x + y + remainder
		res = fmt.Sprintf("%d%s", sum%2, res)
		remainder = sum/2
	}

	if remainder == 1 {
		res = fmt.Sprintf("1%s", res)
	}

	return res

}

func getDigit(num string, index int) int {
	if index >= 0 && len(num) > index {
		i, _ := strconv.Atoi(string(num[index]))
		return i
	}
	return 0
}

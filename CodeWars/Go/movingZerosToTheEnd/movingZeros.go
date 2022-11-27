package movingZeros

import "fmt"

func MoveZeros(arr []int) []int {
	for i, e := range arr {
		fmt.Print(e)
		fmt.Print(arr[:i+1])
		fmt.Print(" ")
		fmt.Print(arr[i+2:])
		fmt.Println()
		if e == 0 {
			arr = append(arr[:i+1], arr[i+2:]...)
		}
	}
	return arr
}

package movingZeros

func MoveZeros(arr []int) []int {
	var array_len = len(arr)
	var last_non_zero_index = 0
	for i := 0; i < array_len; i++ {
		if arr[i] != 0 {
			arr[last_non_zero_index], arr[i] = arr[i], arr[last_non_zero_index]
			last_non_zero_index++
		}
	}
	return arr
}

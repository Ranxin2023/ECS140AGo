package min

// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	// TODO: implement this function.
	if len(arr) == 0 {
		return 0
	}
	var min = arr[0]
	for i := range arr {
		if min > arr[i] {
			min = arr[i]
		}

	}
	return min
}

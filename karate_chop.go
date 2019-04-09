package kata

// chop searches the target in the sorted array source
func chop(target int, source []int) int {
	min := 0
	max := len(source) - 1
	for min <= max {
		mid := min + (max-min)/2
		if source[mid] == target {
			return mid
		} else if source[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}

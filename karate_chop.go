package kata

// chop searches the target in the sorted array source
func chop(target int, source []int) int {
	min := 0
	max := len(source) - 1
	return half(min, max, target, source)
}

// half searches the target at mid recursivly in sorted array source
// it divides the search area by two at each call
func half(min, max, target int, source []int) int {
	if min > max {
		return -1
	}
	mid := min + (max-min)/2
	if source[mid] == target {
		return mid
	} else if source[mid] < target {
		return half(mid+1, max, target, source)
	} else {
		return half(min, mid-1, target, source)
	}
}

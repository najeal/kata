package karatechop

// RecursiveChop implements a recursive chop method
type RecursiveChop struct{}

// Chop searches the target in the sorted array source
func (rc *RecursiveChop) Chop(target int, source []int) int {
	min := 0
	max := len(source) - 1
	return rc.half(min, max, target, source)
}

// half searches the target at mid recursivly in sorted array source
// it divides the search area by two at each call
func (rc *RecursiveChop) half(min, max, target int, source []int) int {
	if min > max {
		return -1
	}
	mid := min + (max-min)/2
	if source[mid] == target {
		return mid
	} else if source[mid] < target {
		return rc.half(mid+1, max, target, source)
	} else {
		return rc.half(min, mid-1, target, source)
	}
}

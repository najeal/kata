package karatechop

// SecondRecursiveChop implements a recursive chop method
type SecondRecursiveChop struct{}

// Chop searches the target in the sorted array source
func (rc *SecondRecursiveChop) Chop(target int, source []int) int {
	mid := len(source) / 2
	max := len(source) - 1
	return rc.half(mid, max, target, source)
}

// half searches the target at mid recursivly in sorted array source
// it divides the search area by two at each call
func (rc *SecondRecursiveChop) half(mid, max, target int, source []int) int {
	if mid > max {
		return -1
	}
	if source[mid] == target {
		return mid
	} else if source[mid] < target {
		return rc.half((mid+(max-mid)/2)+1, max, target, source)
	} else {
		return rc.half(mid/2, mid-1, target, source)
	}
}

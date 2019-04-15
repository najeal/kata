package karatechop

const (
	// NotFound is used when target is not found in array
	NotFound = -1
)

// FunctionalChop implements a functional chop method
type FunctionalChop struct{}

// Chop searches the target in the sorted array source
func (rc *FunctionalChop) Chop(target int, source []int) int {
	min := 0
	length := len(source)
	max := length - 1
	mid := min + (max-min)/2
	if min > max {
		return NotFound
	}
	if source[mid] == target {
		return mid
	} else if source[mid] < target {
		return rc.sum(mid+1, rc.Chop(target, source[mid+1:length]))
	} else {
		return rc.sum(min, rc.Chop(target, source[0:mid]))
	}
}

func (rc *FunctionalChop) sum(track, res int) int {
	if res != NotFound {
		return track + res
	}
	return res
}

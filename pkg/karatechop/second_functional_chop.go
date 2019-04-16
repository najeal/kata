package karatechop

// SecondFunctionalChop implements an iterative chop method
type SecondFunctionalChop struct{}

// Chop searches the target in the sorted array source
func (rc *SecondFunctionalChop) Chop(target int, source []int) int {
	length := len(source)
	if length == 0 {
		return rc.emptySource()
	}
	if length == 1 {
		return rc.oneElemSource(target, source)
	}
	return rc.multipleElemSource(length, target, source)
}

func (rc *SecondFunctionalChop) emptySource() int {
	return NotFound
}

func (rc *SecondFunctionalChop) oneElemSource(target int, source []int) int {
	if source[0] == target {
		return 0
	}
	return NotFound
}

func (rc *SecondFunctionalChop) multipleElemSource(length, target int, source []int) int {
	mid := length / 2
	if source[mid] == target {
		return mid
	} else if source[mid] > target {
		return rc.sum(0, rc.Chop(target, source[0:mid]))
	}
	return rc.sum(mid, rc.Chop(target, source[mid:length]))
}

func (rc *SecondFunctionalChop) sum(track, res int) int {
	if res != NotFound {
		return track + res
	}
	return res
}

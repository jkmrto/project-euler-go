package fibonacci

func GetSequenceUpToLimit(limit int) []int {
	return doGetSequenceUpToLimit([]int{1, 2}, limit)
}

func doGetSequenceUpToLimit(sequence []int, limit int) []int {
	old := sequence[len(sequence)-2]
	current := sequence[len(sequence)-1]
	next := old + current
	if next < limit {
		return doGetSequenceUpToLimit(append(sequence, next), limit)
	} else {
		return sequence
	}

}

package utils

func Intersect(I, J []int) []int {
	intersection := []int{Max(I[0], J[0]), Min(I[1], J[1])}
	if intersection[1] > intersection[0] { // Not empty
		return intersection
	}
	return []int{}
}

func Substract(a, b []int) [][]int { // a - b
	// May return 0, 1 or 2 intervals (2 when b includes a)
	r := [][]int{}
	before := []int{a[0], Min(b[0]-1, a[1])}
	after := []int{Max(b[1]+1, a[0]), a[1]}
	if before[1] > before[0] { // If non empty
		r = append(r, before)
	}
	if after[1] > after[0] {
		r = append(r, after)
	}
	return r
}

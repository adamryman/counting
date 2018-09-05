package counting

// Sort uses counting sort to sort the input slice. Max is the highest value in
// the slice, and by default 0 is the smallest
// TODO: Consider Adding minimum
func Sort(input []int, max int) []int {
	count := make([]int, max, max)
	for _, v := range input {
		count[v]++
	}

	prev := 0
	for i, v := range count {
		count[i] = v + prev
		prev = count[i]
	}

	var index int
	out := make([]int, len(input), len(input))
	for _, v := range input {
		index = count[v]
		out[index-1] = v
		count[v]--
	}

	return out
}

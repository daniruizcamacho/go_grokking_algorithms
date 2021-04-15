package quicksort

func quicksortIterative(values []int) []int {
	if len(values) < 2 {
		return values
	}

	stack := make([]int, len(values))
	top := -1

	top++
	stack[top] = 0
	top++
	stack[top] = len(values) - 1

	for top >= 0 {
		endIndex := stack[top]
		top--
		startIndex := stack[top]
		top--

		pivot := partition(values, startIndex, endIndex)

		if pivot-1 > startIndex {
			top++
			stack[top] = startIndex
			top++
			stack[top] = pivot - 1
		}

		if pivot+1 < endIndex {
			top++
			stack[top] = pivot + 1
			top++
			stack[top] = endIndex
		}
	}

	return values
}

func partition(values []int, startIndex, endIndex int) int {
	x := values[endIndex]
	i := startIndex - 1

	for j := startIndex; j < endIndex; j++ {
		if values[j] <= x {
			i++
			values[i], values[j] = values[j], values[i]
		}
	}
	i++
	values[i], values[endIndex] = values[endIndex], values[i]

	return i
}

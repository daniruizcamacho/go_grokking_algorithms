package quicksort

func quicksortFirstElementPivot(values []int) []int {
	var less, greater, result []int
	valuesCopy := make([]int, len(values))
	copy(valuesCopy, values)

	if len(valuesCopy) < 2 {
		return valuesCopy
	}

	pivot := valuesCopy[0]
	for _, value := range valuesCopy[1:] {
		if value <= pivot {
			less = append(less, value)
		}

		if value > pivot {
			greater = append(greater, value)
		}
	}

	result = append(result, quicksortFirstElementPivot(less)...)
	result = append(result, pivot)
	result = append(result, quicksortFirstElementPivot(greater)...)
	return result
}

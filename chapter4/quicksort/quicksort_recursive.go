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

func quicksortLastElementPivot(values []int) []int {
	var less, greater, result []int
	valuesCopy := make([]int, len(values))
	copy(valuesCopy, values)

	if len(valuesCopy) < 2 {
		return valuesCopy
	}

	pivot := valuesCopy[len(valuesCopy)-1]
	for _, value := range valuesCopy[:len(valuesCopy)-1] {
		if value <= pivot {
			less = append(less, value)
		}

		if value > pivot {
			greater = append(greater, value)
		}
	}

	result = append(result, quicksortLastElementPivot(less)...)
	result = append(result, pivot)
	result = append(result, quicksortLastElementPivot(greater)...)
	return result
}

func quicksortRandomElementPivot(values []int) []int {
	var less, greater, result []int
	valuesCopy := make([]int, len(values))
	copy(valuesCopy, values)

	if len(valuesCopy) < 2 {
		return valuesCopy
	}

	pivotKey := len(valuesCopy) - 1
	pivot := valuesCopy[pivotKey]
	for key, value := range valuesCopy {
		if pivotKey == key {
			continue
		}

		if value <= pivot {
			less = append(less, value)
		}

		if value > pivot {
			greater = append(greater, value)
		}
	}

	result = append(result, quicksortRandomElementPivot(less)...)
	result = append(result, pivot)
	result = append(result, quicksortRandomElementPivot(greater)...)
	return result
}

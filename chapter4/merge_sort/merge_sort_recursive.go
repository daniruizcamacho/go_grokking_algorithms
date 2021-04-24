package merge_sort

func merge(values []int, startIndex, endIndex, mediumDistance int) {
	leftSize := mediumDistance - startIndex + 1
	rightSize := endIndex - mediumDistance

	leftArray := make([]int, leftSize)
	rightArray := make([]int, rightSize)

	for i := 0; i < leftSize; i++ {
		leftArray[i] = values[startIndex+i]
	}
	for j := 0; j < rightSize; j++ {
		rightArray[j] = values[1+mediumDistance+j]
	}

	i, j, k := 0, 0, startIndex
	for i < leftSize && j < rightSize {
		if leftArray[i] <= rightArray[j] {
			values[k] = leftArray[i]
			i++
		} else {
			values[k] = rightArray[j]
			j++
		}
		k++
	}

	for i < leftSize {
		values[k] = leftArray[i]
		k++
		i++
	}

	for j < rightSize {
		values[k] = rightArray[j]
		k++
		j++
	}
}

func mergeSortRecursive(values []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	mediumDistance := (int)(startIndex + (endIndex-startIndex)/2)
	mergeSortRecursive(values, startIndex, mediumDistance)
	mergeSortRecursive(values, mediumDistance+1, endIndex)
	merge(values, startIndex, endIndex, mediumDistance)
}

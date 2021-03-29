package sorted_elements_binary_search

import "errors"

var ErrorNotFound = errors.New("Element not found.")

func binary_search(elements []int, find int) (int, error) {
	left, right, middle := 0, len(elements), 0

	for left <= right {
		middleDistance := (right - left) / 2
		middle = left + middleDistance
		if find == elements[middle] {
			return middle, nil
		}

		if find < elements[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return 0, ErrorNotFound
}

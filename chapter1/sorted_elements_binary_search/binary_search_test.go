package sorted_elements_binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := map[string] func(*testing.T) ([]int, int, int, error){
		"left position": func(t *testing.T) ([]int, int, int, error) {
			return []int{3, 10, 30, 50, 60, 100, 200, 1000},
				10,
				1,
				nil
		},
		"first position": func(t *testing.T) ([]int, int, int, error) {
			return []int{3, 10, 30, 50, 60, 100, 200, 1000},
				3,
				0,
				nil
		},
		"right position": func(t *testing.T) ([]int, int, int, error) {
			return []int{3, 10, 30, 50, 60, 100, 200, 1000},
				200,
				6,
				nil
		},
		"last position": func(t *testing.T) ([]int, int, int, error) {
			return []int{3, 10, 30, 50, 60, 100, 200, 1000},
				1000,
				7,
				nil
		},
		"middle position": func(t *testing.T) ([]int, int, int, error) {
			return []int{3, 10, 30, 50, 60, 100, 200, 1000},
				50,
				3,
				nil
		},
		"not found": func(t *testing.T) ([]int, int, int, error) {
			return []int{3, 10, 30, 50, 60, 100, 200, 1000},
				1,
				0,
				ErrorNotFound
		},
	}


	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			elements, elementToFind, expectedPosition, expectedError := tc(t)

			gotPosition, gotError := binary_search(elements, elementToFind)

			if expectedError != gotError {
				t.Fatalf("Expected error: %s; got: %s", expectedError.Error(), gotError.Error())
			}

			if gotPosition != expectedPosition {
				t.Errorf("Expected position: %d; got: %d", expectedPosition, gotPosition)
			}
		})
	}
}
package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := map[string]struct {
		arg  []int
		want []int
	}{
		"Base case empty slice": {
			arg:  []int{},
			want: []int{},
		},
		"Base case one element slice": {
			arg:  []int{1},
			want: []int{1},
		},
		"Sorted slice": {
			arg:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		"Unsorted slice": {
			arg:  []int{5, 3, 1, 2, 4},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mergeSortRecursive(tc.arg, 0, len(tc.arg)-1)

			if false == reflect.DeepEqual(tc.arg, tc.want) {
				t.Errorf("Want %v; got %v", tc.want, tc.arg)
			}
		})
	}
}

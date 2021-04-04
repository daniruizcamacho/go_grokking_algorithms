package quicksort

import (
	"reflect"
	"testing"
)

func TestQuicksortFirstElementPivot(t *testing.T) {
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
			got := quicksortFirstElementPivot(tc.arg)

			if false == reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want %v; got %v", tc.want, got)
			}
		})
	}
}

func TestQuicksortLastElementPivot(t *testing.T) {
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
			got := quicksortLastElementPivot(tc.arg)

			if false == reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want %v; got %v", tc.want, got)
			}
		})
	}
}

func TestQuicksortRandomElementPivot(t *testing.T) {
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
			got := quicksortRandomElementPivot(tc.arg)

			if false == reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want %v; got %v", tc.want, got)
			}
		})
	}
}

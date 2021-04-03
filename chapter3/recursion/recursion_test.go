package recursion

import (
	"testing"
)

func TestRecursion(t *testing.T) {
	tests := map[string]struct {
		arg  int
		want int
	}{
		"Base case 0": {
			arg:  0,
			want: 0,
		},
		"Base case 1": {
			arg:  1,
			want: 1,
		},
		"Arg 6": {
			arg:  6,
			want: 8,
		},
		"Arg 9": {
			arg:  9,
			want: 34,
		},
		"Arg 12": {
			arg:  12,
			want: 144,
		},
		"Arg 15": {
			arg:  15,
			want: 610,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fibonacci(tc.arg)

			if got != tc.want {
				t.Errorf("Want %d; got %d", tc.want, got)
			}
		})
	}
}
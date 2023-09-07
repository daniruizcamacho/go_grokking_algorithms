package knn

import "testing"

func TestKnn(t *testing.T) {
	tests := map[string]struct {
		samples []sample
		new     point
		k       int
		want    string
	}{
		"Empty samples": {
			samples: nil,
			new:     point{},
			k:       0,
			want:    "",
		},
		"Empty new sample": {
			samples: []sample{
				sample{
					class: "class1",
					point: point{
						x: 1,
						y: 1,
					},
				},
			},
			new:  point{},
			k:    1,
			want: "class1",
		},
		"Empty k": {
			samples: []sample{
				sample{
					class: "class1",
					point: point{
						x: 1,
						y: 1,
					},
				},
			},
			new:  point{},
			k:    0,
			want: "",
		},
		"Single sample": {
			samples: []sample{
				sample{
					class: "class1",
					point: point{
						x: 1,
						y: 1,
					},
				},
			},
			new: point{
				x: 1,
				y: 1,
			},
			k:    1,
			want: "class1",
		},
		"Two samples": {
			samples: []sample{
				sample{
					class: "class1",
					point: point{
						x: 1,
						y: 1,
					},
				},
				sample{
					class: "class2",
					point: point{
						x: 2,
						y: 2,
					},
				},
			},
			new: point{
				x: 1,
				y: 1,
			},
			k:    1,
			want: "class1",
		},
		"Three samples": {
			samples: []sample{
				sample{
					class: "class1",
					point: point{
						x: 1,
						y: 1,
					},
				},
				sample{
					class: "class2",
					point: point{
						x: 2,
						y: 2,
					},
				},
				sample{
					class: "class3",
					point: point{
						x: 3,
						y: 3,
					},
				},
			},
			new: point{
				x: 1,
				y: 1,
			},
			k:    1,
			want: "class1",
		},
		"Two samples, k=2": {
			samples: []sample{
				sample{
					class: "class1",
					point: point{
						x: 1,
						y: 1,
					},
				},
				sample{
					class: "class2",
					point: point{
						x: 2,
						y: 2,
					},
				},
			},
			new: point{
				x: 1,
				y: 1,
			},
			k:    2,
			want: "class1",
		},
		"Three samples, k=2": {
			samples: []sample{
				sample{
					class: "class1",
					point: point{
						x: 1,
						y: 1,
					},
				},
				sample{
					class: "class2",
					point: point{
						x: 2,
						y: 2,
					},
				},
				sample{
					class: "class3",
					point: point{
						x: 3,
						y: 3,
					},
				},
			},
			new: point{
				x: 2,
				y: 3,
			},
			k:    2,
			want: "class2",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := knn(tc.samples, tc.new, tc.k)

			if got != tc.want {
				t.Fatalf("knn(%v, %v, %v) = %v; want %v", tc.samples, tc.new, tc.k, got, tc.want)
			}
		})
	}
}

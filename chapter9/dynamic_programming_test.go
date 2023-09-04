package dynamic_programming

import (
	"reflect"
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := map[string]struct {
		items    []item
		capacity int
		want     itemsInKnapsack
	}{
		"Empty items": {
			items:    nil,
			capacity: 0,
			want:     itemsInKnapsack{},
		},
		"Empty capacity": {
			items: []item{
				item{
					name:   "item1",
					weight: 1,
					value:  1,
				},
			},
			capacity: 0,
			want:     itemsInKnapsack{},
		},
		"Single item": {
			items: []item{
				item{
					name:   "item1",
					weight: 1,
					value:  1,
				},
			},
			capacity: 1,
			want: itemsInKnapsack{
				totalValue: 1,
				items: []item{
					item{
						name:   "item1",
						weight: 1,
						value:  1,
					},
				},
			},
		},
		"Two items": {
			items: []item{
				item{
					name:   "item1",
					weight: 1,
					value:  1,
				},
				item{
					name:   "item2",
					weight: 2,
					value:  2,
				},
			},
			capacity: 3,
			want: itemsInKnapsack{
				totalValue: 3,
				items: []item{
					item{
						name:   "item1",
						weight: 1,
						value:  1,
					},
					item{
						name:   "item2",
						weight: 2,
						value:  2,
					},
				},
			},
		},
		"Three items": {
			items: []item{
				item{
					name:   "item1",
					weight: 1,
					value:  1,
				},
				item{
					name:   "item2",
					weight: 2,
					value:  2,
				},
				item{
					name:   "item3",
					weight: 3,
					value:  4,
				},
			},
			capacity: 4,
			want: itemsInKnapsack{
				totalValue: 5,
				items: []item{
					item{
						name:   "item1",
						weight: 1,
						value:  1,
					},
					item{
						name:   "item3",
						weight: 3,
						value:  4,
					},
				},
			},
		},
		"Four items": {
			items: []item{
				item{
					name:   "item1",
					weight: 1,
					value:  1,
				},
				item{
					name:   "item2",
					weight: 2,
					value:  2,
				},
				item{
					name:   "item3",
					weight: 3,
					value:  3,
				},
				item{
					name:   "item4",
					weight: 4,
					value:  4,
				},
			},
			capacity: 5,
			want: itemsInKnapsack{
				totalValue: 5,
				items: []item{
					item{
						name:   "item2",
						weight: 2,
						value:  2,
					},
					item{
						name:   "item3",
						weight: 3,
						value:  3,
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := knapsack(tc.items, tc.capacity)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("knapsack(%v, %v) = %v, want %v", tc.items, tc.capacity, got, tc.want)
			}
		})
	}
}

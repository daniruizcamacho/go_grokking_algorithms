package dijkstra

import "testing"

func TestDijkstraShortestPath(t *testing.T) {
	tests := map[string]struct {
		graph    *graph
		wantCost int
		wantPath []string
	}{
		"Empty graph": {
			graph: &graph{
				adjacencyList: nil,
			},
			wantCost: 0,
			wantPath: nil,
		},
		"Single node graph": {
			graph: &graph{
				adjacencyList: map[string]map[string]int{
					"start": nil,
				},
			},
			wantCost: 0,
			wantPath: nil,
		},
		"Two nodes graph": {
			graph: &graph{
				adjacencyList: map[string]map[string]int{
					"start": map[string]int{
						"end": 1,
					},
					"end": nil,
				},
			},
			wantCost: 1,
			wantPath: []string{"start", "end"},
		},
		"Three nodes graph": {
			graph: &graph{
				adjacencyList: map[string]map[string]int{
					"start": map[string]int{
						"a": 6,
						"b": 2,
					},
					"a": map[string]int{
						"end": 1,
					},
					"b": map[string]int{
						"a":   3,
						"end": 5,
					},
					"end": nil,
				},
			},
			wantCost: 6,
			wantPath: []string{"start", "b", "a", "end"},
		},
		"Four nodes graph": {
			graph: &graph{
				adjacencyList: map[string]map[string]int{
					"start": map[string]int{
						"a": 5,
						"b": 2,
					},
					"a": map[string]int{
						"c": 4,
						"d": 2,
					},
					"b": map[string]int{
						"a": 8,
						"d": 7,
					},
					"c": map[string]int{
						"d":   6,
						"end": 3,
					},
					"d": map[string]int{
						"end": 1,
					},
					"end": nil,
				},
			},
			wantCost: 8,
			wantPath: []string{"start", "a", "d", "end"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotCost, gotPath := tc.graph.findShortestPath("start", "end")

			if tc.wantCost != gotCost {
				t.Errorf("Want %v; got %v", tc.wantCost, gotCost)
			}

			if len(tc.wantPath) != len(gotPath) {
				t.Errorf("Want %v; got %v", tc.wantPath, gotPath)
			}

			for i, v := range tc.wantPath {
				if v != gotPath[i] {
					t.Errorf("Want %v; got %v", tc.wantPath, gotPath)
				}
			}
		})
	}
}

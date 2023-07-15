package breadth_first_search

import (
	"testing"
)

func TestFindRelationFrom(t *testing.T) {
	tests := map[string]struct {
		graph *graph
		start string
		end   string
		want  bool
	}{
		"Empty graph": {
			graph: &graph{
				adjacencyList: nil,
			},
			start: "A",
			end:   "B",
			want:  false,
		},
		"Missing start element": {
			graph: &graph{
				adjacencyList: map[string][]string{
					"B": {"C"},
					"C": {"D"},
					"D": {"E"},
					"E": {"F"},
				},
			},
			start: "A",
			end:   "F",
			want:  false,
		},
		"Missing end element": {
			graph: &graph{
				adjacencyList: map[string][]string{
					"A": {"B"},
					"B": {"C"},
					"C": {"D"},
					"D": {"E"},
				},
			},
			start: "A",
			end:   "F",
			want:  false,
		},
		"Missing connection": {
			graph: &graph{
				adjacencyList: map[string][]string{
					"A": {"B"},
					"B": {"C"},
					"C": {},
					"D": {"E"},
				},
			},
			start: "A",
			end:   "F",
			want:  false,
		},
		"Find relation": {
			graph: &graph{
				adjacencyList: map[string][]string{
					"A": {"B"},
					"B": {"C", "D"},
					"C": {"D"},
					"D": {"E", "F"},
					"E": {"F"},
				},
			},
			start: "A",
			end:   "F",
			want:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.graph.FindRelationFrom(tc.start, tc.end)

			if tc.want != got {
				t.Errorf("Want %v; got %v", tc.want, got)
			}
		})
	}
}

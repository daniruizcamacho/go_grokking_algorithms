package greedy

import (
	"testing"
)

func TestFindBestStation(t *testing.T) {
	tests := map[string]struct {
		stations *stations
		want     []string
		places   []string
	}{
		"Empty stations": {
			stations: &stations{
				stations: nil,
			},
			want:   nil,
			places: nil,
		},
		"Single station": {
			stations: &stations{
				stations: []station{
					station{
						name:   "station1",
						places: []string{"place1", "place2"},
					},
				},
			},
			want:   []string{"station1"},
			places: []string{"place1", "place2"},
		},
		"Two stations": {
			stations: &stations{
				stations: []station{
					station{
						name:   "station1",
						places: []string{"place1", "place2"},
					},
					station{
						name:   "station2",
						places: []string{"place2", "place3"},
					},
				},
			},
			want:   []string{"station1", "station2"},
			places: []string{"place1", "place2", "place3"},
		},
		"Three stations": {
			stations: &stations{
				stations: []station{
					station{
						name:   "station1",
						places: []string{"place1", "place2", "place3"},
					},
					station{
						name:   "station2",
						places: []string{"place2", "place3"},
					},
					station{
						name:   "station3",
						places: []string{"place2", "place3", "place4"},
					},
				},
			},
			want:   []string{"station1", "station3"},
			places: []string{"place1", "place2", "place3", "place4"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.stations.findBestStations(tc.places)
			if len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("got %v, want %v", got, tc.want)
				}
			}
		})
	}
}

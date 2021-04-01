package selection_sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	input := []ArtistInfo{
		{
			name: "Metallica",
			playCount: 50183,
		},
		{
			name: "Radiohead",
			playCount: 23975,
		},
		{
			name: "Guns and Roses",
			playCount: 40953,
		},
		{
			name: "U2",
			playCount: 3385,
		},
	}

	expected := []ArtistInfo{
		{
			name: "U2",
			playCount: 3385,
		},
		{
			name: "Radiohead",
			playCount: 23975,
		},
		{
			name: "Guns and Roses",
			playCount: 40953,
		},
		{
			name: "Metallica",
			playCount: 50183,
		},
	}

	got := selectionSort(input)
	if false == reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected %v; got %v", expected, got)
	}
}
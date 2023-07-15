package hash_table

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	tests := map[string]struct {
		keyValues map[int]int
		searchKey int
		wantValue int
		wantOk    bool
	}{
		"Empty hash table": {
			keyValues: nil,
			searchKey: 1,
			wantValue: 0,
			wantOk:    false,
		},
		"Insert one element": {
			keyValues: map[int]int{1: 1},
			searchKey: 1,
			wantValue: 1,
			wantOk:    true,
		},
		"Insert two elements": {
			keyValues: map[int]int{1: 1, 2: 2},
			searchKey: 2,
			wantValue: 2,
			wantOk:    true,
		},
		"Missing key": {
			keyValues: map[int]int{1: 1, 2: 2},
			searchKey: 3,
			wantValue: 0,
			wantOk:    false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ht := HashTable{}
			for key, value := range tc.keyValues {
				ht.Insert(key, value)
			}

			gotValue, gotOk := ht.Search(tc.searchKey)

			if tc.wantValue != gotValue {
				t.Errorf("Want %v; got %v", tc.wantValue, gotValue)
			}

			if tc.wantOk != gotOk {
				t.Errorf("Want %v; got %v", tc.wantOk, gotOk)
			}
		})
	}
}

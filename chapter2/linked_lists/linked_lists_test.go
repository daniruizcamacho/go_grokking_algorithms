package linked_lists

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := newList("testing list")

	expected := []string{"elementOne", "elementTwo", "elementThree", "elementFour", "elementFive"}

	for _, element := range expected {
		list.add(element)
	}

	got := list.showAll()
	if false == reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected %v; got %v", expected, got)
	}
}

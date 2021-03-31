package linked_lists

type List struct {
	name string
	head *Element
}

type Element struct {
	value string
	next  *Element
}

func newList(name string) *List {
	return &List{
		name: name,
		head: nil,
	}
}

func (list *List) showAll() []string {
	element := list.head
	var values []string
	for element != nil {
		values = append(values, element.value)
		element = element.next
	}

	return values
}

func (list *List) add(value string) {
	newNode := &Element{
		value: value,
		next:  nil,
	}

	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
	}
}

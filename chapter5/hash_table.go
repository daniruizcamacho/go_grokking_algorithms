package hash_table

type HashTable struct {
	data map[int]int
}

func (ht *HashTable) init() {
	ht.data = make(map[int]int)
}

func (ht *HashTable) Insert(key, value int) {
	if nil == ht.data {
		ht.init()
	}

	ht.data[key] = value
}

func (ht *HashTable) Search(key int) (int, bool) {
	if nil == ht.data {
		return 0, false
	}

	value, ok := ht.data[key]

	return value, ok
}

func (ht *HashTable) Delete(key int) {
	if nil == ht.data {
		return
	}

	delete(ht.data, key)
}

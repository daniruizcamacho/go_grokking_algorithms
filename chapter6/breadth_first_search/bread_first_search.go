package breadth_first_search

type graph struct {
	adjacencyList map[string][]string
}

func (g *graph) init() {
	g.adjacencyList = make(map[string][]string)
}

func (g *graph) FindRelationFrom(start, end string) bool {
	if nil == g.adjacencyList {
		return false
	}

	queue := []string{}
	queue = append(queue, start)

	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		if element == end {
			return true
		}

		for _, connection := range g.adjacencyList[element] {
			if _, ok := visited[connection]; !ok {
				visited[connection] = true
				queue = append(queue, connection)
			}
		}
	}

	return false
}

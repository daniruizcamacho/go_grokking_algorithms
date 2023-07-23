package dijkstra

import "math"

type graph struct {
	adjacencyList map[string]map[string]int
}

func (g *graph) findShortestPath(start, end string) (int, []string) {
	if nil == g.adjacencyList {
		return 0, nil
	}

	costs := make(map[string]int)
	for k, _ := range g.adjacencyList {
		costs[k] = math.MaxInt
	}
	costs[start] = 0

	parents := make(map[string]string)
	visited := make(map[string]bool)

	node := g.findLowestCostNode(costs, visited)

	for node != "" {
		cost := costs[node]
		neighbors := g.adjacencyList[node]
		for n, c := range neighbors {
			newCost := cost + c
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		visited[node] = true
		node = g.findLowestCostNode(costs, visited)
	}

	if _, ok := costs[end]; !ok {
		return 0, nil
	}

	path := []string{end}

	parent := parents[end]
	for parent != "" {
		path = append([]string{parent}, path...)
		parent = parents[parent]
	}

	return costs[end], path
}

func (g *graph) findLowestCostNode(costs map[string]int, visited map[string]bool) string {
	lowestCost := math.MaxInt
	lowestCostNode := ""
	for node, cost := range costs {
		if _, ok := visited[node]; !ok && cost < lowestCost {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

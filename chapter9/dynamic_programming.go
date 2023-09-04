package dynamic_programming

type item struct {
	name   string
	weight int
	value  int
}

type itemsInKnapsack struct {
	totalValue int
	items      []item
}

func knapsack(items []item, capacity int) itemsInKnapsack {
	if len(items) == 0 || capacity == 0 {
		return itemsInKnapsack{}
	}

	minWeight := getMinWeight(items)
	if capacity < minWeight {
		return itemsInKnapsack{}
	}

	gcdValue := greatestCommonDivisor(items)
	matrix := make(map[string]map[int]itemsInKnapsack)
	for i := 0; i < len(items); i++ {
		matrix[items[i].name] = make(map[int]itemsInKnapsack)
		for j := minWeight; j <= capacity; j += gcdValue {
			matrix[items[i].name][j] = itemsInKnapsack{}
		}
	}

	for i := 0; i < len(items); i++ {
		for j := minWeight; j <= capacity; j += gcdValue {
			if i == 0 {
				if items[i].weight <= j {
					matrix[items[i].name][j] = itemsInKnapsack{
						totalValue: items[i].value,
						items:      []item{items[i]},
					}
				}
				continue
			}

			if j < items[i].weight {
				matrix[items[i].name][j] = matrix[items[i-1].name][j]
				continue
			}

			if j >= items[i].weight {
				previousItems := matrix[items[i-1].name][j-items[i].weight]
				previousItems.items = append(previousItems.items, items[i])
				previousItems.totalValue += items[i].value

				if previousItems.totalValue > matrix[items[i-1].name][j].totalValue {
					matrix[items[i].name][j] = previousItems
				} else {
					matrix[items[i].name][j] = matrix[items[i-1].name][j]
				}
			}
		}
	}

	return matrix[items[len(items)-1].name][capacity]
}

func getMinWeight(items []item) int {
	min := items[0].weight
	for _, item := range items {
		if item.weight < min {
			min = item.weight
		}
	}
	return min
}

func greatestCommonDivisor(items []item) int {
	if len(items) == 0 {
		return 0
	}

	if len(items) == 1 {
		return items[0].weight
	}

	gcdValue := items[0].weight
	for i := 1; i < len(items); i++ {
		gcdValue = gcd(gcdValue, items[i].weight)
	}

	return gcdValue
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	if a%b == 0 {
		return b
	}

	return gcd(b, a%b)
}

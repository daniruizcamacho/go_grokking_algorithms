package knn

import (
	"math"
	"sort"
)

type point struct {
	x int
	y int
}

type sample struct {
	class string
	point point
}

type distance struct {
	sample   sample
	distance int
}

func knn(samples []sample, newSample point, k int) string {
	if len(samples) == 0 {
		return ""
	}

	if len(samples) < k {
		return ""
	}

	if k < 1 {
		return ""
	}

	distances := []distance{}

	for _, v := range samples {
		calculatedDistance := calculateDistance(v.point, newSample)
		distances = append(distances, distance{sample: v, distance: calculatedDistance})
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	classes := make(map[string]int)

	for _, v := range distances[:k] {
		classes[v.sample.class]++
	}

	max := 0
	class := ""

	for k, v := range classes {
		if v > max || class == "" {
			max = v
			class = k
		}
	}

	return class
}

func calculateDistance(sample1, sample2 point) int {
	return int(math.Sqrt(float64((sample1.x-sample2.x)*(sample1.x-sample2.x) + (sample1.y-sample2.y)*(sample1.y-sample2.y))))
}

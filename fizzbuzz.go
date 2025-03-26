package main

import "fmt"

type cost struct {
	day int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	result := []float64{}

	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(result) {
				result = append(result, 0.0)
		}
		result[cost.day] += cost.value
	}

	return result
}

func main() {
	costs := []cost{
		{ value: 0.56, day: 0 },
		{ value: 0.45, day: 3 },
		{ value: 0.72, day: 4 },
		{ value: 0.21, day: 6 },
		{ value: 0.09, day: 7 },
	}

	fmt.Println(getCostByDay(costs))
}
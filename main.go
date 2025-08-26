package main

import (
	"fmt"
	"math"
	"sort"
)

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		var rangeKey int
		if temp < 0 {
			rangeKey = int(math.Ceil(temp/10)) * 10
		} else {
			rangeKey = int(math.Floor(temp/10)) * 10
		}

		groups[rangeKey] = append(groups[rangeKey], temp)
	}

	return groups
}

func printGroups(groups map[int][]float64) {
	var keys []int
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println("Группировка:")
	for _, key := range keys {
		temps := groups[key]
		fmt.Printf("%d: {", key)
		for i, temp := range temps {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%.1f", temp)
		}
		fmt.Println("}")
	}
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemperatures(temperatures)

	printGroups(groups)
}

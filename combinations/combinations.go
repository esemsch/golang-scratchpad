package main

import (
	"fmt"
	"math"
)

func main() {
	vals := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	numOfCombs := int64(math.Pow(float64(2), float64(len(vals))))

	var combinations [][]string
	var i int64 = 0
	for ; i < numOfCombs; i++ {
		var combination []string
		for j := 0; j < len(vals); j++ {
			mask := int64(math.Pow(2, float64(j)))
			if i&mask > 0 {
				combination = append(combination, vals[j])
			}
		}
		combinations = append(combinations, combination)
	}

	for _, c := range combinations {
		fmt.Println(c)
	}
	fmt.Println("Total of", len(combinations), "combinations")
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Quicksort")

	//vals := []int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1, 2, 3, 4, 626, 626, 4, 3, 3, 34, 4, 4, 45, 35, 25, 45, 4, 54, 52, 5, 5, 25, 25, 25}
	vals := []int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}

	quicksort(vals)

	fmt.Println(vals)
}

func quicksort(toSort []int) {
	//fmt.Println(toSort)
	if len(toSort) == 1 {
		return
	} else {
		pivot := toSort[len(toSort)/2]
		rightStop := len(toSort) - 1
		//fmt.Printf("pivot = %d\n", pivot)
		for i := 0; i <= rightStop; i++ {
			//fmt.Printf("left: %d[%d]\n", toSort[i], i)
			if toSort[i] >= pivot {
				for ; rightStop > i; rightStop-- {
					//fmt.Printf("right: %d[%d]\n", toSort[rightStop], rightStop)
					if toSort[rightStop] <= pivot {
						//fmt.Printf("%d[%d] <-> %d[%d]\n", toSort[i], i, toSort[rightStop], rightStop)
						aux := toSort[rightStop]
						toSort[rightStop] = toSort[i]
						toSort[i] = aux
						break
					}
				}
			}
		}
		//fmt.Println(toSort)
		//fmt.Printf("rightStop = %d[%d]\n", toSort[rightStop], rightStop)

		quicksort(toSort[:rightStop])
		quicksort(toSort[rightStop:])
	}
}

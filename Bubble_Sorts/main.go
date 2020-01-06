// https://www.codeabbey.com/index/task_view/bubble-sort

package main

import (
	"fmt"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanMassive() (massive_ []int64) {
	count := ScanInt64()
	massive_ = make([]int64, 0, count)
	for i := int64(0); i < count; i++ {
		nextValue := ScanInt64()
		massive_ = append(massive_, nextValue)
	}
	return
}

func BubbleSort(unsortedMassive []int64) (passes_, swaps_ int64) {
	for {
		saveSwaps := swaps_
		for i, stop := 0, len(unsortedMassive)-2; i <= stop; i++ {
			if unsortedMassive[i] > unsortedMassive[i+1] {
				unsortedMassive[i], unsortedMassive[i+1] = unsortedMassive[i+1], unsortedMassive[i]
				swaps_++
			}
		}
		passes_++
		if saveSwaps == swaps_ {
			break
		}
	}
	return
}

func main() {
	passes, swaps := BubbleSort(ScanMassive())
	fmt.Printf("%v %v", passes, swaps)
}

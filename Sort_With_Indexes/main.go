// https://www.codeabbey.com/index/task_view/sort-with-indexes

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

func IndexesSort(unsortedMassive []int64) (sortedIndexes_ []int64) {
	sortedIndexes_ = make([]int64, len(unsortedMassive))
	for i := range sortedIndexes_ {
		sortedIndexes_[i] = int64(i + 1)
	}
	for {
		swaps := int64(0)
		for i, stop := 0, len(unsortedMassive)-2; i <= stop; i++ {
			if unsortedMassive[i] > unsortedMassive[i+1] {
				unsortedMassive[i], unsortedMassive[i+1] = unsortedMassive[i+1], unsortedMassive[i]
				sortedIndexes_[i], sortedIndexes_[i+1] = sortedIndexes_[i+1], sortedIndexes_[i]
				swaps++
			}
		}
		if swaps == 0 {
			break
		}
	}
	return
}

func main() {
	for _, value := range IndexesSort(ScanMassive()) {
		fmt.Printf("%v ", value)
	}
}

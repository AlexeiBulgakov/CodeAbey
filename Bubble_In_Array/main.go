// https://www.codeabbey.com/index/task_view/bubble-in-array--ru

package main

import (
	"fmt"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanMassive() (massive_ []int64) {
	for {
		nextValue := ScanInt64()
		if nextValue != -1 {
			massive_ = append(massive_, nextValue)
		} else {
			break
		}
	}
	return
}

func ComputeChecksum(array []int64) (checksum_ int64) {
	for i, _ := range array {
		checksum_ += array[i]
		checksum_ *= 113
		checksum_ %= 10000007
	}
	return
}

func Bubbling(unsortedMassive []int64) (swaps_ int64) {
	for i, stop := 0, len(unsortedMassive)-2; i <= stop; i++ {
		if unsortedMassive[i] > unsortedMassive[i+1] {
			unsortedMassive[i], unsortedMassive[i+1] = unsortedMassive[i+1], unsortedMassive[i]
			swaps_++
		}
	}
	return
}

func main() {
	massive := ScanMassive()
	swaps := Bubbling(massive)
	checksum := ComputeChecksum(massive)
	fmt.Printf("%v %v", swaps, checksum)
}

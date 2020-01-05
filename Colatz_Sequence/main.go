// https://www.codeabbey.com/index/task_view/collatz-sequence

package main

import (
	"fmt"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		counter, startNumber := int64(0), ScanInt64()
		for {
			if startNumber == 1 {
				fmt.Printf("%v ", counter)
				break
			} else if startNumber%2 == 0 {
				startNumber /= 2
				counter++
			} else {
				startNumber *= 3
				startNumber++
				counter++
			}
		}
	}
}

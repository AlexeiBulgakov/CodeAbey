// https://www.codeabbey.com/index/task_view/arithmetic-progression

package main

import (
	"fmt"
)

func main() {
	var count1 uint32
	fmt.Scan(&count1)
	for i := uint32(0); i < count1; i++ {
		var first, step, count2, result int32
		fmt.Scan(&first)
		fmt.Scan(&step)
		fmt.Scan(&count2)
		for j := int32(0); j < count2; j++ {
			result += (first + (step * j))
		}
		fmt.Printf("%v ", result)
	}
}

// https://www.codeabbey.com/index/task_view/median-of-three

package main

import (
	"fmt"
	"sort"
)

func main() {
	var count int32
	fmt.Scan(&count)
	for i := int32(0); i < count; i++ {
		digits := make([]int, 0, 3)
		for j := 0; j < 3; j++ {
			var value int
			fmt.Scan(&value)
			digits = append(digits, value)
		}
		sort.Ints(digits)
		fmt.Printf("%v ", digits[1])
	}
}

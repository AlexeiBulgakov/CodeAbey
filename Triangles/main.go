// https://www.codeabbey.com/index/task_view/triangles

package main

import (
	"fmt"
	"sort"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanTriangleSide() (a_, b_, c_ int) {
	fmt.Scanf("%d %d %d\n", &a_, &b_, &c_)
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		a, b, c := ScanTriangleSide()
		triangle := []int{a, b, c}
		sort.Ints(triangle)
		if (triangle[0] + triangle[1]) > triangle[2] {
			fmt.Print("1 ")
		} else {
			fmt.Print("0 ")
		}
	}
}

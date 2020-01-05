// https://www.codeabbey.com/index/task_view/linear-function

package main

import (
	"fmt"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanParams() (x1, y1, x2, y2 int64) {
	x1 = ScanInt64()
	y1 = ScanInt64()
	x2 = ScanInt64()
	y2 = ScanInt64()
	return
}

func ComputeAB(x1, y1, x2, y2 int64) (a_, b_ float64) {
	a_ = float64(y2-y1) / float64(x2-x1)
	b_ = float64(y1) - a_*float64(x1)
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		a, b := ComputeAB(ScanParams())
		fmt.Printf("(%v %v) ", a, b)
	}
}

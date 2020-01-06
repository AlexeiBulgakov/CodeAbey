// https://www.codeabbey.com/index/task_view/square-root

package main

import (
	"fmt"
	"math/big"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ComputeSqurt(digit, presicion int64) (squrt_ float64) {
	squrt_ = 1
	for i := int64(0); i < presicion; i++ {
		squrt_ = (squrt_ + float64(digit)/squrt_) / 2
	}
	return
}

func main() {
	for i, coutn := int64(0), ScanInt64(); i < coutn; i++ {
		digit := ScanInt64()
		presicion := ScanInt64()
		fmt.Printf("%v ", big.NewFloat(ComputeSqurt(digit, presicion)).Text('f', 7))
	}
}

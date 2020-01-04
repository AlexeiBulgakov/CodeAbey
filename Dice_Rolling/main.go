// https://www.codeabbey.com/index/task_view/dice-rolling

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanFloan64() (f_ float64) {
	fmt.Scanf("%f", &f_)
	return
}

func DiceRolling(target float64) (result_ int64) {
	result_ = int64(target*float64(6)) + 1
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		fmt.Printf("%v ", DiceRolling(ScanFloan64()))
	}
}

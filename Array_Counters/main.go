// https://www.codeabbey.com/index/task_view/array-counters

package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader = bufio.NewReader(os.Stdin)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func main() {
	count := ScanInt64()
	max := ScanInt64()
	massive := make([]int64, max+1)
	for i := int64(0); i < count; i++ {
		element := ScanInt64()
		massive[element]++
	}
	for _, element := range massive[1:] {
		fmt.Printf("%v ", element)
	}
}

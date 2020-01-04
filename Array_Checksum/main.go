// https://www.codeabbey.com/index/task_view/array-checksum

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputReader = bufio.NewReader(os.Stdin)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanArrayOfInt64() (array_ []int64) {
	arrayInStr, _ := inputReader.ReadString('\n')
	for _, digitInStr := range strings.Fields(arrayInStr) {
		digit, _ := strconv.ParseInt(digitInStr, 10, 64)
		array_ = append(array_, digit)
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

func main() {
	_ = ScanInt64()
	fmt.Printf("%v ", ScanArrayOfInt64())
}

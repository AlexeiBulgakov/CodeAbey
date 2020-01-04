// https://www.codeabbey.com/index/task_view/average-of-array

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanArrayOfInt64() (array_ []int64) {
	arrayInStr, _ := reader.ReadString('\n')
	for _, digitInStr := range strings.Fields(arrayInStr) {
		digit, _ := strconv.ParseInt(digitInStr, 10, 64)
		array_ = append(array_, digit)
	}
	// последний элемент массива, получаемого из потока ввода, не является целевым элементом для дальнейших вычислений:
	// его значение всегда '0', а смысл - конец строки
	array_ = array_[:len(array_)-1]
	return
}

func AverageOfArray(array []int64) (average_ int64) {
	for i, _ := range array {
		average_ += array[i]
	}
	if div, mod := math.Modf(float64(average_) / float64(len(array))); mod >= 0.5 {
		average_ = int64(div) + 1
	} else {
		average_ = int64(div)
	}
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		fmt.Printf("%v ", AverageOfArray(ScanArrayOfInt64()))
	}
}

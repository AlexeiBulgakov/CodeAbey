// https://www.codeabbey.com/index/task_view/sum-in-loop

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count int64
var values []int64

func init() {
	// чтение количества значений
	reader := bufio.NewReader(os.Stdin)
	countInStr, _ := reader.ReadString('\n')
	if valueInInt, err := strconv.ParseInt(countInStr, 0, 64); err == nil {
		count = valueInInt
	} /*else {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}*/
	// чтение массива со значениями
	valuesInStr, _ := reader.ReadString('\n')
	for _, valueInStr := range strings.Fields(valuesInStr) {
		if valueInInt, err := strconv.ParseInt(valueInStr, 0, 64); err == nil {
			values = append(values, valueInInt)
		} /*else {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}*/
	}
}

func main() {
	// TODO: для устранения ошибок переполнения использовать пакет math/big
	var sum int64
	for _, value := range values {
		sum += value
	}
	fmt.Println(sum)
}

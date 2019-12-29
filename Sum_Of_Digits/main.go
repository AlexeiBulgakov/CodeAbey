// https://www.codeabbey.com/index/task_view/sum-of-digits

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var countOfValues uint32
	fmt.Scan(&countOfValues)
	for i := uint32(0); i < countOfValues; i++ {
		var A, B, C uint64
		fmt.Scan(&A)
		fmt.Scan(&B)
		fmt.Scan(&C)
		res := A*B + C
		// Замечание: для суммирования цифр числа в 8-ми ричной СЧ, необходимо оперировать байтами числа; в 16-ти ричное - срезом по 2-байта
		var sumOfDigits uint64
		for _, sign := range strconv.FormatUint(res, 10) {
			oneSign, _ := strconv.ParseUint(string(sign), 0, 64)
			sumOfDigits += oneSign
		}
		fmt.Printf("%v ", sumOfDigits)
	}
}

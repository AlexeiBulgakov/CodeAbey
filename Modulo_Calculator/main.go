// https://www.codeabbey.com/index/task_view/modular-calculator

package main

import (
	"fmt"
	"math/big"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanNextOperands() (operator_ string, values_ int64) {
	fmt.Scanf("%s %d\n", &operator_, &values_)
	return
}

func main() {
	var firstValue = big.NewInt(ScanInt64())
	for {
		operator, secondValue := ScanNextOperands()
		if operator == "+" {
			firstValue = firstValue.Add(firstValue, big.NewInt(secondValue))
		} else if operator == "*" {
			firstValue = firstValue.Mul(firstValue, big.NewInt(secondValue))
		} else if operator == "%" {
			firstValue = firstValue.Mod(firstValue, big.NewInt(secondValue))
			break
		}
	}
	fmt.Printf("%v ", firstValue)
}

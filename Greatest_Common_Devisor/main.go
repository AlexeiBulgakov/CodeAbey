// https://www.codeabbey.com/index/task_view/bubble-sort

package main

import (
	"fmt"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

// НОД
func GCD(a, b int64) (gcd_ int64) {
	if a == b {
		gcd_ = a
	} else {
		for {
			if a > b {
				a = a % b
				if a == 0 {
					gcd_ = b
					break
				}
			} else {
				b = b % a
				if b == 0 {
					gcd_ = a
					break
				}
			}
		}
	}
	return
}

// НОК
func LCM(a, b int64) (lcm_ int64) {
	lcm_ = a * b / GCD(a, b)
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		a := ScanInt64()
		b := ScanInt64()
		fmt.Printf("(%v %v) ", GCD(a, b), LCM(a, b))
	}
}

// https://www.codeabbey.com/index/task_view/weighted-sum-of-digits

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanWord() (word_ string) {
	fmt.Scanf("%s", &word_)
	word_ = strings.TrimSpace(word_)
	return
}

func WeightingDigit(digitInStr string) (weightedSum_ int64) {
	for i := 0; i < len(digitInStr); i++ {
		loneDigit, _ := strconv.ParseInt(string(digitInStr[i]), 10, 64)
		weightedSum_ += loneDigit * int64((i + 1))
	}
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		fmt.Printf("%v ", WeightingDigit(ScanWord()))
	}
}

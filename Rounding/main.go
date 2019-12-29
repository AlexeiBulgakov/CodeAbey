// https://www.codeabbey.com/index/task_view/rounding

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var countOfValues uint32
	fmt.Scan(&countOfValues)
	stdinStreamReader := bufio.NewReader(os.Stdin)
	for i := uint32(0); i < countOfValues; i++ {
		twoValuesFromImput, _ := stdinStreamReader.ReadString('\n')
		twoValuesFromImput = strings.TrimSuffix(twoValuesFromImput, "\n")
		valuesInStr := strings.Fields(twoValuesFromImput)
		value1, _ := strconv.ParseInt(valuesInStr[0], 0, 64)
		value2, _ := strconv.ParseInt(valuesInStr[1], 0, 64)
		div, mod := math.Modf(float64(value1) / float64(value2))
		if math.Abs(mod) >= float64(0.5) {
			if div >= 0 {
				div++
			} else {
				div--
			}
		}
		if div == 0 {
			//  т.к. '0', '-0' и '+0' - это три разных вывода!
			div = 0
		}
		fmt.Printf("%v ", div)
	}
}

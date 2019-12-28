// https://www.codeabbey.com/index/task_view/min-of-two

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var count uint32
	fmt.Scan(&count)
	reader := bufio.NewReader(os.Stdin)
	for i := uint32(0); i < count; i++ {
		valuesInStr, _ := reader.ReadString('\n')
		valuesInStr = strings.TrimSuffix(valuesInStr, "\n")
		valuesInMas := strings.Fields(valuesInStr)
		value1, _ := strconv.ParseInt(valuesInMas[0], 0, 64)
		value2, _ := strconv.ParseInt(valuesInMas[1], 0, 64)
		if value1 < value2 {
			fmt.Printf("%v ", value1)
		} else {
			fmt.Printf("%v ", value2)
		}
	}
}

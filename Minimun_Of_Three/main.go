// https://www.codeabbey.com/index/task_view/min-of-three

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	toIntMas := func(massiveInStr []string) (massiveInInt_ []int) {
		for _, valueInStr := range massiveInStr {
			valueInInt64, _ := strconv.ParseInt(valueInStr, 0, 64)
			massiveInInt_ = append(massiveInInt_, int(valueInInt64))
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)
	valuesInStr, _ := reader.ReadString('\n')
	valuesInStr = strings.TrimSuffix(valuesInStr, "\n")
	values := toIntMas(strings.Fields(valuesInStr))
	sort.Ints(values)
	fmt.Printf("%v %v", values[0], values[len(values)-1])
}

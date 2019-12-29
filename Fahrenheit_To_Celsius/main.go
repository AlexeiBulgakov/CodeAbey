// https://www.codeabbey.com/index/task_view/fahrenheit-celsius

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
	stdinStreamReader := bufio.NewReader(os.Stdin)
	twoValuesFromImput, _ := stdinStreamReader.ReadString('\n')
	twoValuesFromImput = strings.TrimSuffix(twoValuesFromImput, "\n")
	valuesInStr := strings.Fields(twoValuesFromImput)
	for _, valueInStr := range valuesInStr[1:] {
		fahrengeit, _ := strconv.ParseInt(valueInStr, 0, 64)
		celsius, mod := math.Modf(float64(fahrengeit-32) / 1.8)
		if mod > 0.5 {
			celsius++
		}
		fmt.Printf("%v ", celsius)
	}
}

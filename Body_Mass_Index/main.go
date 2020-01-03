// https://www.codeabbey.com/index/task_view/body-mass-index

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

func ScanLine() (line_ string) {
	line_, _ = reader.ReadString('\n')
	line_ = strings.TrimSuffix(line_, "\n")
	return
}

func ScanInt64() (i_ int64) {
	i_, _ = strconv.ParseInt(ScanLine(), 10, 64)
	return
}

func ScanTargetValues() (height_, weight_ float64) {
	lines := strings.Fields(ScanLine())
	height_, _ = strconv.ParseFloat(lines[0], 64)
	weight_, _ = strconv.ParseFloat(lines[1], 64)
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		weight, height := ScanTargetValues()
		if IBM := float64(weight) / math.Pow(height, 2); IBM < 18.5 {
			fmt.Print("under ")
		} else if IBM < 25.0 {
			fmt.Print("normal ")
		} else if IBM < 30.0 {
			fmt.Print("over ")
		} else {
			fmt.Print("obese ")
		}
	}
}

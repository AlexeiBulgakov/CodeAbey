// https://www.codeabbey.com/index/task_view/sums-in-loop

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	l uint64
	r uint64
}

func (_pair Pair) Summ() (summ_ uint64) {
	return _pair.l + _pair.r
}

var pairs []Pair

func init() {
	reader := bufio.NewReader(os.Stdin)
	// получение количесва пар
	coutPairsInStr, _ := reader.ReadString('\n')
	coutPairsInInt, _ := strconv.ParseInt(strings.TrimSuffix(coutPairsInStr, "\n"), 0, 64)
	// получение пар
	parseOnePair := func() (pair_ Pair) {
		pairInStr, _ := reader.ReadString('\n')
		pairInMas := strings.Fields(pairInStr)
		l, _ := strconv.ParseInt(pairInMas[0], 0, 64)
		r, _ := strconv.ParseInt(pairInMas[1], 0, 64)
		pair_ = Pair{uint64(l), uint64(r)}
		return
	}
	for i := int64(0); i < coutPairsInInt; i++ {
		pairs = append(pairs, parseOnePair())
	}
}

func main() {
	for _, pair := range pairs {
		fmt.Printf("%v ", pair.Summ())
	}
}

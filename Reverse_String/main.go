// https://www.codeabbey.com/index/task_view/reverse-string

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader = bufio.NewReader(os.Stdin)

func ScanStr() (str_ string) {
	str_, _ = inputReader.ReadString('\n')
	str_ = strings.TrimSuffix(str_, "\n")
	return
}

func ReverseString(targetStr string) (reverseStr_ string) {
	targetRunes, reverseRunes := []rune(targetStr), make([]rune, 0, len(targetStr))
	for i := len(targetRunes) - 1; i >= 0; i-- {
		reverseRunes = append(reverseRunes, targetRunes[i])
	}
	reverseStr_ = string(reverseRunes)
	return
}

func main() {
	fmt.Printf("%v ", ReverseString(ScanStr()))
}

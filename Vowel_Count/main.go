// https://www.codeabbey.com/index/task_view/vowel-count

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var countOfValues uint32
	fmt.Scan(&countOfValues)
	stdinStreamReader := bufio.NewReader(os.Stdin)
	for i := uint32(0); i < countOfValues; i++ {
		targetStr, _ := stdinStreamReader.ReadString('\n')
		fmt.Printf("%v ",
			strings.Count(targetStr, "a")+
				strings.Count(targetStr, "o")+
				strings.Count(targetStr, "u")+
				strings.Count(targetStr, "i")+
				strings.Count(targetStr, "e")+
				strings.Count(targetStr, "y"))
	}
}

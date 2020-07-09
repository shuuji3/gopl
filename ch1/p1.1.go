// 問題1.1 コマンド名も表示しなさい
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("command:", os.Args[0])
	fmt.Println("args:", strings.Join(os.Args[1:], " "))
}

// 問題1.2 引数のインデックスと値のペアをプリントしなさい
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, "-", arg)
	}
}

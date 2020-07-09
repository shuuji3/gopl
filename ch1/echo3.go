// 効率的な strings.Join() を利用してプログラムの引数をプリントする
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
}

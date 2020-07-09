// 2種類のechoの実行時間を測定しなさい
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

// 非効率バージョンを10億回実行する
func inefficient() {
	for i := 0; i < int(math.Pow10(9)); i++ {
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "

		}
		//fmt.Println(s)
	}

}

// 効率バージョンを10億回実行する
func efficient() {
	for i := 0; i < int(math.Pow10(9)); i++ {
		strings.Join(os.Args[1:], " ")
		//fmt.Println(strings.Join(os.Args[1:], " "))
	}
}

func main() {
	start := time.Now()
	inefficient()
	t1 := time.Since(start)
	fmt.Println("inefficient:", t1.Seconds())

	start = time.Now()
	efficient()
	t2 := time.Since(start)
	fmt.Println("efficient:", t2.Seconds())
}

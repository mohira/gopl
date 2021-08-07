package main

import (
	"fmt"
	"github.com/mohira/gopl/ch2/2.7/7/foo"
)

func main() {
	fmt.Println("main")

	fmt.Println(foo.Cwd == "") // true; つまり foo.Cwdが文字列のゼロ値のまま！
}

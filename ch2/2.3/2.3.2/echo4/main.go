package main

import (
	"flag"
	"fmt"
	"strings"
)

// それぞれポインタ型を返す
var n *bool = flag.Bool("n", false, "omit trailing newline(改行の省略)")
var sep *string = flag.String("s", " ", "separator")

// p.36
func main() {
	flag.Parse()

	fmt.Printf("%T\n", n)   // *bool
	fmt.Printf("%T\n", sep) // *string

	fmt.Println(flag.Args())
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}

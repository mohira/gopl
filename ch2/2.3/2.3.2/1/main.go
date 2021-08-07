package main

import "fmt"

// p.35
func main() {
	x := 1
	p := &x
	fmt.Println(*p) // 1

	*p = 2
	fmt.Println(x) // 2
}

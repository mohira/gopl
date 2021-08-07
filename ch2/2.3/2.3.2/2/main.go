package main

import "fmt"

// p.35 中段
func main() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
}

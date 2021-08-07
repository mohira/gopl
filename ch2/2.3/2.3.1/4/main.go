package main

import "fmt"

func main() {
	// ./main.go:7:6: a redeclared in this block
	//        previous declaration at ./main.go:6:6
	var a = 1
	var a = 2

	fmt.Println(a)
}

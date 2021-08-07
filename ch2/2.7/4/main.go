package main

import "fmt"

func f() int {
	return 0
}

func g(n int) int {
	return 0
}

// p.53
func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}

	fmt.Printf(x, y) // undefined x, y
}

package main

import (
	"fmt"
)

func main() {
	p := new(int)
	fmt.Printf("%T\n", p) // *int
	fmt.Println(p)        // 0xc0000140c8
	fmt.Println(*p)       // 0

	*p = 2
	fmt.Println(*p) // 2
}

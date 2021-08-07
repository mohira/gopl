package main

import "fmt"

// アドレスが同じになる例
func main() {
	fmt.Printf("%p\n", make([]int, 0)) // 0x57b418
	fmt.Printf("%p\n", make([]int, 0)) // 0x57b418
	fmt.Printf("%p\n", make([]int, 0)) // 0x57b418

	fmt.Printf("%p\n", new([0]int)) // 0x57b418
	fmt.Printf("%p\n", new([0]int)) // 0x57b418
	fmt.Printf("%p\n", new([0]int)) // 0x57b418
}

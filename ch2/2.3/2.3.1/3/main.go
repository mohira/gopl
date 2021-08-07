package main

import "fmt"

// ただ多値を返すことを表現するだけの関数
func tuple() (int, int) {
	return 1, 2
}
func main() {
	a, x := tuple()
	a, y := tuple()

	fmt.Println(a, x, y)
}

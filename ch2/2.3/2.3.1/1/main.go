package main

import "os"

// p.34中段: 宣言 と 代入 は違う！
func main() {
	f, err := os.Open("foo")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("bar") // no new variables on left side of :=

	if err != nil {
		panic(err)
	}
}

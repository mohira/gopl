package main

import "fmt"

// p.49
// c, a, b の順で評価される == 依存関係をうまいこと解決している
var a = valueA()
var b = valueB()
var c = 1

func valueA() int {
	fmt.Println("call valueA()")
	return b + c
}

func valueB() int {
	fmt.Println("call valueB()")
	return c + 1
}

func main() {
	fmt.Println("main start")
	fmt.Println("main done")
}

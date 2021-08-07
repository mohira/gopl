package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

var p = f()

// p.35 下段
func main() {
	fmt.Println(p == p)     // true
	fmt.Println(*p == *f()) // true; 値は同じ
	fmt.Println(p == f())   // false; アドレスは違う
	fmt.Println(f() == f()) // false
}

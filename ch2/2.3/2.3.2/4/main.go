package main

import "fmt"

func incr(p *int) int {
	*p++ // 実体に加算
	return *p
}

func main() {
	v := 1
	fmt.Println(v) // 1

	incr(&v)
	fmt.Println(v) // 2

	fmt.Println(incr(&v)) // 3
}

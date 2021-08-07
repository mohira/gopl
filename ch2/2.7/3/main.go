package main

import "fmt"

// p.52 {}がレキシカルブロックってわけじゃない！
// for文の暗黙のレキシカルブロックを作っている
func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("outer: i=%d", i)
		i := i * i
		fmt.Printf(" inner: i=%d\n", i)
	}
}

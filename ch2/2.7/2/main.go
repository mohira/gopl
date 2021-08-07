package main

import "fmt"

// p.52: 3つのレキシカルブロック
// Goroutineの i:=i も同じっぽい
// 他の言語とくらべると混乱するやつ
func main() {
	x := "hello!"

	for i := 0; i < len(x); i++ {
		fmt.Printf("i=%d\n", i)
		fmt.Printf("\tfor block: x=%v\n", x) // ループごとに x を「宣言」している！
		x := x[i]                            // 2回めのループでのxはどうなるんだ？
		if x != '!' {
			x := x + 'A' - 'a' // 英小文字 を 英大文字 に変換
			fmt.Printf("\t if block: x=%c\n", x)
		}
	}
}

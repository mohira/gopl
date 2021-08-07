package main

import "fmt"

// シャドーイング
var x = 1

// [Big Sky :: Go 言語で変数のシャドウイングを避けたいなら shadow を使おう。](https://mattn.kaoriya.net/software/lang/go/20200227102218.htm)
func main() {
	x := 2
	fmt.Println(x)
}

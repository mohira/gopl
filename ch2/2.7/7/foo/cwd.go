package foo

import (
	"log"
	"os"
)

var Cwd string // 実は使われてない！！！ GoLandは警告出してくれるね

func init() {
	Cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v\n", err)
	}
	// ここの Cwd は init()内のレキシカルスコープで「宣言」された Cwd
	// なので、パッケージレベルの変数cwdはゼロ値のまま！
	// なので、他のファイルから Cwd を 利用しようとすると...
	log.Printf("Working directory = %s\n", Cwd)
}

package main

import (
	"log"
	"os"
)

func main() {

	if f, err := os.Open("tmp"); err != nil {
		log.Println(err)
	} else {
		// else句もブロックに含まれるのでセーフ
		f.Stat()
	}

}

package main

import (
	"fmt"
	"test/wc"
)

func main() {
	//s := "go is awesome, php is not"
	s := "in a coat of gold or a coat of red"
	w := wc.MakeWords(s)

	fmt.Println(w.Index("a"))
	// 0
	fmt.Println(w.Index("is"))
	// 1
	fmt.Println(w.Index("is not"))
	// -1
	fmt.Println(w.Index("python"))
	// -1
}

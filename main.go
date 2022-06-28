package main

import (
	"Golang/hangman"
	"fmt"
)

func main() {
	g := hangman.New(8, "GOLANG")
	fmt.Println(g)
}

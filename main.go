package main

import (
	"Golang/hangman"
	"fmt"
)

func main() {
	g := hangman.New(8, "GOLANG")
	fmt.Println(g)
	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(l)
}

package main

import (
	"Golang/hangman"
	"fmt"
	"os"
)

func main() {
	g := hangman.New(8, "GOLANG")
	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Println("Error: %v", err)
			os.Exit(1)
		}
		guess = l
		fmt.Println(l)
	}

}

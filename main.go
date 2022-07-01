package main

import (
	"fmt"
	"hangman/dictionary"
	"hangman/hangman"
	"os"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		panic(err)
	}

	g, err := hangman.New(8, dictionary.PickWord())
	if err != nil {
		panic(err)
	}
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
		g.MakeAGuess(guess)
	}

}

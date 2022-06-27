package hangman

import "strings"

type Game struct {
	State       string
	Letters     []string
	FndLetters  []string
	UsedLetters []string
	TurnsLeft   int
}

func New(turns int, words string) *Game {
	letters := strings.Split(strings.ToUpper(words), "")
	found := make([]string, len(letters))

	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State:       "",
		Letters:     letters,
		FndLetters:  found,
		UsedLetters: []string{},
		TurnsLeft:   turns,
	}

	return g
}

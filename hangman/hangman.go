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

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FndLetters) {
			g.State = "won"
		}
	}
}

func letterInWord(guess string, letter []string) bool {
	for _, l := range letter {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FndLetters[i] = guess
		}
	}
}

func hasWon(letters []string, fondLetters []string) bool {
	for i := range letters {
		if letters[i] != fondLetters[i] {
			return false
		}
	}
	return true
}

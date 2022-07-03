package hangman

import (
	"testing"
)

func TestNewGameWork(t *testing.T) {
	g, _ := New(8, "golang")
	letter := []string{"G", "O", "L", "A", "N", "G"}
	for i, v := range g.Letters {
		if v != letter[i] {
			t.Errorf("Error: array word")
		}
	}
}

func TestNewGameEmptyWord(t *testing.T) {
	_, err := New(8, "")
	if err == nil {
		t.Errorf("Invalid word to guess")
	}
}

func TestLetterNotInWord(t *testing.T) {
	letter := []string{"A", "D", "N"}
	guess := "B"
	b := letterInWord(guess, letter)
	if b {
		t.Errorf("Word not in letter")
	}
}

func TestLetterInWord(t *testing.T) {
	letter := []string{"A", "D", "N"}
	guess := "A"
	b := letterInWord(guess, letter)
	if !b {
		t.Errorf("Word in letter")
	}
}

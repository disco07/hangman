package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
 __ __   ____  ____    ____  ___ ___   ____  ____  
|  |  | /    ||    \  /    ||   |   | /    ||    \ 
|  |  ||  o  ||  _  ||   __|| _   _ ||  o  ||  _  |
|  _  ||     ||  |  ||  |  ||  \_/  ||     ||  |  |
|  |  ||  _  ||  |  ||  |_ ||   |   ||  _  ||  |  |
|  |  ||  |  ||  |  ||     ||   |   ||  |  ||  |  |
|__|__||__|__||__|__||___,_||___|___||__|__||__|__|
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FndLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word", guess)
	case "lost":
		fmt.Print("You lost :( ! Teh word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You won! Teh word was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
			
		`
	}
	println(draw)
}

func main() {

}

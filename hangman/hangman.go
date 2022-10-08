package hangman

import "strings"

const ALREADY_GUESSED = "alreadyGuessed"
const GOOD_GUESS = "goodGuess"
const BAD_GUESS = "badGuess"
const LOST = "lost"
const WON = "won"

type Game struct {
	State        string   //Game state
	Letters      []string //Letters in the word to find
	FoundLetters []string
	UsedLetters  []string
	TurnsLeft    int
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	foundLetters := make([]string, len(word))
	for i := range letters {
		foundLetters[i] = "_"
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: foundLetters,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if letterInWorld(guess, g.UsedLetters) {
		g.State = ALREADY_GUESSED
	} else if letterInWorld(guess, g.Letters) {
		g.State = GOOD_GUESS
		g.revealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = WON
		}
	} else {
		g.UsedLetters = append(g.UsedLetters, guess)
		g.TurnsLeft--
		if g.TurnsLeft > 0 {
			g.State = BAD_GUESS
		} else {
			g.State = LOST
		}

	}
}

func (g *Game) revealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, v := range g.Letters {
		if guess == v {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func letterInWorld(guess string, letters []string) bool {
	for _, v := range letters {
		if guess == v {
			return true
		}
	}
	return false
}

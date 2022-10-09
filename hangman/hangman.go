package hangman

import (
	"fmt"
	"strings"
)

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

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("word '%s' must be at least 3 characters, got=%v", word, len(word))
	}
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
	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	switch g.State {
	case WON, LOST:
		return
	}
	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = ALREADY_GUESSED
	} else if letterInWord(guess, g.Letters) {
		g.State = GOOD_GUESS
		g.revealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = WON
		}
	} else {
		g.State = BAD_GUESS

		g.UsedLetters = append(g.UsedLetters, guess)
		g.TurnsLeft--

		if g.TurnsLeft <= 0 {
			g.State = "lost"
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

func letterInWord(guess string, letters []string) bool {
	for _, v := range letters {
		if guess == v {
			return true
		}
	}
	return false
}

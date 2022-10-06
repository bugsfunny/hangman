package hangman

import "strings"

type Game struct {
	State        string   //Game state
	Letters      []string //Letters in the word to find
	FoundLetters []string
	UsedLetters  []string
	TurnLeft     int
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
		TurnLeft:     turns,
	}
	return g
}

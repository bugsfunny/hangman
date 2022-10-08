package main

import (
	"fmt"
	"os"

	"github.com/bugsfunny/hangman/hangman"
	"github.com/bugsfunny/hangman/hangman/dictionary"
)

func main() {
	if err := dictionary.Load("words.txt"); err != nil {
		fmt.Printf("Could not load word from file: %v", err)
		os.Exit(1)
	}
	g := hangman.New(8, dictionary.PickWord())
	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case hangman.WON, hangman.LOST:
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}

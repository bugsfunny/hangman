package main

import (
	"fmt"
	"os"

	"github.com/bugsfunny/hangman/hangman"
)

func main() {
	game := hangman.New(8, "qwerty")
	fmt.Println(game)
	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Could not read from terminal: %v", err)
		os.Exit(1)
	}
	fmt.Println(l)
}

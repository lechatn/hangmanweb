package main

import ("../hangman"
		"fmt"
	"c:\Users\najma\hangman")

func main(){ 
	clear := "\033[H\033[2J"
	fmt.Print(clear)
	hangman.Menu()
}

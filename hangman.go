package main

<<<<<<< HEAD
import (
	"fmt"

	"../hangman"
)

func main() {
=======
import ("../hangman"
		"fmt")

func main(){ 
>>>>>>> main
	clear := "\033[H\033[2J"
	fmt.Print(clear)
	hangman.Menu()
}

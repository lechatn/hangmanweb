package hangman

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// La structure HangmanGame représente l'état du jeu du pendu.
type HangmanGame struct {
	WordToGuess      string
	GuessesRemaining int
	IncorrectGuesses []string
}

// Fonction pour obtenir un mot aléatoire du package "hangman".
func getRandomWord() string {
	words := []string{"golang", "programming", "developer", "hangman", "web"}
	rand.Seed(time.Now().Unix())
	return words[rand.Intn(len(words))]
}

func HangmanHandler(w http.ResponseWriter, r *http.Request) {
	// Générez un mot aléatoire.
	wordToGuess := getRandomWord()
	// Analysez le modèle HTML.
	t, err := template.New("hangman").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Les données à passer au modèle HTML.
	game := HangmanGame{
		WordToGuess:      wordToGuess,
		GuessesRemaining: 6, // Par exemple, 6 tentatives avant de perdre.
		IncorrectGuesses: []string{},
	}

	// Générez la page HTML en utilisant le modèle et les données.
	if err := t.Execute(w, game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Associez la fonction HangmanHandler à un chemin de votre serveur.
	http.HandleFunc("/hangman", HangmanHandler)

	// Démarrez le serveur HTTP sur le port 8080.
	http.ListenAndServe(":8080", nil)
}

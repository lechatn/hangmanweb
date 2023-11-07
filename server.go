package main

import (
	"fmt"
	"github.com/GuillaumeAntier/hangman"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/game", game)
	http.HandleFunc("/letter", handleLetter) // Définissez une route distincte pour la gestion de la lettre
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tIndex, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	tIndex.Execute(w, nil)
}

func game(w http.ResponseWriter, r *http.Request) {
	tGame, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	Dyna2 := hangman.DisplayWord(hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt")))
	// Générez le contenu HTML avec la variable dynamique
	htmlContent := fmt.Sprintf("%s", Dyna2)
	data := struct{ Res string }{htmlContent}
	// Écrivez la réponse HTML dans la sortie HTTP
	tGame.Execute(w, data)

}

func handleLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		lettre := r.PostFormValue("letterInput") // Utilisez "letterInput" au lieu de "test"
		// Insérez ici le code de traitement de la lettre.
		fmt.Fprintf(w, "La lettre saisie est : %s", lettre)
	} else {
		// Générez le formulaire ou redirigez si nécessaire.
	}
}

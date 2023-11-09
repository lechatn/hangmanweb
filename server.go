package main

import (
	"fmt"
	"github.com/lechatn/hangman"
	"net/http"
	"text/template"
)

func main() {
	word := hangman.DisplayWord(hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt")))
	life := 10
	http.HandleFunc("/", index)
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		game(w, r, word)
	})
	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		letter(w, r, word, &life)
	})
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

func game(w http.ResponseWriter, r *http.Request, word string) {
	tGame, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	Dyna2 := word
	// Générez le contenu HTML avec la variable dynamique
	htmlContent := fmt.Sprintf("%s", Dyna2)
	life := 10
	htmlContent2 :=fmt.Sprintf("%d", life)
	data := struct{
		Res  string
		Life string
	}{
		Res:  htmlContent,
		Life: htmlContent2,
	}
	// Écrivez la réponse HTML dans la sortie HTTP
	tGame.Execute(w, data)

}

func letter(w http.ResponseWriter, r *http.Request, word string, life *int) {
	tletter, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	Dyna2 := word
	// Générez le contenu HTML avec la variable dynamique
	htmlContent := fmt.Sprintf("%s", Dyna2)
	*life = *life - 1
	htmlContent2 :=fmt.Sprintf("%d", *life)
	data := struct{
		Res  string
		Life string
	}{
		Res:  htmlContent,
		Life: htmlContent2,
	}
	tletter.Execute(w, data)
	// Écrivez la réponse HTML dans la sortie HTTP
    //lettre := r.PostFormValue("letterInput")

}

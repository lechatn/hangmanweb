package main

import (
	"fmt"
	"github.com/lechatn/hangman"
	"net/http"
	"text/template"
)

func main() {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt"))
	display := hangman.DisplayWord(word) 
	indexHangman := 0
	failed_letter := ""
	life := 10
	http.HandleFunc("/", index)
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		game(w, r, display,life)
	})
	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		display,life,indexHangman,failed_letter = letter(w, r, word, life,display,indexHangman,failed_letter)
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

func game(w http.ResponseWriter, r *http.Request, display string, life int) {
	tGame, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	Dyna := display
	// Générez le contenu HTML avec la variable dynamique
	htmlContent := fmt.Sprintf("%s", Dyna)
	htmlContent2 :=fmt.Sprintf("%d", life)
	data := struct{
		Display  string
		Life string
	}{
		Display:  htmlContent,
		Life: htmlContent2,
	}
	// Écrivez la réponse HTML dans la sortie HTTP
	tGame.Execute(w, data)

}

func letter(w http.ResponseWriter, r *http.Request, word string, life int, display string, indexHangman int, failed_letter string) (string,int,int,string) {
	tletter, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	// Générez le contenu HTML avec la variable dynamique
	
	// Écrivez la réponse HTML dans la sortie HTTP
    letter := r.PostFormValue("letterInput")
	display,life, indexHangman, failed_letter = hangman.IsPresent(letter,word,display,life,indexHangman,failed_letter)
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 :=fmt.Sprintf("%d", life)
	htmlContent3 :=fmt.Sprintf("%s", display)
	htmlContent4 :=fmt.Sprintf("%d", indexHangman)
	htmlContent5 :=fmt.Sprintf("%s", failed_letter)
	data := struct{
		Display  string
		Life string
		Word string
		indexHangman string
		failed_letter string
	}{	
		Display:  htmlContent,
		Life: htmlContent2,
		Word: htmlContent3,
		indexHangman: htmlContent4,
		failed_letter: htmlContent5,
	}
	tletter.Execute(w, data)

	return display,life,indexHangman,failed_letter
}

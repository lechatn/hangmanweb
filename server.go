package main

import (
	"fmt"
	"github.com/lechatn/hangman"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt"))
	print(word, "\n")
	Display := hangman.DisplayWord(word)
	IndexHangman := 0
	Failed_letter := ""
	life := 10
	http.HandleFunc("/", index)
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		game(w, r, Display, life)
	})
	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		Display, life, IndexHangman, Failed_letter = letter(w, r, word, life, Display, IndexHangman, Failed_letter)
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

func game(w http.ResponseWriter, r *http.Request, Display string, life int) {
	tGame, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	Dyna := Display
	// Générez le contenu HTML avec la variable dynamique
	htmlContent := fmt.Sprintf("%s", Dyna)
	htmlContent2 := fmt.Sprintf("%d", life)
	data := struct {
		Display string
		Life    string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
	}
	// Écrivez la réponse HTML dans la sortie HTTP
	tGame.Execute(w, data)

}

func letter(w http.ResponseWriter, r *http.Request, word string, life int, Display string, IndexHangman int, Failed_letter string) (string, int, int, string) {
	tletter, err := template.ParseFiles("letter.html")
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	// Générez le contenu HTML avec la variable dynamique

	// Écrivez la réponse HTML dans la sortie HTTP
	letter := r.PostFormValue("letterInput")
	file, err := ioutil.ReadFile("affichage/hangman.txt")
	Display, life, IndexHangman, Failed_letter = hangman.IsPresent(strings.ToUpper(letter), word, Display, life, IndexHangman, Failed_letter)
	htmlContent := fmt.Sprintf("%s", Display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", word)
	htmlContent4 := ""
	htmlContent5 := fmt.Sprintf("%s", Failed_letter)
	lines := strings.Split(string(file), "\n") // We cut the file hangman.txt line by line
	if life == 10 {
		for i := 0; i < 7; i++ {
			htmlContent4 = htmlContent4 + lines[i] + "<br>"
		}
	} else {
		for i := IndexHangman; i < IndexHangman+7; i++ {
			htmlContent4 = htmlContent4 + lines[i] + "<br>"
		}
	}
	if word == Display {
		win(w, r)
		return Display, life, IndexHangman, Failed_letter
	}
	data := struct {
		Display       string
		Life          string
		Word          string
		IndexHangman  string
		Failed_letter string
	}{
		Display:       htmlContent,
		Life:          htmlContent2,
		Word:          htmlContent3,
		IndexHangman:  htmlContent4,
		Failed_letter: htmlContent5,
	}
	tletter.Execute(w, data)

	return Display, life, IndexHangman, Failed_letter
}

func win(w http.ResponseWriter, r *http.Request) {
	tWin, err := template.ParseFiles("win.html")
	if err != nil {
		panic(err)
	}
	tWin.Execute(w, nil)
}

package main

import (
	"fmt"
	"github.com/lechatn/hangman"
	"net/http"
	"strings"
	"text/template"
	"strconv"
)

func main() {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt"))
	print(word, "\n")
	Display := hangman.DisplayWord(word)
	Failed_letter := ""
	life := 10
	http.HandleFunc("/", index)
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		game(w, r, Display, life)
	})
	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		Display, life, Failed_letter,word = restart(w, r, word, life, Display, Failed_letter)
	})
	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		Display, life, Failed_letter = letter(w, r, word, life, Display, Failed_letter)
	})
	http.HandleFunc("/regle", func(w http.ResponseWriter, r *http.Request) {
		regle(w, r)
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

func letter(w http.ResponseWriter, r *http.Request, word string, life int, Display string, Failed_letter string) (string, int, string) {
	tletter, err := template.ParseFiles("letter.html")
	IndexHangman := 0
	if err != nil {
		panic(err)
	}
	// Créez une variable dynamique en Go
	// Générez le contenu HTML avec la variable dynamique

	// Écrivez la réponse HTML dans la sortie HTTP
	letter := r.PostFormValue("letterInput")
	Display, life,IndexHangman, Failed_letter = hangman.IsPresent(strings.ToUpper(letter), word, Display, life, IndexHangman, Failed_letter)
	htmlContent := fmt.Sprintf("%s", Display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", word)
	htmlContent4 := ""
	htmlContent5 := fmt.Sprintf("%s", Failed_letter)
	htmlContent6 := "static/images/hangman-"+strconv.Itoa(9-life)+".png"
	print(htmlContent6)
	if word == Display {
		win(w, r)
		return Display, life, Failed_letter
	}
	if life == 0 {	
		loose(w, r, word)
		return Display, life, Failed_letter
	}
	data := struct {
		Display       string
		Life          string
		Word          string
		IndexHangman  string
		Failed_letter string
		ImageName     string
	}{
		Display:       htmlContent,
		Life:          htmlContent2,
		Word:          htmlContent3,
		IndexHangman:  htmlContent4,
		Failed_letter: htmlContent5,
		ImageName:     htmlContent6,
	}
	tletter.Execute(w, data)

	return Display, life, Failed_letter
}

func win(w http.ResponseWriter, r *http.Request) {
	tWin, err := template.ParseFiles("win.html")
	if err != nil {
		panic(err)
	}
	tWin.Execute(w, nil)
}

func regle(w http.ResponseWriter, r *http.Request) {	
	tRegles, err := template.ParseFiles("regle.html")
	if err != nil {
		panic(err)
	}
	tRegles.Execute(w, nil)
}

func restart(w http.ResponseWriter, r *http.Request, word string, life int, Display string, Failed_letter string) (string, int, string, string){
	tRestart, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	word = hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt"))
	Display = hangman.DisplayWord(word)
	Failed_letter = ""
	life = 10
	htmlContent := fmt.Sprintf("%s", Display)
	htmlContent2 := fmt.Sprintf("%d", life)
	data := struct {	
		Display string
		Life    string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
	
	}
	tRestart.Execute(w, data)

	return Display,life, Failed_letter,word
}

func loose(w http.ResponseWriter, r *http.Request, word string) {
	tWin, err := template.ParseFiles("loose.html")
	if err != nil {
		panic(err)
	}
	tWin.Execute(w, nil)
}
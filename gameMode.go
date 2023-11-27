package hangmanweb

import (
	"fmt"
	"github.com/lechatn/hangman"
	"net/http"
	"text/template"
)

func Gamemode(w http.ResponseWriter, r *http.Request) {
	tGamemode, err := template.ParseFiles("template/game_mode.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)
}

func Gamemode_french(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_french"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "French")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_english(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/english.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_english"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "English")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_german(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/german.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_german"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "German")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_drinks(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/drinks.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_drinks"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Drinks")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_capitals(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/capitals.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_capitals"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Capitals")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_spanish(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/spanish.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_spanish"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Spanish")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_food(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/food.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_food"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Food")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_italiano(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/italiano.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_italiano"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Italiano")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_brands(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/brands.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_brands"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Brands")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_countrys(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/countrys.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_country"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Countrys")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_portugese(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/portugese.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_portugese"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Portugese")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_sports(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/sports.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_sports"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "Sports")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}

func Gamemode_french_citys(w http.ResponseWriter, r *http.Request, score int) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/french_citys.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_citys"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent3 := fmt.Sprintf("%s", "French Citys")
	htmlContent, htmlContent2, htmlContent4 := PrintHtml(display, life, score)
	data := CreateData(htmlContent, htmlContent2, htmlContent3, htmlContent4)
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode, htmlContent3
}
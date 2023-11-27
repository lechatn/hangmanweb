package main

import (
	"fmt"
	"github.com/lechatn/hangman"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	//Define the different values
	var word string
	var life int
	var display string
	var Failed_letter string
	var game_mode string
	var name_mode string

	//Define the different routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Failed_letter = index(w, r)
	})
	http.HandleFunc("/gamemode", func(w http.ResponseWriter, r *http.Request) {
		gamemode(w, r)
	})
	http.HandleFunc("/gamemode_french", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_french(w, r)
	})
	http.HandleFunc("/gamemode_english", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_english(w, r)
	})
	http.HandleFunc("/gamemode_german", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_german(w, r)
	})
	http.HandleFunc("/gamemode_drinks", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_drinks(w, r)
	})
	http.HandleFunc("/gamemode_capitals", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_capitals(w, r)
	})
	http.HandleFunc("/gamemode_spanish", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_spanish(w, r)
	})
	http.HandleFunc("/gamemode_food", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_food(w, r)
	})
	http.HandleFunc("/gamemode_italiano", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_italiano(w, r)
	})
	http.HandleFunc("/gamemode_brands", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_brands(w, r)
	})
	http.HandleFunc("/gamemode_countrys", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_countrys(w, r)
	})
	http.HandleFunc("/gamemode_portugese", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_portugese(w, r)
	})
	http.HandleFunc("/gamemode_sports", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_sports(w, r)
	})
	http.HandleFunc("/gamemode_french_citys", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, game_mode,name_mode = gamemode_french_citys(w, r)
	})
	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		display, life, Failed_letter = letter(w, r, word, life, display, Failed_letter, game_mode, name_mode)
	})
	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		Failed_letter = restart(w, r)
	})
	http.HandleFunc("/regle", func(w http.ResponseWriter, r *http.Request) {
		regle(w, r)
	})
	http.HandleFunc("/contact", contact) 
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

//Functions

func index(w http.ResponseWriter, r *http.Request) string{
	Failed_letter := ""
	tIndex, err := template.ParseFiles("template/index.html")
	if err != nil {
		panic(err)
	}
	tIndex.Execute(w, nil)
	return Failed_letter
}

func letter(w http.ResponseWriter, r *http.Request, word string, life int, Display string, Failed_letter string, game_mode string,name_mode string) (string, int, string) {
	tletter, err := template.ParseFiles("template/letter.html")
	IndexHangman := 0
	if err != nil {
		panic(err)
	}
	letter := r.URL.Query().Get("id")
	Display, life, IndexHangman, Failed_letter = hangman.IsPresent(strings.ToUpper(letter), word, Display, life, IndexHangman, Failed_letter)
	htmlContent := fmt.Sprintf("%s", Display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", word)
	htmlContent4 := ""
	htmlContent5 := fmt.Sprintf("%s", Failed_letter)
	htmlContent6 := "static/images/hangman-" + strconv.Itoa(10-life) + ".png"
	htmlContent7 := fmt.Sprintf("%s", name_mode)
	fmt.Println(Failed_letter)
	if word == Display {
		win(w, r,word)
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
		Game_mode     string
	}{
		Display:       htmlContent,
		Life:          htmlContent2,
		Word:          htmlContent3,
		IndexHangman:  htmlContent4,
		Failed_letter: htmlContent5,
		ImageName:     htmlContent6,
		Game_mode:     htmlContent7,
	}
	tletter.Execute(w, data)
	return Display, life, Failed_letter
}

func win(w http.ResponseWriter, r *http.Request, word string) {
	tWin, err := template.ParseFiles("template/win.html")
	if err != nil {
		panic(err)
	}
	htmlContent := fmt.Sprintf("%s", word)
	data := struct {
		Word string
	}{
		Word: htmlContent,
	}
	tWin.Execute(w, data)
}

func regle(w http.ResponseWriter, r *http.Request) {
	tRegles, err := template.ParseFiles("template/regle.html")
	if err != nil {
		panic(err)
	}
	tRegles.Execute(w, nil)
}

func restart(w http.ResponseWriter, r *http.Request) string {
	tRestart, err := template.ParseFiles("template/game_mode.html")
	if err != nil {
		panic(err)
	}
	tRestart.Execute(w, nil)
	Failed_letter := ""
	return Failed_letter
}

func loose(w http.ResponseWriter, r *http.Request, word string) {
	tWin, err := template.ParseFiles("template/loose.html")
	if err != nil {
		panic(err)
	}
	htmlContent := fmt.Sprintf("%s", word)
	data := struct {
		Word string
	}{
		Word: htmlContent,
	}

	tWin.Execute(w, data)
}

func gamemode(w http.ResponseWriter, r *http.Request) {
	tGamemode, err := template.ParseFiles("template/game_mode.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)
}

func gamemode_french(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_french"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "French")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_english(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/english.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_english"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "English")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_german(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/german.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_german"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "German")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_drinks(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/drinks.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_drinks"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Drinks")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_capitals(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/capitals.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_capitals"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Capitals")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_spanish(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/spanish.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_spanish"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Spanish")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_food(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/food.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_food"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Food")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_italiano(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/italiano.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_italiano"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Italiano")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_brands(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/brands.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_brands"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Brands")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_countrys(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/countrys.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_country"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Countrys")
	data := struct {
		Display string
		Life    string
		Game_mode string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_portugese(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/portugese.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_portugese"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3 := fmt.Sprintf("%s", "Portugese")
	data := struct {
		Display string
		Life    string
		Game_mode  string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_sports(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/sports.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_sports"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3  := fmt.Sprintf("%s", "Sports")
	data := struct {
		Display string
		Life    string
		Game_mode  string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func gamemode_french_citys(w http.ResponseWriter, r *http.Request) (string, int, string, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/french_citys.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_citys"
	tGamemode, err := template.ParseFiles("template/game.html")
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent3  := fmt.Sprintf("%s", "French Citys")
	data := struct {
		Display string
		Life    string
		Game_mode  string
	}{
		Display: htmlContent,
		Life:    htmlContent2,
		Game_mode: htmlContent3,
	}
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, data)
	return display, life, word, game_mode,htmlContent3
}

func contact(w http.ResponseWriter, r *http.Request) {
    tcontact, err := template.ParseFiles("template/contact.html")
    if err != nil {
        panic(err)
    }
    tcontact.Execute(w, nil)
}


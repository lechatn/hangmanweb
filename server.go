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
	var word string
	var life int
	var display string
	var Failed_letter string
	var game_mode string
	http.HandleFunc("/", index)
	http.HandleFunc("/gamemode", func(w http.ResponseWriter, r *http.Request) {
		gamemode(w, r)
	})
	http.HandleFunc("/gamemode_french", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_french(w,r)
	})

	http.HandleFunc("/gamemode_english", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_english(w,r)
	})

	http.HandleFunc("/gamemode_german", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_german(w,r)
	})

	http.HandleFunc("/gamemode_drinks", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_drinks(w,r)
	})

	http.HandleFunc("/gamemode_capitals", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_capitals(w,r)
	})

	http.HandleFunc("/gamemode_spanish", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_spanish(w,r)
	})

	http.HandleFunc("/gamemode_food", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_food(w,r)
	})

	http.HandleFunc("/gamemode_italiano", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_italiano(w,r)
	})

	http.HandleFunc("/gamemode_brands", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_brands(w,r)
	})

	http.HandleFunc("/gamemode_countrys", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_countrys(w,r)
	})

	http.HandleFunc("/gamemode_portugese", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_portugese(w,r)
	})

	http.HandleFunc("/gamemode_sports", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_sports(w,r)
	})

	http.HandleFunc("/gamemode_french_citys", func(w http.ResponseWriter, r *http.Request) {
		display,life,word,game_mode = gamemode_french_citys(w,r)
	})

	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		display,life,Failed_letter = letter(w,r,word,life,display,Failed_letter,game_mode)
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

func letter(w http.ResponseWriter, r *http.Request, word string, life int, Display string, Failed_letter string,game_mode string) (string, int, string) {
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
	htmlContent6 := "static/images/hangman-"+strconv.Itoa(10-life)+".png"
	htmlContent7 := fmt.Sprintf("%s", game_mode)
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

func gamemode(w http.ResponseWriter, r *http.Request) {
	tGamemode, err := template.ParseFiles("game_mode.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)
}

func gamemode_french(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/words.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_french"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode
}

func gamemode_english(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/english.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_english"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode
}

func gamemode_german(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/german.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_german"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode
}

func gamemode_drinks(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/drinks.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_drinks"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_capitals(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/capitals.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_capitals"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_spanish(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/spanish.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_spanish"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_food(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/food.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_food"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_italiano(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/italiano.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_italiano"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}

	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_brands(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/brands.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_brands"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	
	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_countrys(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/countrys.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_country"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}

	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_portugese(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/portugese.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_portugese"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}

	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_sports(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/sports.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_sports"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}

	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}

func gamemode_french_citys(w http.ResponseWriter, r *http.Request) (string, int, string, string) {
	word := hangman.RandomWord(hangman.LoadWords("base_de_donnée/french_citys.txt"))
	display := hangman.DisplayWord(word)
	life := 10
	game_mode := "/gamemode_citys"
	tGamemode, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}

	tGamemode.Execute(w, nil)

	return display, life, word, game_mode

}



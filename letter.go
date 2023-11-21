package hangmanweb

import ( "net/http"
		 "text/template"
		 "fmt"
		 "github.com/lechatn/hangman"
		 "strconv"
		 "strings"
		 )

func Letter(w http.ResponseWriter, r *http.Request, word string, life int, Display string, Failed_letter string, game_mode string, name_mode string, score int, win_series int) (string, int, string, int, int) {
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
	htmlContent8 := fmt.Sprintf("%d", score)
	fmt.Println(Failed_letter)
	if word == Display {
		score, win_series = Win(w, r, word, Failed_letter, life, score, win_series)
		return Display, life, Failed_letter, score, win_series
	}
	if life == 0 {
		score = Loose(w, r, word, Failed_letter, life, score)
		win_series = 0
		return Display, life, Failed_letter, score, win_series
	}
	data := struct {
		Display       string
		Life          string
		Word          string
		IndexHangman  string
		Failed_letter string
		ImageName     string
		Game_mode     string
		Score         string
	}{
		Display:       htmlContent,
		Life:          htmlContent2,
		Word:          htmlContent3,
		IndexHangman:  htmlContent4,
		Failed_letter: htmlContent5,
		ImageName:     htmlContent6,
		Game_mode:     htmlContent7,
		Score:         htmlContent8,
	}
	tletter.Execute(w, data)
	return Display, life, Failed_letter, score, win_series
}
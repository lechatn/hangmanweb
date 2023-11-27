package hangmanweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func Win(w http.ResponseWriter, r *http.Request, word string, Failed_letter string, life int, score int, win_series int) (int, int) {
	tWin, err := template.ParseFiles("template/win.html")
	if err != nil {
		panic(err)
	}
	win_series++
	score = score + life/2*win_series
	htmlContent := fmt.Sprintf("%s", word)
	htmlContent2 := fmt.Sprintf("%d", score)
	data := struct {
		Word  string
		Score string
	}{
		Word:  htmlContent,
		Score: htmlContent2,
	}
	tWin.Execute(w, data)
	return score, win_series
}

func Loose(w http.ResponseWriter, r *http.Request, word string, Failed_letter string, life int, score int) int {
	tWin, err := template.ParseFiles("template/loose.html")
	if err != nil {
		panic(err)
	}
	htmlContent := fmt.Sprintf("%s", word)
	htmlContent2 := fmt.Sprintf("%d", score)
	data := struct {
		Word  string
		Score string
	}{
		Word:  htmlContent,
		Score: htmlContent2,
	}
	tWin.Execute(w, data)
	return score
}

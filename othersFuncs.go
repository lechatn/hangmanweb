package hangmanweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) (string, int, int) {
	score := 0
	win_series := 0
	Failed_letter := ""
	tIndex, err := template.ParseFiles("template/index.html")
	if err != nil {
		panic(err)
	}
	tIndex.Execute(w, nil)
	return Failed_letter, score, win_series
}

func Regle(w http.ResponseWriter, r *http.Request) {
	tRegles, err := template.ParseFiles("template/regle.html")
	if err != nil {
		panic(err)
	}
	tRegles.Execute(w, nil)
}

func Restart(w http.ResponseWriter, r *http.Request) string {
	tRestart, err := template.ParseFiles("template/game_mode.html")
	if err != nil {
		panic(err)
	}
	tRestart.Execute(w, nil)
	Failed_letter := ""
	return Failed_letter
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tcontact, err := template.ParseFiles("template/contact.html")
	if err != nil {
		panic(err)
	}
	tcontact.Execute(w, nil)
}

func PrintHtml(display string, life int, score int) (string, string, string) {
	htmlContent := fmt.Sprintf("%s", display)
	htmlContent2 := fmt.Sprintf("%d", life)
	htmlContent4 := fmt.Sprintf("%d", score)
	return htmlContent, htmlContent2, htmlContent4
}

func CreateData(htmlContent string, htmlContent2 string, htmlContent3 string, htmlContent4 string) struct {
	Display   string
	Life      string
	Game_mode string
	Score     string
} {
	data := struct {
		Display   string
		Life      string
		Game_mode string
		Score     string
	}{
		Display:   htmlContent,
		Life:      htmlContent2,
		Game_mode: htmlContent3,
		Score:     htmlContent4,
	}
	return data
}
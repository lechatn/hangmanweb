package hangmanweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func Win(w http.ResponseWriter, r *http.Request, word string, Failed_letter string, life int, score int, win_series int) (int, int) {
	tWin, err := template.ParseFiles("template/win.html") // Parse the html file
	if err != nil {
		panic(err)
	}
	win_series++ // Increment the win series
	score = score + life/2*win_series // Calculate the score with this formula : score = score + life/2*win_series
	htmlContent := fmt.Sprintf("%s", word) // Define the htmlContent for the display in the html file
	htmlContent2 := fmt.Sprintf("%d", score)
	data := struct { // Create the data for the html file
		Word  string
		Score string
	}{
		Word:  htmlContent,
		Score: htmlContent2,
	}
	tWin.Execute(w, data) // Execute the html file
	return score, win_series
}

func Loose(w http.ResponseWriter, r *http.Request, word string, Failed_letter string, life int, score int) int {
	tWin, err := template.ParseFiles("template/loose.html") // Parse the html file
	if err != nil {
		panic(err)
	}
	htmlContent := fmt.Sprintf("%s", word) // Define the htmlContent for the display in the html file
	htmlContent2 := fmt.Sprintf("%d", score)
	data := struct { // Create the data for the html file
		Word  string
		Score string
	}{
		Word:  htmlContent,
		Score: htmlContent2,
	}
	tWin.Execute(w, data) // Execute the html file
	return score
}

package main

import (
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/game", game)
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

func game(w http.ResponseWriter, r *http.Request) {
	tGame, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}
	tGame.Execute(w, nil)
}


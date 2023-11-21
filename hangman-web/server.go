package main

import (
	"net/http"
	"hangmanweb"
)

func main() {
	//Define the different values
	var word string
	var life int
	var display string
	var Failed_letter string
	var name_mode string
	var score int
	var win_series int
	var Game_mode string

	//Define the different routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Failed_letter, score, win_series = hangmanweb.Index(w, r)
	})
	http.HandleFunc("/gamemode", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Gamemode(w, r)
	})
	http.HandleFunc("/gamemode_french", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_french(w, r, score)
	})
	http.HandleFunc("/gamemode_english", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_english(w, r, score)
	})
	http.HandleFunc("/gamemode_german", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_german(w, r, score)
	})
	http.HandleFunc("/gamemode_drinks", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_drinks(w, r, score)
	})
	http.HandleFunc("/gamemode_capitals", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_capitals(w, r, score)
	})
	http.HandleFunc("/gamemode_spanish", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_spanish(w, r, score)
	})
	http.HandleFunc("/gamemode_food", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_food(w, r, score)
	})
	http.HandleFunc("/gamemode_italiano", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_italiano(w, r, score)
	})
	http.HandleFunc("/gamemode_brands", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_brands(w, r, score)
	})
	http.HandleFunc("/gamemode_countrys", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_countrys(w, r, score)
	})
	http.HandleFunc("/gamemode_portugese", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_portugese(w, r, score)
	})
	http.HandleFunc("/gamemode_sports", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_sports(w, r, score)
	})
	http.HandleFunc("/gamemode_french_citys", func(w http.ResponseWriter, r *http.Request) {
		display, life, word, Game_mode, name_mode = hangmanweb.Gamemode_french_citys(w, r, score)
	})
	http.HandleFunc("/letter", func(w http.ResponseWriter, r *http.Request) {
		display, life, Failed_letter, score, win_series = hangmanweb.Letter(w, r, word, life, display, Failed_letter, Game_mode, name_mode, score, win_series)
	})
	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		Failed_letter = hangmanweb.Restart(w, r)
	})
	http.HandleFunc("/regle", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Regle(w, r)
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		hangmanweb.Contact(w, r)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}


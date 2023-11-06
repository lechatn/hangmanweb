package main

import (
    "net/http"
	"text/template"
)

func main() {

    http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
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

func test(w http.ResponseWriter, r *http.Request) {
	tTest, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}
	tTest.Execute(w, nil)
}
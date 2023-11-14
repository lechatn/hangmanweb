package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type PageData struct {
	Life int
}

var wordsToGuess = []string
var wordToGuess string
var guessedWord []string
var incorrectGuesses []string
var maxIncorrectGuesses = 10

func init() {
	rand.Seed(time.Now().UnixNano())
	wordToGuess = getRandomWord()
	guessedWord = make([]string, len(wordToGuess))

	for i := 0; i < len(guessedWord); i++ {
		guessedWord[i] = "_"
	}
}

func getRandomWord() string {
	return wordsToGuess[rand.Intn(len(wordsToGuess))]
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmpl, err := template.New("index").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Life: maxIncorrectGuesses - len(incorrectGuesses),
	}
	renderTemplate(w, yourHTMLString, data)
}

func letterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		letter := strings.ToUpper(r.FormValue("letterInput"))

		// Vérifier si la lettre est correcte
		if strings.Contains(wordToGuess, letter) {
			// Mettre à jour le mot deviné
			for i, char := range wordToGuess {
				if string(char) == letter {
					guessedWord[i] = letter
				}
			}
		} else {
			// Ajouter la lettre incorrecte à la liste
			incorrectGuesses = append(incorrectGuesses, letter)
		}

		// Vérifier si le joueur a gagné
		if strings.Join(guessedWord, "") == wordToGuess {
			// Afficher le message de victoire
			renderTemplate(w, yourHTMLString, PageData{Life: 0}) // 0 pour indiquer que le joueur a gagné
			return
		}

		// Vérifier si le joueur a perdu
		if len(incorrectGuesses) >= maxIncorrectGuesses {
			// Afficher le message de défaite
			renderTemplate(w, yourHTMLString, PageData{Life: 0}) // 0 pour indiquer que le joueur a perdu
			return
		}

		// Rediriger vers la page du jeu avec les mises à jour
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/letter", letterHandler)
	http.ListenAndServe(":8080", nil)
}

// Mettez votre HTML ici
var yourHTMLString = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hangman Game</title>
   <link rel="stylesheet" href="./static/styles.css">
</head>
<body>
  <div id="game" >
      <h1>hangman game</h1>
      <p>Find the hidden word - Enter one letter at a time</p>
          <div class="jeu-contenant">
                <svg height="250" width="200" class=figure-contenant>
                   <!--support-->
                   <line x1="30"  y1="220" x2="100" y2="220" />
                   <line x1="60"  y1="20" x2="60" y2="230" />
                   <line x1="60"  y1="20" x2="140" y2="20" />
                   <line x1="140"  y1="20" x2="140" y2="50" />
                   
                  <!--tete-->
                   <circle cx="140" cy="70" r="20" class="figure-partie"/>
              
                  <!--corps-->
                   <line x1="140"  y1="90" x2="140" y2="150" class="figure-partie"/>
                   <!--bras-->
                   <line x1="140"  y1="120" x2="120" y2="100" class="figure-partie"/>
                   <line x1="140"  y1="120" x2="160" y2="100" class="figure-partie"/>
                   <!--jambes-->
                   <line x1="140"  y1="150" x2="120" y2="180" class="figure-partie"/>
                   <line x1="140"  y1="150" x2="160" y2="180" class="figure-partie"/>
              </svg></div>
            <br>

            <div class="mauvaises-lettres-contenant"></div>
                     <p>wrong letter</p>
                      <div id="mauvaises-lettres"></div>
          <div id="mot" id="mot"></div>
            
                 <!--message final-->
            <div class="popup-contenant" id="popup-contenant" >
                      <div class="popup" >
                        <h2 id="message-final">bravo, tu as gagné</h2>
                        <button id="play-bouton">Rejouer</button>
                      </div>
            </div>
        <p> incorrect guesses : {{.Life}}</p>   <b>0 / 10</b> 
        <form action="/letter" method="post">
            <label for="letterInput">Entrez une lettre : </label>
            <input type="text" id="letterInput" name="letterInput" maxlength="1" pattern="[A-Za-z]" title="Entrez une seule lettre alphabétique" required>
            <button type="submit">Tester</button> 

             <!--notification-->
            <div class="notification-contenant" id="notification-contenant">
              <p>you have already tried this letter</p> 
            </div>
            <p>The correct word was : </p>
  </div>      
        </form>
</body>
</html>
`

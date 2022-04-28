package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// CSS-display (grid 등)
// CSS selector
// tag별 default CSS

var board Board

func main() {
	for i := range board.BigSquares {
		for j := range board.BigSquares[i].Squares {
			board.BigSquares[i].Squares[j].i = i
			board.BigSquares[i].Squares[j].j = j
		}
	}
	// Components routing:
	app.Route("/", &board)
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

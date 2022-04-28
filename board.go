package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Board struct {
	app.Compo
	Count      int
	BigSquares [9]BigSquare
}

func (b *Board) Render() app.UI {
	return app.Div().Class("board").Body(
		app.Div().Class("grid-outer").Body(
			app.Range(b.BigSquares).Slice(func(i int) app.UI {
				return app.Div().Class("grid-inner").Body(
					app.Range(b.BigSquares[i].Squares).Slice(func(j int) app.UI {
						return b.BigSquares[i].Squares[j].Render()
					}),
				)
			}),
		),
		app.Text(b.Status()),
	)
}
func (b *Board) HandleClick(i, j int) {
	if b.BigSquares[i].Squares[j].Value != Empty {
		return
	}
	if CalcWin(b.BigSquares[i].Squares) != Empty || CalcFinalWin(b.BigSquares) != Empty {
		return
	}

	b.BigSquares[i].Squares[j].Value = b.Turn()
}

func (b Board) Turn() int {
	var mark int
	switch b.Count % 2 {
	case 0:
		mark = Circle
	case 1:
		mark = Cross
	}
	return mark
}

var Wins = [...][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 7}}

func CalcWin(s [9]Square) int {
	for _, idxs := range Wins {
		zero := s[idxs[0]].Value
		one := s[idxs[1]].Value
		two := s[idxs[2]].Value
		if zero != Empty && zero == one && zero == two {
			return zero
		}
	}
	return Empty
}
func CalcFinalWin(ss [9]BigSquare) int {
	var s [9]Square
	for i := range s {
		s[i].Value = CalcWin(ss[i].Squares)
	}
	return CalcWin(s)
}

func (b Board) Status() string {
	if m := CalcFinalWin(b.BigSquares); m != Empty {
		return fmt.Sprintf("Winner: %s\n", Mark(m))
	}
	return fmt.Sprintf("Next player: %s\n", Mark(b.Count))
}

package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

const (
	Cross = -1 + iota
	Empty
	Circle
)

type BigSquare struct {
	app.Compo
	Squares [9]Square
	Value   int
}
type Square struct {
	app.Compo
	Value int
	i     int
	j     int
}

func (s *Square) OnClick(ctx app.Context, e app.Event) {
	board.HandleClick(s.i, s.j)
}

func (s *Square) Render() app.UI {
	return app.Div().Body(
		app.Button().Body(
			app.Button().Class("square"),
			app.Button().OnClick(s.OnClick),
			app.Text(Mark(s.Value)),
		),
	)
}

func Mark(m int) string {
	var s string
	switch m {
	case Cross:
		s = "❌"
	case Circle:
		s = "⚪"
	}
	return s
}

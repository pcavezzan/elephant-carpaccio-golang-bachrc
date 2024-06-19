package app

import "github.com/bachrc/test-stonal/internal/state"



type App struct {
	States []state.State
}

func New() App {
	return App{
		States: []state.State{{
			Name: "UTAH",
		}},
	}
}
 
func (app *App) ListStates() []state.State {
	return app.States;
}
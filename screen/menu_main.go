package screen

import (
	"eklase/state"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// MainMenu defines a main menu screen layout.
func MainMenu(th *material.Theme, state *state.State) Screen {
	var (
		add      widget.Clickable
		list     widget.Clickable
		quit     widget.Clickable
		explorer widget.Clickable
	)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(material.Button(th, &add, "Add student").Layout)),
			layout.Rigid(rowInset(material.Button(th, &list, "List students").Layout)),
			layout.Rigid(rowInset(material.Button(th, &quit, "Quit").Layout)),
			layout.Rigid(rowInset(material.Button(th, &explorer, "Explorer").Layout)),
		)
		if add.Clicked() {
			return AddStudent(th, state), d
		}
		if list.Clicked() {
			return ListStudent(th, state), d
		}
		if explorer.Clicked() {
			return Explorer(th, state), d
		}
		if quit.Clicked() {
			state.Quit()
		}
		return nil, d
	}
}

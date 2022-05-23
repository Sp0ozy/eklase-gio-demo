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
		matAddBut := material.Button(th, &add, "Add student")
		matAddBut.Font = ButtonFontMain()
		matAddBut.Background = ButtonBacgGroundMain()
		matListBut := material.Button(th, &list, "List students")
		matListBut.Font = ButtonFontMain()
		matListBut.Background = ButtonBacgGroundMain()
		matExpBut := material.Button(th, &explorer, "Explore")
		matExpBut.Font = ButtonFontMain()
		matExpBut.Background = ButtonBacgGroundMain()
		matQuitBut := material.Button(th, &quit, "Quit")
		matQuitBut.Font = ButtonFontMain()
		matQuitBut.Background = ButtonBacgGroundMain()

		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(matAddBut.Layout)),
			layout.Rigid(rowInset(matListBut.Layout)),
			layout.Rigid(rowInset(matExpBut.Layout)),
			layout.Rigid(rowInset(matQuitBut.Layout)),
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

package screen

import (
	"eklase/state"
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
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
		matAddBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matAddBut.Background = color.NRGBA{A: 0xff, R: 0x3C, G: 0x3C, B: 0x3C}
		matListBut := material.Button(th, &list, "List students")
		matListBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matListBut.Background = color.NRGBA{A: 0xff, R: 0x3C, G: 0x3C, B: 0x3C}
		matQuitBut := material.Button(th, &quit, "Quit")
		matQuitBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matQuitBut.Background = color.NRGBA{A: 0xff, R: 0x3C, G: 0x3C, B: 0x3C}
		matExpBut := material.Button(th, &explorer, "Explore")
		matExpBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matExpBut.Background = color.NRGBA{A: 0xff, R: 0x3C, G: 0x3C, B: 0x3C}

		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(matAddBut.Layout)),
			layout.Rigid(rowInset(matListBut.Layout)),
			layout.Rigid(rowInset(matQuitBut.Layout)),
			layout.Rigid(rowInset(matExpBut.Layout)),
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

package screen

import (
	"eklase/explorer"
	"eklase/state"
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Explorera(th *material.Theme, state *state.State) Screen {
	d := explorer.Drives()
	var root []string
	fmt.Println(d, len(d))
	listf := explorer.List(root)

	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}

	lightContrast := th.ContrastBg
	lightContrast.A = 0x11
	darkContrast := th.ContrastBg
	darkContrast.A = 0x33

	filesLayout := func(gtx layout.Context) layout.Dimensions {
		return material.List(th, &list).Layout(gtx, len(listf), func(gtx layout.Context, index int) layout.Dimensions {
			file := listf[index]
			fmt.Println(file)
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					color := lightContrast
					if index%2 == 0 {
						color = darkContrast
					}
					max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
					paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Stacked(rowInset(material.Body1(th, file).Layout)),
			)
		})
	}

	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(1, rowInset(filesLayout)),
			layout.Rigid(rowInset(material.Button(th, &close, "Close").Layout)),
		)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}

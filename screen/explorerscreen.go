package screen

import (
	"eklase/explorer"
	"eklase/state"
	"eklase/storage"
	"image"
	"log"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func generateFileList(th *material.Theme, list widget.List, files []string, button []widget.Clickable) func(gtx layout.Context) layout.Dimensions {
	lightContrast := th.ContrastBg
	lightContrast.A = 0x11
	darkContrast := th.ContrastBg
	darkContrast.A = 0x33
	return func(gtx layout.Context) layout.Dimensions {
		return material.List(th, &list).Layout(gtx, len(files), func(gtx layout.Context, index int) layout.Dimensions {
			file := files[index]

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
				layout.Stacked(rowInset(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(material.Button(th, &button[index], file).Layout),
					)
				})),
			)
		})
	}
}

func Explorer(th *material.Theme, state *state.State) Screen {
	var files, f, root []string
	root = explorer.GetRoot()
	if root == nil {
		f = explorer.Drives()
	} else {
		f = explorer.List(root)
	}
	for _, str := range f {
		var cnt int
		if str[0] == '.' {
			continue
		}
		for i := range str {
			if str[i] == '.' {
				cnt++
				if str[i+1] == 'd' && str[i+2] == 'b' && len(str) == i+3 {
					files = append(files, str)
					continue
				}
			}
		}
		if cnt == 0 {
			files = append(files, str)
		}
	}

	var close widget.Clickable
	var back widget.Clickable

	list := widget.List{List: layout.List{Axis: layout.Vertical}}
	students, err := state.Students()
	if err != nil {
		log.Printf("failed to fetch students: %v", err)
		return nil
	}

	button := make([]widget.Clickable, len(students))

	studentsLayout := generateFileList(th, list, files, button)

	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(1, rowInset(studentsLayout)),
			layout.Rigid(rowInset(material.Button(th, &back, "Back").Layout)),
			layout.Rigid(rowInset(material.Button(th, &close, "Close").Layout)),
		)
		for i := range button {
			if button[i].Clicked() {
				if checnkdb(files[i]) {
					storage.Open(explorer.Root(root) + files[i])
					return MainMenu(th, state), d
				}
				root = append(root, files[i])
				explorer.SaveRoot(root)
				return Explorer(th, state), d
			}
		}
		if back.Clicked() {
			if root == nil {
				return Explorer(th, state), d
			}
			root = root[:len(root)-1]
			explorer.SaveRoot(root)
			return Explorer(th, state), d
		}
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}

func checnkdb(str string) bool {
	for i := range str {
		if str[i] == '.' {

			if str[i+1] == 'd' && str[i+2] == 'b' && len(str) == i+3 {
				return true

			}
		}
	}
	return false
}

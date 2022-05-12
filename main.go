package main

import (
	"eklase/explorer"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {

	// // Run the main event loop.
	// go func() {
	// 	w := app.NewWindow(app.Title("e-Klasse"))
	// 	if err := mainLoop(w); err != nil {
	// 		log.Fatalf("failed to handle events: %v", err)
	// 	}
	// 	// Gracefully exit the application at the end.
	// 	os.Exit(0)
	// }()
	// app.Main()

	//temporary code
	d := explorer.Drives()
	var root []string
	fmt.Println("which dir to use?")
	var in int
	fmt.Scanln(&in)
	root = append(root, d[in-1])
	for {
		list := explorer.List(root)
		fmt.Println("which dir to use?")
		fmt.Scanln(&in)
		if in == 0 {
			root = root[:len(root)-1]
		} else {
			root = append(root, list[in-1])
		}

	}

}

// func mainLoop(w *app.Window) error {
// 	storage := storage.Must(storage.New("school.db"))
// 	defer storage.Close()

// 	appState := state.New(storage)

// 	th := material.NewTheme(gofont.Collection())
// 	currentLayout := screen.MainMenu(th, appState)

// 	for {
// 		select {
// 		case e := <-w.Events():
// 			switch e := e.(type) {
// 			case system.FrameEvent:
// 				gtx := layout.NewContext(&op.Ops{}, e)
// 				layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
// 					nextLayout, d := currentLayout(gtx)
// 					if nextLayout != nil {
// 						currentLayout = nextLayout
// 					}
// 					return d
// 				})
// 				if appState.ShouldQuit() {
// 					w.Perform(system.ActionClose)
// 				}
// 				e.Frame(gtx.Ops)
// 			case system.DestroyEvent:
// 				return e.Err
// 			}
// 		}
// 	}
// }

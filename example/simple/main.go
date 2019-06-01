package main

import (
	"fmt"
	"os"

	"github.com/ysh86/gui"
)

func main() {
	app := gui.NewApplication()

	if err := app.Init(); err != nil {
		panic(err)
	}
	defer app.Deinit()

	renderer, err := NewD2D1Renderer()
	if err != nil {
		panic(err)
	}

	windowName := "Direct2D Demo App"
	errc := app.Loop(windowName, 640, 480, renderer)
	select {
	case e := <-errc:
		if e != nil {
			panic(e)
		}
	}

	fmt.Fprintln(os.Stderr, "Done")
}

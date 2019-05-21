package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/ysh86/gui"
	"github.com/ysh86/svg"
)

func main() {
	// load
	dec := xml.NewDecoder(os.Stdin)

	svg := new(svg.Root)
	err := svg.Parse(dec)
	if err != nil {
		panic(err)
	}
	fmt.Println("load done")

	// app
	app := gui.NewApplication()

	if err := app.Init(); err != nil {
		panic(err)
	}
	defer app.Deinit()

	// renderer
	// TODO: logger
	renderer, err := NewSVGRenderer(svg)
	if err != nil {
		panic(err)
	}

	// run
	windowName := "Direct2D SVG Viewer"
	errc := app.Loop(
		windowName,
		int32(svg.ViewBox.Width),
		int32(svg.ViewBox.Height),
		renderer)
	select {
	case e := <-errc:
		if e != nil {
			panic(e)
		}
	}

	fmt.Fprintln(os.Stderr, "Done")
}

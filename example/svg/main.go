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
	fmt.Printf("load from stdin...")
	dec := xml.NewDecoder(os.Stdin)

	svg := new(svg.Root)
	err := svg.Parse(dec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("done")
	fmt.Printf(" (WxH: %dx%d)\n", svg.ViewBox.Width, svg.ViewBox.Height)

	// app
	app := gui.NewApplication()

	if err := app.Init(); err != nil {
		panic(err)
	}
	defer app.Deinit()

	// renderer
	renderer, err := NewSVGRenderer(svg)
	if err != nil {
		panic(err)
	}
	renderer.enableLog()

	// run
	fmt.Printf("drawing...")
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
	fmt.Println("done")
}

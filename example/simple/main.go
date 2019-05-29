// Creating a Simple Direct2D Application
// https://docs.microsoft.com/en-us/windows/desktop/direct2d/direct2d-quickstart

package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"

	d2d "github.com/ysh86/go-d2d"
	"github.com/ysh86/gui"
)

type d2d1Renerer struct {
	factory *d2d.ID2D1Factory

	dpiX float32
	dpiY float32

	renderTarget        *d2d.ID2D1HwndRenderTarget
	lightSlateGrayBrush *d2d.ID2D1Brush
	cornflowerBlueBrush *d2d.ID2D1Brush
}

// Init calls methods for instantiating drawing resources
func (r *d2d1Renerer) Init() error {
	err := gui.CoInitializeEx(0, gui.COINIT_APARTMENTTHREADED|gui.COINIT_DISABLE_OLE1DDE)
	if err != nil {
		return fmt.Errorf("CoInitializeEx")
	}

	// Initialize device-indpendent resources, such
	// as the Direct2D factory.
	err = r.createDeviceIndependentResources()
	if err != nil {
		return err
	}

	// Because the CreateWindow function takes its size in pixels,
	// obtain the system DPI and use it to scale the window size.
	r.dpiX, r.dpiY = r.factory.GetDesktopDpi()

	return nil
}

// Deinit releases resources
func (r *d2d1Renerer) Deinit() {
	r.discardDeviceResources()

	if r.factory != nil {
		r.factory.Release()
		r.factory = nil
	}

	gui.CoUninitialize()
}

func (r *d2d1Renerer) Dpi() (float32, float32) {
	return r.dpiX, r.dpiY
}

func (r *d2d1Renerer) Update(width, height uint32) error {
	if r.renderTarget != nil {
		r.renderTarget.Resize(
			&d2d.D2D1_SIZE_U{Width: width, Height: height})
	}

	return nil
}

func (r *d2d1Renerer) Draw(nativeWindow uintptr) error {
	err := r.createDeviceResources(nativeWindow)
	if err != nil {
		return err
	}

	r.renderTarget.BeginDraw()

	identityMatrix := d2d.D2D1_MATRIX_3X2_F{
		A11: 1.,
		A22: 1.,
	}
	r.renderTarget.SetTransform(&identityMatrix)

	white := d2d.D2D1_COLOR_F{R: 1, G: 1, B: 1, A: 1}
	r.renderTarget.Clear(&white)

	size := r.renderTarget.GetSize()

	// Draw a grid background.
	width := int(size.Width)
	height := int(size.Height)
	for x := 0; x < width; x += 10 {
		r.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{X: float32(x), Y: 0.0},
			d2d.D2D1_POINT_2F{X: float32(x), Y: size.Height},
			r.lightSlateGrayBrush,
			0.5,
			nil)
	}
	for y := 0; y < height; y += 10 {
		r.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{X: 0.0, Y: float32(y)},
			d2d.D2D1_POINT_2F{X: size.Width, Y: float32(y)},
			r.lightSlateGrayBrush,
			0.5,
			nil)
	}

	// Draw two rectangles.
	rectangle1 := d2d.D2D1_RECT_F{
		Left:   size.Width/2.0 - 50.0,
		Top:    size.Height/2.0 - 50.0,
		Right:  size.Width/2.0 + 50.0,
		Bottom: size.Height/2.0 + 50.0,
	}
	rectangle2 := d2d.D2D1_RECT_F{
		Left:   size.Width/2.0 - 100.0,
		Top:    size.Height/2.0 - 100.0,
		Right:  size.Width/2.0 + 100.0,
		Bottom: size.Height/2.0 + 100.0,
	}
	// Draw a filled rectangle.
	r.renderTarget.FillRectangle(
		&rectangle1,
		r.lightSlateGrayBrush)
	// Draw the outline of a rectangle.
	r.renderTarget.DrawRectangle(
		&rectangle2,
		r.cornflowerBlueBrush,
		1,
		nil)

	_, _, err = r.renderTarget.EndDraw()
	if err != nil {
		r.discardDeviceResources()
	}

	return nil
}

// private methods

func (r *d2d1Renerer) createDeviceIndependentResources() (err error) {
	r.factory, err = d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
	return
}

func (r *d2d1Renerer) createDeviceResources(nativeWindow uintptr) error {
	if r.renderTarget != nil {
		return nil // already created
	}

	hwnd := windows.Handle(nativeWindow)

	var rc gui.Rect
	err := gui.GetClientRect(hwnd, &rc)
	if err != nil {
		return err
	}
	r.renderTarget, err = r.factory.CreateHwndRenderTarget(
		d2d.RenderTargetProperties(),
		&d2d.D2D1_HWND_RENDER_TARGET_PROPERTIES{
			Hwnd: uintptr(hwnd),
			PixelSize: d2d.D2D1_SIZE_U{
				Width:  uint32(rc.Right - rc.Left),
				Height: uint32(rc.Bottom - rc.Top)},
			PresentOptions: d2d.D2D1_PRESENT_OPTIONS_NONE,
		})
	if err != nil {
		return err
	}

	lightSlateGray := d2d.D2D1_COLOR_F{R: 0x77 / 255., G: 0x88 / 255., B: 0x99 / 255., A: 1}
	lightSlateGrayBrush, err := r.renderTarget.CreateSolidColorBrush(
		&lightSlateGray,
		nil)
	if err != nil {
		return err
	}
	r.lightSlateGrayBrush = &(lightSlateGrayBrush.ID2D1Brush)

	cornflowerBlue := d2d.D2D1_COLOR_F{R: 0x64 / 255., G: 0x95 / 255., B: 0xED / 255., A: 1}
	cornflowerBlueBrush, err := r.renderTarget.CreateSolidColorBrush(
		&cornflowerBlue,
		nil)
	if err != nil {
		return err
	}
	r.cornflowerBlueBrush = &(cornflowerBlueBrush.ID2D1Brush)

	return nil
}

func (r *d2d1Renerer) discardDeviceResources() {
	if r.renderTarget != nil {
		r.renderTarget.Release()
		r.renderTarget = nil
	}
	if r.lightSlateGrayBrush != nil {
		r.lightSlateGrayBrush.Release()
		r.lightSlateGrayBrush = nil
	}
	if r.cornflowerBlueBrush != nil {
		r.cornflowerBlueBrush.Release()
		r.cornflowerBlueBrush = nil
	}
}

func main() {
	app := gui.NewApplication()

	if err := app.Init(); err != nil {
		panic(err)
	}
	defer app.Deinit()

	windowName := "Direct2D Demo App"
	renderer := &d2d1Renerer{}
	errc := app.Loop(windowName, 640, 480, renderer)
	select {
	case e := <-errc:
		if e != nil {
			panic(e)
		}
	}

	fmt.Fprintln(os.Stderr, "Done")
}

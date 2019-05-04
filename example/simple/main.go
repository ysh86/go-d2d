// Creating a Simple Direct2D Application
// https://docs.microsoft.com/en-us/windows/desktop/direct2d/direct2d-quickstart

package main

import (
	"fmt"
	"math"
	"unsafe"

	"golang.org/x/sys/windows"

	d2d "github.com/ysh86/go-d2d"
	"github.com/ysh86/gui"
)

// DemoApp is the app
type DemoApp struct {
	instance windows.Handle
	cmdLine  string
	cmdShow  int32
	atom     gui.Atom

	hwnd windows.Handle

	factory             *d2d.ID2D1Factory
	renderTarget        *d2d.ID2D1HwndRenderTarget
	lightSlateGrayBrush *d2d.ID2D1Brush
	cornflowerBlueBrush *d2d.ID2D1Brush
}

// Initialize registers the window class and call methods for instantiating drawing resources
func (app *DemoApp) Initialize() error {
	// Initialize device-indpendent resources, such
	// as the Direct2D factory.
	err := app.createDeviceIndependentResources()
	if err != nil {
		return err
	}

	// dummy _tWinMain()
	i, err := gui.GetModuleHandle(nil)
	if err != nil {
		return fmt.Errorf("GetModuleHandle: %v", err)
	}
	app.instance = i
	app.cmdLine = ""
	app.cmdShow = gui.SW_SHOWNORMAL

	// Register the window class.
	className := "D2DDemoApp"
	classNameUTF16, err := windows.UTF16PtrFromString(className)
	if err != nil {
		return fmt.Errorf("UTF16PtrFromString %s: %v", className, err)
	}
	//icon := w32.LoadIcon(0, w32.MakeIntResource(w32.IDI_APPLICATION))
	wndClass := &gui.WndClassEx{
		Size:       0,
		Style:      gui.CS_HREDRAW | gui.CS_VREDRAW,
		WndProc:    windows.NewCallback(app.wndProc),
		ClsExtra:   0,
		WndExtra:   0,
		Instance:   app.instance,
		Icon:       0, //icon,
		Cursor:     0, //w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)), // LoadCursor(NULL, IDI_APPLICATION),
		Background: windows.Handle(gui.COLOR_WINDOW + 1),
		MenuName:   nil,
		ClassName:  classNameUTF16,
		IconSm:     0, //icon,
	}
	wndClass.Size = uint32(unsafe.Sizeof(*wndClass))
	atom, err := gui.RegisterClassEx(wndClass)
	if err != nil {
		return fmt.Errorf("RegisterClassEx %v: %v", wndClass, err)
	}
	app.atom = atom

	// Because the CreateWindow function takes its size in pixels,
	// obtain the system DPI and use it to scale the window size.

	// The factory returns the current system DPI. This is also the value it will use
	// to create its own windows.
	dpiX, dpiY := app.factory.GetDesktopDpi()

	windowName := "Direct2D Demo App"
	windowNameUTF16, err := windows.UTF16PtrFromString(windowName)
	if err != nil {
		return fmt.Errorf("UTF16PtrFromString %s: %v", windowName, err)
	}
	w, err := gui.CreateWindowEx(
		0,
		(*uint16)(unsafe.Pointer(uintptr(app.atom))),
		windowNameUTF16,
		gui.WS_OVERLAPPEDWINDOW,
		gui.CW_USEDEFAULT, gui.CW_USEDEFAULT, // x, y
		int32(math.Ceil(float64(640.0*dpiX/96.0))), // width
		int32(math.Ceil(float64(480.0*dpiY/96.0))), // height
		0,
		0,
		app.instance,
		uintptr(unsafe.Pointer(app)),
	)
	if err != nil {
		return fmt.Errorf("CreateWindowEx: %v", err)
	}
	app.hwnd = w

	_ = gui.ShowWindow(w, app.cmdShow) // ignore return value
	err = gui.UpdateWindow(w)
	if err != nil {
		return fmt.Errorf("UpdateWindow: %v", err)
	}

	return nil
}

// Deinitialize releases resources
func (app *DemoApp) Deinitialize() {
	if app.factory != nil {
		app.factory.Release()
		app.factory = nil
	}
	if app.renderTarget != nil {
		app.renderTarget.Release()
		app.renderTarget = nil
	}
	if app.lightSlateGrayBrush != nil {
		app.lightSlateGrayBrush.Release()
		app.lightSlateGrayBrush = nil
	}
	if app.cornflowerBlueBrush != nil {
		app.cornflowerBlueBrush.Release()
		app.cornflowerBlueBrush = nil
	}
}

// RunMessageLoop processes and dispatches messages
func (app *DemoApp) RunMessageLoop() (err error) {
	var msg gui.Msg
	for {
		if result, e := gui.GetMessage(&msg, 0, 0, 0); e != nil || result == 0 {
			err = e
			break
		}
		gui.TranslateMessage(&msg)
		gui.DispatchMessage(&msg)
	}
	return err
}

// private methods

func (app *DemoApp) createDeviceIndependentResources() (err error) {
	app.factory, err = d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
	return
}

func (app *DemoApp) createDeviceResources() error {
	if app.renderTarget != nil {
		return nil // already created
	}

	var rc gui.Rect
	err := gui.GetClientRect(app.hwnd, &rc)
	if err != nil {
		return err
	}
	app.renderTarget, err = app.factory.CreateHwndRenderTarget(
		d2d.RenderTargetProperties(),
		&d2d.D2D1_HWND_RENDER_TARGET_PROPERTIES{
			Hwnd: uintptr(app.hwnd),
			PixelSize: d2d.D2D1_SIZE_U{
				Width:  uint32(rc.Right - rc.Left),
				Height: uint32(rc.Bottom - rc.Top)},
			PresentOptions: d2d.D2D1_PRESENT_OPTIONS_NONE,
		})
	if err != nil {
		return err
	}

	lightSlateGray := d2d.D2D1_COLOR_F{R: 0x77 / 255., G: 0x88 / 255., B: 0x99 / 255., A: 1}
	lightSlateGrayBrush, err := app.renderTarget.CreateSolidColorBrush(
		&lightSlateGray,
		nil)
	if err != nil {
		return err
	}
	app.lightSlateGrayBrush = &(lightSlateGrayBrush.ID2D1Brush)

	cornflowerBlue := d2d.D2D1_COLOR_F{R: 0x64 / 255., G: 0x95 / 255., B: 0xED / 255., A: 1}
	cornflowerBlueBrush, err := app.renderTarget.CreateSolidColorBrush(
		&cornflowerBlue,
		nil)
	if err != nil {
		return err
	}
	app.cornflowerBlueBrush = &(cornflowerBlueBrush.ID2D1Brush)

	return nil
}

func (app *DemoApp) discardDeviceResources() {
	if app.renderTarget != nil {
		app.renderTarget.Release()
		app.renderTarget = nil
	}
	if app.lightSlateGrayBrush != nil {
		app.lightSlateGrayBrush.Release()
		app.lightSlateGrayBrush = nil
	}
	if app.cornflowerBlueBrush != nil {
		app.cornflowerBlueBrush.Release()
		app.cornflowerBlueBrush = nil
	}
}

func (app *DemoApp) onRender() {
	err := app.createDeviceResources()
	if err != nil {
		return
	}

	app.renderTarget.BeginDraw()

	identityMatrix := d2d.D2D1_MATRIX_3X2_F{
		A11: 1.,
		A22: 1.,
	}
	app.renderTarget.SetTransform(&identityMatrix)

	white := d2d.D2D1_COLOR_F{R: 1, G: 1, B: 1, A: 1}
	app.renderTarget.Clear(&white)

	size := app.renderTarget.GetSize()

	// Draw a grid background.
	width := int(size.Width)
	height := int(size.Height)
	for x := 0; x < width; x += 10 {
		app.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{X: float32(x), Y: 0.0},
			d2d.D2D1_POINT_2F{X: float32(x), Y: size.Height},
			app.lightSlateGrayBrush,
			0.5,
			nil)
	}
	for y := 0; y < height; y += 10 {
		app.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{X: 0.0, Y: float32(y)},
			d2d.D2D1_POINT_2F{X: size.Width, Y: float32(y)},
			app.lightSlateGrayBrush,
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
	app.renderTarget.FillRectangle(
		&rectangle1,
		app.lightSlateGrayBrush)
	// Draw the outline of a rectangle.
	app.renderTarget.DrawRectangle(
		&rectangle2,
		app.cornflowerBlueBrush,
		1,
		nil)

	_, _, err = app.renderTarget.EndDraw()
	if err != nil {
		app.discardDeviceResources()
	}
}

func (app *DemoApp) onResize(w, h uint32) {
	if app.renderTarget != nil {
		app.renderTarget.Resize(
			&d2d.D2D1_SIZE_U{Width: w, Height: h})
	}
}

func (app *DemoApp) wndProc(hwnd windows.Handle, msg uint32, wParam, lParam uintptr) uintptr {
	if hwnd == app.hwnd {
		switch msg {
		case gui.WM_SIZE:
			width := uint32(gui.LOWORD(lParam))
			height := uint32(gui.HIWORD(lParam))
			app.onResize(width, height)
			return 0
		case gui.WM_DISPLAYCHANGE:
			gui.InvalidateRect(hwnd, nil, false)
			return 0
		case gui.WM_PAINT:
			app.onRender()
			gui.ValidateRect(hwnd, nil)
			return 0
		case gui.WM_DESTROY:
			gui.PostQuitMessage(0)
			return 1
		}
	}

	r, _ := gui.DefWindowProc(hwnd, msg, wParam, lParam)
	return r
}

func main() {
	err := gui.CoInitializeEx(0, gui.COINIT_MULTITHREADED)
	if err != nil {
		panic(err)
	}
	defer gui.CoUninitialize()

	app := new(DemoApp)

	err = app.Initialize()
	if err != nil {
		panic(err)
	}
	defer app.Deinitialize()

	err = app.RunMessageLoop()
	if err != nil {
		panic(err)
	}
}

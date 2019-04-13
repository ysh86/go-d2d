// Creating a Simple Direct2D Application
// https://docs.microsoft.com/en-us/windows/desktop/direct2d/direct2d-quickstart

package main

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/AllenDang/w32"
	d2d "github.com/ysh86/go-d2d"
)

// DemoApp is the app
type DemoApp struct {
	hwnd                w32.HWND
	factory             *d2d.ID2D1Factory
	renderTarget        *d2d.ID2D1HwndRenderTarget
	lightSlateGrayBrush *d2d.ID2D1Brush
	cornflowerBlueBrush *d2d.ID2D1Brush
}

// Initialize registers the window class and call methods for instantiating drawing resources
func (app *DemoApp) Initialize() {
	// Initialize device-indpendent resources, such
	// as the Direct2D factory.
	app.createDeviceIndependentResources()

	// Register the window class.
	hInstance := w32.GetModuleHandle("")
	icon := w32.LoadIcon(0, w32.MakeIntResource(w32.IDI_APPLICATION))
	wndProc := func(hwnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
		return app.wndProc(hwnd, msg, wParam, lParam)
	}
	wndClass := w32.WNDCLASSEX{
		Size:       uint32(unsafe.Sizeof(w32.WNDCLASSEX{})),
		Style:      w32.CS_HREDRAW | w32.CS_VREDRAW,
		WndProc:    syscall.NewCallback(wndProc),
		ClsExtra:   0,
		WndExtra:   0,
		Instance:   hInstance,
		Icon:       icon,
		Cursor:     w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		Background: 0,
		MenuName:   nil,
		ClassName:  syscall.StringToUTF16Ptr("D2DDemoApp"),
		IconSm:     icon,
	}
	w32.RegisterClassEx(&wndClass)

	// Because the CreateWindow function takes its size in pixels,
	// obtain the system DPI and use it to scale the window size.

	// The factory returns the current system DPI. This is also the value it will use
	// to create its own windows.
	dpiX, dpiY := app.factory.GetDesktopDpi()

	app.hwnd = w32.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr("D2DDemoApp"),
		syscall.StringToUTF16Ptr("Hello Windows"),
		w32.WS_OVERLAPPEDWINDOW,
		w32.CW_USEDEFAULT,
		w32.CW_USEDEFAULT,
		int(math.Ceil(float64(640*dpiX/96))),
		int(math.Ceil(float64(480*dpiY/96))),
		0,
		0,
		hInstance,
		nil)
	w32.ShowWindow(app.hwnd, w32.SW_SHOW)
	w32.UpdateWindow(app.hwnd)
}

//Dispose releases resources
func (app *DemoApp) Dispose() {
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
func (app *DemoApp) RunMessageLoop() {
	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) > 0 {
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

// private methods

func (app *DemoApp) createDeviceIndependentResources() {
	app.factory, _ = d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
}

func (app *DemoApp) createDeviceResources() {
	if app.renderTarget == nil {
		rc := w32.GetClientRect(app.hwnd)
		app.renderTarget, _ = app.factory.CreateHwndRenderTarget(
			d2d.RenderTargetProperties(),
			&d2d.D2D1_HWND_RENDER_TARGET_PROPERTIES{
				Hwnd: uintptr(unsafe.Pointer(app.hwnd)),
				PixelSize: d2d.D2D1_SIZE_U{
					Width:  uint32(rc.Right - rc.Left),
					Height: uint32(rc.Bottom - rc.Top)},
				PresentOptions: d2d.D2D1_PRESENT_OPTIONS_NONE,
			})

		lightSlateGray := d2d.D2D1_COLOR_F{0x77 / 255., 0x88 / 255., 0x99 / 255., 1}
		lightSlateGrayBrush, _ := app.renderTarget.CreateSolidColorBrush(
			&lightSlateGray,
			nil)
		app.lightSlateGrayBrush = &(lightSlateGrayBrush.ID2D1Brush)

		cornflowerBlue := d2d.D2D1_COLOR_F{0x64 / 255., 0x95 / 255., 0xED / 255., 1}
		cornflowerBlueBrush, _ := app.renderTarget.CreateSolidColorBrush(
			&cornflowerBlue,
			nil)
		app.cornflowerBlueBrush = &(cornflowerBlueBrush.ID2D1Brush)
	}
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
	app.createDeviceResources()

	app.renderTarget.BeginDraw()

	identityMatrix := d2d.D2D1_MATRIX_3X2_F{
		A11: 1.,
		A22: 1.,
	}
	app.renderTarget.SetTransform(&identityMatrix)

	white := d2d.D2D1_COLOR_F{1, 1, 1, 1}
	app.renderTarget.Clear(&white)

	size := app.renderTarget.GetSize()

	// Draw a grid background.
	width := int(size.Width)
	height := int(size.Height)
	for x := 0; x < width; x += 10 {
		app.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{float32(x), 0.0},
			d2d.D2D1_POINT_2F{float32(x), size.Height},
			app.lightSlateGrayBrush,
			0.5,
			nil)
	}
	for y := 0; y < height; y += 10 {
		app.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{0.0, float32(y)},
			d2d.D2D1_POINT_2F{size.Width, float32(y)},
			app.lightSlateGrayBrush,
			0.5,
			nil)
	}

	// Draw two rectangles.
	rectangle1 := d2d.D2D1_RECT_F{
		size.Width/2.0 - 50.0,
		size.Height/2.0 - 50.0,
		size.Width/2.0 + 50.0,
		size.Height/2.0 + 50.0,
	}
	rectangle2 := d2d.D2D1_RECT_F{
		size.Width/2.0 - 100.0,
		size.Height/2.0 - 100.0,
		size.Width/2.0 + 100.0,
		size.Height/2.0 + 100.0,
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

	app.renderTarget.EndDraw()

	// TODO: error handling
	var err error
	if err != nil {
		app.discardDeviceResources()
	}
}

func (app *DemoApp) onResize(w, h uint32) {
	if app.renderTarget != nil {
		app.renderTarget.Resize(
			&d2d.D2D1_SIZE_U{w, h})
	}
}

func (app *DemoApp) wndProc(hwnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if hwnd != app.hwnd {
		return w32.DefWindowProc(hwnd, msg, wParam, lParam)
	}
	switch msg {
	case w32.WM_SIZE:
		width := uint32(w32.LOWORD(uint32(lParam)))
		height := uint32(w32.HIWORD(uint32(lParam)))
		app.onResize(width, height)
		return 0
	case w32.WM_DISPLAYCHANGE:
		w32.InvalidateRect(app.hwnd, nil, false)
		return 0
	case w32.WM_PAINT:
		app.onRender()
		w32.ValidateRect(app.hwnd, nil)
		return 0
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
		return 1
	}
	return w32.DefWindowProc(hwnd, msg, wParam, lParam)
}

func main() {
	w32.CoInitialize()
	defer w32.CoUninitialize()

	app := new(DemoApp)
	defer app.Dispose()

	app.Initialize()
	app.RunMessageLoop()
}

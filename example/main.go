// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/AllenDang/w32"
	d2d "github.com/ysh86/go-d2d"
)

type DemoApp struct {
	hwnd                   w32.HWND
	factory                *d2d.ID2D1Factory
	render_target          *d2d.ID2D1HwndRenderTarget
	light_slate_gray_brush *d2d.ID2D1Brush
	cornflower_blue        *d2d.ID2D1Brush
}

func (app *DemoApp) Initialize() {
	app.CreateDeviceIndependentResources()
	hInstance := w32.GetModuleHandle("")
	icon := w32.LoadIcon(0, w32.MakeIntResource(w32.IDI_APPLICATION))
	wndProc := func(hwnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
		return app.WndProc(hwnd, msg, wParam, lParam)
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

	dpiX, dpiY := app.factory.GetDesktopDpi()

	app.hwnd = w32.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr("D2DDemoApp"),
		syscall.StringToUTF16Ptr("Hello Windows"),
		w32.WS_OVERLAPPEDWINDOW,
		w32.CW_USEDEFAULT,
		w32.CW_USEDEFAULT,
		int(math.Ceil(float64(640*dpiX/96))),
		int(math.Ceil(float64(640*dpiY/96))),
		0,
		0,
		hInstance,
		nil)
	w32.ShowWindow(app.hwnd, w32.SW_SHOW)
	w32.UpdateWindow(app.hwnd)
}

func (app *DemoApp) Dispose() {
	// TODO: safe release がダサい
	if app.factory != nil {
		app.factory.Release()
		app.factory = nil
	}
	if app.render_target != nil {
		app.render_target.Release()
		app.render_target = nil
	}
	if app.light_slate_gray_brush != nil {
		app.light_slate_gray_brush.Release()
		app.light_slate_gray_brush = nil
	}
	if app.cornflower_blue != nil {
		app.cornflower_blue.Release()
		app.cornflower_blue = nil
	}
}

func (app *DemoApp) RunMessageLoop() {
	msg := w32.MSG{}
	for w32.GetMessage(&msg, 0, 0, 0) > 0 {
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

func (app *DemoApp) CreateDeviceIndependentResources() {
	app.factory, _ = d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
}

func (app *DemoApp) CreateDeviceResources() {
	if app.render_target == nil {
		var rc = w32.GetClientRect(app.hwnd)
		var hwndRenderTargetProperties = d2d.HwndRenderTargetProperties(uintptr(unsafe.Pointer(app.hwnd)))
		hwndRenderTargetProperties.PixelSize = d2d.D2D1_SIZE_U{
			uint32(rc.Right - rc.Left),
			uint32(rc.Bottom - rc.Top),
		}
		app.render_target, _ = app.factory.CreateHwndRenderTarget(
			d2d.RenderTargetProperties(),
			hwndRenderTargetProperties)
		var LightSlateGray = d2d.D2D1_COLOR_F{0x77 / 255., 0x88 / 255., 0x99 / 255., 1}
		var CornflowerBlue = d2d.D2D1_COLOR_F{0x64 / 255., 0x95 / 255., 0xED / 255., 1}
		light_slate_gray_brush, _ := app.render_target.CreateSolidColorBrush(
			&LightSlateGray,
			nil)
		defer light_slate_gray_brush.Release()
		p, _ := light_slate_gray_brush.QueryInterface(&d2d.IID_ID2D1Brush)
		// TODO: QueryInterface が生すぎる
		app.light_slate_gray_brush = (*d2d.ID2D1Brush)(p)
		cornflower_blue, _ := app.render_target.CreateSolidColorBrush(
			&CornflowerBlue,
			nil)
		defer cornflower_blue.Release()
		p, _ = cornflower_blue.QueryInterface(&d2d.IID_ID2D1Brush)
		app.cornflower_blue = (*d2d.ID2D1Brush)(p)
	}
}

func (app *DemoApp) DiscardDeviceResources() {
	if app.render_target != nil {
		app.render_target.Release()
		app.render_target = nil
	}
	if app.light_slate_gray_brush != nil {
		app.light_slate_gray_brush.Release()
		app.light_slate_gray_brush = nil
	}
	if app.cornflower_blue != nil {
		app.cornflower_blue.Release()
		app.cornflower_blue = nil
	}
}

func (app *DemoApp) WndProc(hwnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if hwnd != app.hwnd {
		return w32.DefWindowProc(hwnd, msg, wParam, lParam)
	}
	switch msg {
	case w32.WM_SIZE:
		width := w32.LOWORD(uint32(lParam))
		height := w32.HIWORD(uint32(lParam))
		app.OnResize(width, height)
		return 0
	case w32.WM_DISPLAYCHANGE:
		w32.InvalidateRect(app.hwnd, nil, false)
		return 0
	case w32.WM_PAINT:
		app.OnRender()
		w32.ValidateRect(app.hwnd, nil)
		return 0
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
		return 1
	}

	return w32.DefWindowProc(hwnd, msg, wParam, lParam)
}

func (app *DemoApp) OnResize(w, h uint16) {
	if app.render_target != nil {
		app.render_target.Resize(
			&d2d.D2D1_SIZE_U{uint32(w), uint32(h)})
	}
}

func (app *DemoApp) OnRender() {
	app.CreateDeviceResources()
	app.render_target.BeginDraw()
	var identityMatrix = d2d.D2D1_MATRIX_3X2_F{
		A11: 1.,
		A22: 1.,
	}
	app.render_target.SetTransform(&identityMatrix)
	var White = d2d.D2D1_COLOR_F{1, 1, 1, 1}
	app.render_target.Clear(&White)
	size := app.render_target.GetSize()
	width := int(size.Width)
	height := int(size.Height)
	for x := 0; x < width; x += 10 {
		app.render_target.DrawLine(
			d2d.D2D1_POINT_2F{float32(x), 0},
			d2d.D2D1_POINT_2F{float32(x), size.Height},
			app.light_slate_gray_brush,
			0.5,
			nil)
	}
	for y := 0; y < height; y += 10 {
		app.render_target.DrawLine(
			d2d.D2D1_POINT_2F{0, float32(y)},
			d2d.D2D1_POINT_2F{size.Width, float32(y)},
			app.light_slate_gray_brush,
			0.5,
			nil)
	}
	rectangle1 := d2d.D2D1_RECT_F{
		size.Width/2 - 50,
		size.Height/2 - 50,
		size.Width/2 + 50,
		size.Height/2 + 50,
	}
	rectangle2 := d2d.D2D1_RECT_F{
		size.Width/2 - 100,
		size.Height/2 - 100,
		size.Width/2 + 100,
		size.Height/2 + 100,
	}
	app.render_target.FillRectangle(
		&rectangle1,
		app.light_slate_gray_brush)
	app.render_target.DrawRectangle(
		&rectangle2,
		app.cornflower_blue,
		1,
		nil)
	app.render_target.EndDraw()
	if r := recover(); r != nil {
		app.DiscardDeviceResources()
	}
}

func main() {
	w32.CoInitialize()
	defer w32.CoUninitialize()
	app := new(DemoApp)
	defer app.Dispose()
	app.Initialize()
	app.RunMessageLoop()
}

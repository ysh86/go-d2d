// +build windows

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 06152247-6f50-465a-9245-118bfd3b6007
var IID_ID2D1Factory = GUID{0x06152247, 0x6f50, 0x465a, [8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}

type ID2D1FactoryVtbl struct {
	IUnknownVtbl
	ReloadSystemMetrics            uintptr
	GetDesktopDpi                  uintptr
	CreateRectangleGeometry        uintptr
	CreateRoundedRectangleGeometry uintptr
	CreateEllipseGeometry          uintptr
	CreateGeometryGroup            uintptr
	CreateTransformedGeometry      uintptr
	CreatePathGeometry             uintptr
	CreateStrokeStyle              uintptr
	CreateDrawingStateBlock        uintptr
	CreateWicBitmapRenderTarget    uintptr
	CreateHwndRenderTarget         uintptr
	CreateDxgiSurfaceRenderTarget  uintptr
	CreateDCRenderTarget           uintptr
}

type ID2D1Factory struct {
	IUnknown
}

func (obj *ID2D1Factory) vtbl() *ID2D1FactoryVtbl {
	return (*ID2D1FactoryVtbl)(obj.unsafeVtbl)
}

func (obj *ID2D1Factory) ReloadSystemMetrics() (
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().ReloadSystemMetrics,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call ReloadSystemMetrics: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) GetDesktopDpi() (
	dpiX float32,
	dpiY float32) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetDesktopDpi,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&dpiX)),
		uintptr(unsafe.Pointer(&dpiY)))
	return
}

func (obj *ID2D1Factory) CreateRectangleGeometry(
	rectangle *D2D1_RECT_F) (
	rectangleGeometry *ID2D1RectangleGeometry,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CreateRectangleGeometry,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(rectangle)),
		uintptr(unsafe.Pointer(&rectangleGeometry)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateRectangleGeometry: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateRoundedRectangleGeometry(
	roundedRectangle *D2D1_ROUNDED_RECT) (
	roundedRectangleGeometry *ID2D1RoundedRectangleGeometry,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CreateRoundedRectangleGeometry,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(roundedRectangle)),
		uintptr(unsafe.Pointer(&roundedRectangleGeometry)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateRoundedRectangleGeometry: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateEllipseGeometry(
	ellipse *D2D1_ELLIPSE) (
	ellipseGeometry *ID2D1EllipseGeometry,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CreateEllipseGeometry,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ellipse)),
		uintptr(unsafe.Pointer(&ellipseGeometry)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateEllipseGeometry: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateGeometryGroup(
	fillMode D2D1_FILL_MODE,
	geometries []*ID2D1Geometry) (
	geometryGroup *ID2D1GeometryGroup,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateGeometryGroup,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(fillMode),
		uintptr(unsafe.Pointer(&(geometries[0]))),
		uintptr(len(geometries)),
		uintptr(unsafe.Pointer(&geometryGroup)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateGeometryGroup: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateTransformedGeometry(
	sourceGeometry *ID2D1Geometry,
	transform *D2D1_MATRIX_3X2_F) (
	transformedGeometry *ID2D1TransformedGeometry,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateTransformedGeometry,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(sourceGeometry)),
		uintptr(unsafe.Pointer(transform)),
		uintptr(unsafe.Pointer(&transformedGeometry)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateTransformedGeometry: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreatePathGeometry() (
	pathGeometry *ID2D1PathGeometry,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CreatePathGeometry,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&pathGeometry)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreatePathGeometry: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateStrokeStyle(
	strokeStyleProperties *D2D1_STROKE_STYLE_PROPERTIES,
	dashes []float32) (
	strokeStyle *ID2D1StrokeStyle,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateStrokeStyle,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(strokeStyleProperties)),
		uintptr(unsafe.Pointer(&(dashes[0]))),
		uintptr(len(dashes)),
		uintptr(unsafe.Pointer(&strokeStyle)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateStrokeStyle: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateDrawingStateBlock(
	drawingStateDescription *D2D1_DRAWING_STATE_DESCRIPTION,
	textRenderingParams *IDWriteRenderingParams) (
	drawingStateBlock *ID2D1DrawingStateBlock,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateDrawingStateBlock,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(drawingStateDescription)),
		uintptr(unsafe.Pointer(textRenderingParams)),
		uintptr(unsafe.Pointer(&drawingStateBlock)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateDrawingStateBlock: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateWicBitmapRenderTarget(
	target *IWICBitmap,
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES) (
	renderTarget *ID2D1RenderTarget,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateWicBitmapRenderTarget,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(target)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(&renderTarget)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateWicBitmapRenderTarget: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateHwndRenderTarget(
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES,
	hwndRenderTargetProperties *D2D1_HWND_RENDER_TARGET_PROPERTIES) (
	hwndRenderTarget *ID2D1HwndRenderTarget,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateHwndRenderTarget,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(hwndRenderTargetProperties)),
		uintptr(unsafe.Pointer(&hwndRenderTarget)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateHwndRenderTarget: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateDxgiSurfaceRenderTarget(
	dxgiSurface *IDXGISurface,
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES) (
	renderTarget *ID2D1RenderTarget,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateDxgiSurfaceRenderTarget,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dxgiSurface)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(&renderTarget)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateDxgiSurfaceRenderTarget: %#x", ret)
	}
	return
}

func (obj *ID2D1Factory) CreateDCRenderTarget(
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES) (
	dcRenderTarget *ID2D1DCRenderTarget,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CreateDCRenderTarget,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(&dcRenderTarget)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateDCRenderTarget: %#x", ret)
	}
	return
}

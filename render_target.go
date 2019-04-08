// +build windows

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 2cd90694-12e2-11dc-9fed-001143a055f9
var IID_ID2D1RenderTarget = GUID{0x2cd90694, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1RenderTargetVtbl struct {
	ID2D1ResourceVtbl
	CreateBitmap                 uintptr
	CreateBitmapFromWicBitmap    uintptr
	CreateSharedBitmap           uintptr
	CreateBitmapBrush            uintptr
	CreateSolidColorBrush        uintptr
	CreateGradientStopCollection uintptr
	CreateLinearGradientBrush    uintptr
	CreateRadialGradientBrush    uintptr
	CreateCompatibleRenderTarget uintptr
	CreateLayer                  uintptr
	CreateMesh                   uintptr
	DrawLine                     uintptr
	DrawRectangle                uintptr
	FillRectangle                uintptr
	DrawRoundedRectangle         uintptr
	FillRoundedRectangle         uintptr
	DrawEllipse                  uintptr
	FillEllipse                  uintptr
	DrawGeometry                 uintptr
	FillGeometry                 uintptr
	FillMesh                     uintptr
	FillOpacityMask              uintptr
	DrawBitmap                   uintptr
	DrawText                     uintptr
	DrawTextLayout               uintptr
	DrawGlyphRun                 uintptr
	SetTransform                 uintptr
	GetTransform                 uintptr
	SetAntialiasMode             uintptr
	GetAntialiasMode             uintptr
	SetTextAntialiasMode         uintptr
	GetTextAntialiasMode         uintptr
	SetTextRenderingParams       uintptr
	GetTextRenderingParams       uintptr
	SetTags                      uintptr
	GetTags                      uintptr
	PushLayer                    uintptr
	PopLayer                     uintptr
	Flush                        uintptr
	SaveDrawingState             uintptr
	RestoreDrawingState          uintptr
	PushAxisAlignedClip          uintptr
	PopAxisAlignedClip           uintptr
	Clear                        uintptr
	BeginDraw                    uintptr
	EndDraw                      uintptr
	GetPixelFormat               uintptr
	SetDpi                       uintptr
	GetDpi                       uintptr
	GetSize                      uintptr
	GetPixelSize                 uintptr
	GetMaximumBitmapSize         uintptr
	IsSupported                  uintptr
}

type ID2D1RenderTarget struct {
	ID2D1Resource
}

func (obj *ID2D1RenderTarget) vtbl() *ID2D1RenderTargetVtbl {
	return (*ID2D1RenderTargetVtbl)(obj.unsafeVtbl)
}

func (obj *ID2D1RenderTarget) CreateBitmap(
	size D2D1_SIZE_U,
	srcData *void,
	pitch uint32,
	bitmapProperties *D2D1_BITMAP_PROPERTIES) (
	bitmap *ID2D1Bitmap,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateBitmap,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(size),
		uintptr(unsafe.Pointer(srcData)),
		uintptr(pitch),
		uintptr(unsafe.Pointer(bitmapProperties)),
		uintptr(unsafe.Pointer(&bitmap)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateBitmap: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateBitmapFromWicBitmap(
	wicBitmapSource *IWICBitmapSource,
	bitmapProperties *D2D1_BITMAP_PROPERTIES) (
	bitmap *ID2D1Bitmap,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateBitmapFromWicBitmap,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(wicBitmapSource)),
		uintptr(unsafe.Pointer(bitmapProperties)),
		uintptr(unsafe.Pointer(&bitmap)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateBitmapFromWicBitmap: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateSharedBitmap(
	riid REFIID,
	data *void,
	bitmapProperties *D2D1_BITMAP_PROPERTIES) (
	bitmap *ID2D1Bitmap,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateSharedBitmap,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(riid),
		uintptr(unsafe.Pointer(data)),
		uintptr(unsafe.Pointer(bitmapProperties)),
		uintptr(unsafe.Pointer(&bitmap)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateSharedBitmap: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateBitmapBrush(
	bitmap *ID2D1Bitmap,
	bitmapBrushProperties *D2D1_BITMAP_BRUSH_PROPERTIES,
	brushProperties *D2D1_BRUSH_PROPERTIES) (
	bitmapBrush *ID2D1BitmapBrush,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateBitmapBrush,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(bitmapBrushProperties)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(&bitmapBrush)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateBitmapBrush: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateSolidColorBrush(
	color *D2D1_COLOR_F,
	brushProperties *D2D1_BRUSH_PROPERTIES) (
	solidColorBrush *ID2D1SolidColorBrush,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateSolidColorBrush,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(color)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(&solidColorBrush)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateSolidColorBrush: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateGradientStopCollection(
	gradientStops []D2D1_GRADIENT_STOP,
	colorInterpolationGamma D2D1_GAMMA,
	extendMode D2D1_EXTEND_MODE) (
	gradientStopCollection *ID2D1GradientStopCollection,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateGradientStopCollection,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(gradientStops[0]))),
		uintptr(len(gradientStops)),
		uintptr(colorInterpolationGamma),
		uintptr(extendMode),
		uintptr(unsafe.Pointer(&gradientStopCollection)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateGradientStopCollection: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateLinearGradientBrush(
	linearGradientBrushProperties *D2D1_LINEAR_GRADIENT_BRUSH_PROPERTIES,
	brushProperties *D2D1_BRUSH_PROPERTIES,
	gradientStopCollection *ID2D1GradientStopCollection) (
	linearGradientBrush *ID2D1LinearGradientBrush,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateLinearGradientBrush,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(linearGradientBrushProperties)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(gradientStopCollection)),
		uintptr(unsafe.Pointer(&linearGradientBrush)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateLinearGradientBrush: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateRadialGradientBrush(
	radialGradientBrushProperties *D2D1_RADIAL_GRADIENT_BRUSH_PROPERTIES,
	brushProperties *D2D1_BRUSH_PROPERTIES,
	gradientStopCollection *ID2D1GradientStopCollection) (
	radialGradientBrush *ID2D1RadialGradientBrush,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateRadialGradientBrush,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(radialGradientBrushProperties)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(gradientStopCollection)),
		uintptr(unsafe.Pointer(&radialGradientBrush)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateRadialGradientBrush: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateCompatibleRenderTarget(
	desiredSize *D2D1_SIZE_F,
	desiredPixelSize *D2D1_SIZE_U,
	desiredFormat *D2D1_PIXEL_FORMAT,
	options D2D1_COMPATIBLE_RENDER_TARGET_OPTIONS) (
	bitmapRenderTarget *ID2D1BitmapRenderTarget,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CreateCompatibleRenderTarget,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(desiredSize)),
		uintptr(unsafe.Pointer(desiredPixelSize)),
		uintptr(unsafe.Pointer(desiredFormat)),
		uintptr(options),
		uintptr(unsafe.Pointer(&bitmapRenderTarget)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateCompatibleRenderTarget: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateLayer(
	size *D2D1_SIZE_F) (
	layer *ID2D1Layer,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CreateLayer,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(size)),
		uintptr(unsafe.Pointer(&layer)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateLayer: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) CreateMesh() (
	mesh *ID2D1Mesh,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CreateMesh,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&mesh)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CreateMesh: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) DrawLine(
	point0 D2D1_POINT_2F,
	point1 D2D1_POINT_2F,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawLine,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(point0),
		uintptr(point1),
		uintptr(unsafe.Pointer(brush)),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)))
	return
}

func (obj *ID2D1RenderTarget) DrawRectangle(
	rect *D2D1_RECT_F,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawRectangle,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)
	return
}

func (obj *ID2D1RenderTarget) FillRectangle(
	rect *D2D1_RECT_F,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().FillRectangle,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(brush)))
	return
}

func (obj *ID2D1RenderTarget) DrawRoundedRectangle(
	roundedRect *D2D1_ROUNDED_RECT,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawRoundedRectangle,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(roundedRect)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)
	return
}

func (obj *ID2D1RenderTarget) FillRoundedRectangle(
	roundedRect *D2D1_ROUNDED_RECT,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().FillRoundedRectangle,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(roundedRect)),
		uintptr(unsafe.Pointer(brush)))
	return
}

func (obj *ID2D1RenderTarget) DrawEllipse(
	ellipse *D2D1_ELLIPSE,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawEllipse,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ellipse)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)
	return
}

func (obj *ID2D1RenderTarget) FillEllipse(
	ellipse *D2D1_ELLIPSE,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().FillEllipse,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ellipse)),
		uintptr(unsafe.Pointer(brush)))
	return
}

func (obj *ID2D1RenderTarget) DrawGeometry(
	geometry *ID2D1Geometry,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawGeometry,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(geometry)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)
	return
}

func (obj *ID2D1RenderTarget) FillGeometry(
	geometry *ID2D1Geometry,
	brush *ID2D1Brush,
	opacityBrush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().FillGeometry,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(geometry)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(opacityBrush)),
		0,
		0)
	return
}

func (obj *ID2D1RenderTarget) FillMesh(
	mesh *ID2D1Mesh,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().FillMesh,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(mesh)),
		uintptr(unsafe.Pointer(brush)))
	return
}

func (obj *ID2D1RenderTarget) FillOpacityMask(
	opacityMask *ID2D1Bitmap,
	brush *ID2D1Brush,
	content D2D1_OPACITY_MASK_CONTENT,
	destinationRectangle *D2D1_RECT_F,
	sourceRectangle *D2D1_RECT_F) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().FillOpacityMask,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(opacityMask)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(content),
		uintptr(unsafe.Pointer(destinationRectangle)),
		uintptr(unsafe.Pointer(sourceRectangle)))
	return
}

func (obj *ID2D1RenderTarget) DrawBitmap(
	bitmap *ID2D1Bitmap,
	destinationRectangle *D2D1_RECT_F,
	opacity float32,
	interpolationMode D2D1_BITMAP_INTERPOLATION_MODE,
	sourceRectangle *D2D1_RECT_F) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawBitmap,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(destinationRectangle)),
		uintptr(*(*uint32)(unsafe.Pointer(&opacity))),
		uintptr(interpolationMode),
		uintptr(unsafe.Pointer(sourceRectangle)))
	return
}

func (obj *ID2D1RenderTarget) DrawText(
	string []WCHAR,
	textFormat *IDWriteTextFormat,
	layoutRect *D2D1_RECT_F,
	defaultForegroundBrush *ID2D1Brush,
	options D2D1_DRAW_TEXT_OPTIONS,
	measuringMode DWRITE_MEASURING_MODE) {
	var _, _, _ = syscall.Syscall9(
		obj.vtbl().DrawText,
		8,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(string[0]))),
		uintptr(len(string)),
		uintptr(unsafe.Pointer(textFormat)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(defaultForegroundBrush)),
		uintptr(options),
		uintptr(measuringMode),
		0)
	return
}

func (obj *ID2D1RenderTarget) DrawTextLayout(
	origin D2D1_POINT_2F,
	textLayout *IDWriteTextLayout,
	defaultForegroundBrush *ID2D1Brush,
	options D2D1_DRAW_TEXT_OPTIONS) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawTextLayout,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(origin),
		uintptr(unsafe.Pointer(textLayout)),
		uintptr(unsafe.Pointer(defaultForegroundBrush)),
		uintptr(options),
		0)
	return
}

func (obj *ID2D1RenderTarget) DrawGlyphRun(
	baselineOrigin D2D1_POINT_2F,
	glyphRun *DWRITE_GLYPH_RUN,
	foregroundBrush *ID2D1Brush,
	measuringMode DWRITE_MEASURING_MODE) {
	var _, _, _ = syscall.Syscall6(
		obj.vtbl().DrawGlyphRun,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(baselineOrigin),
		uintptr(unsafe.Pointer(glyphRun)),
		uintptr(unsafe.Pointer(foregroundBrush)),
		uintptr(measuringMode),
		0)
	return
}

func (obj *ID2D1RenderTarget) SetTransform(
	transform *D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetTransform,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(transform)),
		0)
	return
}

func (obj *ID2D1RenderTarget) GetTransform() (
	transform D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetTransform,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&transform)),
		0)
	return
}

func (obj *ID2D1RenderTarget) SetAntialiasMode(
	antialiasMode D2D1_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetAntialiasMode,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(antialiasMode),
		0)
	return
}

func (obj *ID2D1RenderTarget) GetAntialiasMode() (
	result D2D1_ANTIALIAS_MODE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetAntialiasMode,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_ANTIALIAS_MODE)(ret)
	return
}

func (obj *ID2D1RenderTarget) SetTextAntialiasMode(
	textAntialiasMode D2D1_TEXT_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetTextAntialiasMode,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(textAntialiasMode),
		0)
	return
}

func (obj *ID2D1RenderTarget) GetTextAntialiasMode() (
	result D2D1_TEXT_ANTIALIAS_MODE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetTextAntialiasMode,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_TEXT_ANTIALIAS_MODE)(ret)
	return
}

func (obj *ID2D1RenderTarget) SetTextRenderingParams(
	textRenderingParams *IDWriteRenderingParams) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetTextRenderingParams,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(textRenderingParams)),
		0)
	return
}

func (obj *ID2D1RenderTarget) GetTextRenderingParams() (
	textRenderingParams *IDWriteRenderingParams) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetTextRenderingParams,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&textRenderingParams)),
		0)
	return
}

func (obj *ID2D1RenderTarget) SetTags(
	tag1 D2D1_TAG,
	tag2 D2D1_TAG) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetTags,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(tag1),
		uintptr(tag2))
	return
}

func (obj *ID2D1RenderTarget) GetTags() (
	tag1 D2D1_TAG,
	tag2 D2D1_TAG) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetTags,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&tag1)),
		uintptr(unsafe.Pointer(&tag2)))
	return
}

func (obj *ID2D1RenderTarget) PushLayer(
	layerParameters *D2D1_LAYER_PARAMETERS,
	layer *ID2D1Layer) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().PushLayer,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(layerParameters)),
		uintptr(unsafe.Pointer(layer)))
	return
}

func (obj *ID2D1RenderTarget) PopLayer() {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().PopLayer,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return
}

func (obj *ID2D1RenderTarget) Flush() (
	tag1 D2D1_TAG,
	tag2 D2D1_TAG,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().Flush,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&tag1)),
		uintptr(unsafe.Pointer(&tag2)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Flush: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) SaveDrawingState(
	drawingStateBlock *ID2D1DrawingStateBlock) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SaveDrawingState,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(drawingStateBlock)),
		0)
	return
}

func (obj *ID2D1RenderTarget) RestoreDrawingState(
	drawingStateBlock *ID2D1DrawingStateBlock) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().RestoreDrawingState,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(drawingStateBlock)),
		0)
	return
}

func (obj *ID2D1RenderTarget) PushAxisAlignedClip(
	clipRect *D2D1_RECT_F,
	antialiasMode D2D1_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().PushAxisAlignedClip,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(clipRect)),
		uintptr(antialiasMode))
	return
}

func (obj *ID2D1RenderTarget) PopAxisAlignedClip() {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().PopAxisAlignedClip,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return
}

func (obj *ID2D1RenderTarget) Clear(
	clearColor *D2D1_COLOR_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().Clear,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(clearColor)),
		0)
	return
}

func (obj *ID2D1RenderTarget) BeginDraw() {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().BeginDraw,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return
}

func (obj *ID2D1RenderTarget) EndDraw() (
	tag1 D2D1_TAG,
	tag2 D2D1_TAG,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().EndDraw,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&tag1)),
		uintptr(unsafe.Pointer(&tag2)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call EndDraw: %#x", ret)
	}
	return
}

func (obj *ID2D1RenderTarget) GetPixelFormat() (
	result D2D1_PIXEL_FORMAT) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetPixelFormat,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_PIXEL_FORMAT)(ret)
	return
}

func (obj *ID2D1RenderTarget) SetDpi(
	dpiX float32,
	dpiY float32) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetDpi,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint32)(unsafe.Pointer(&dpiX))),
		uintptr(*(*uint32)(unsafe.Pointer(&dpiY))))
	return
}

func (obj *ID2D1RenderTarget) GetDpi() (
	dpiX float32,
	dpiY float32) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetDpi,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&dpiX)),
		uintptr(unsafe.Pointer(&dpiY)))
	return
}

func (obj *ID2D1RenderTarget) GetSize() (
	result D2D1_SIZE_F) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetSize,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_SIZE_F)(ret)
	return
}

func (obj *ID2D1RenderTarget) GetPixelSize() (
	result D2D1_SIZE_U) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetPixelSize,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_SIZE_U)(ret)
	return
}

func (obj *ID2D1RenderTarget) GetMaximumBitmapSize() (
	result uint32) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetMaximumBitmapSize,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (uint32)(ret)
	return
}

func (obj *ID2D1RenderTarget) IsSupported(
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES) (
	result BOOL) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().IsSupported,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		0)
	result = (BOOL)(ret)
	return
}

// 2cd90698-12e2-11dc-9fed-001143a055f9
var IID_ID2D1HwndRenderTarget = GUID{0x2cd90698, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1HwndRenderTargetVtbl struct {
	ID2D1RenderTargetVtbl
	CheckWindowState uintptr
	Resize           uintptr
	GetHwnd          uintptr
}

type ID2D1HwndRenderTarget struct {
	ID2D1RenderTarget
}

func (obj *ID2D1HwndRenderTarget) vtbl() *ID2D1HwndRenderTargetVtbl {
	return (*ID2D1HwndRenderTargetVtbl)(obj.unsafeVtbl)
}

func (obj *ID2D1HwndRenderTarget) CheckWindowState() (
	result D2D1_WINDOW_STATE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().CheckWindowState,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_WINDOW_STATE)(ret)
	return
}

func (obj *ID2D1HwndRenderTarget) Resize(
	pixelSize *D2D1_SIZE_U) (
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().Resize,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pixelSize)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Resize: %#x", ret)
	}
	return
}

func (obj *ID2D1HwndRenderTarget) GetHwnd() (
	result HWND) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetHwnd,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (HWND)(ret)
	return
}

// 1c51bc64-de61-46fd-9899-63a5d8f03950
var IID_ID2D1DCRenderTarget = GUID{0x1c51bc64, 0xde61, 0x46fd, [8]byte{0x98, 0x99, 0x63, 0xa5, 0xd8, 0xf0, 0x39, 0x50}}

type ID2D1DCRenderTargetVtbl struct {
	ID2D1RenderTargetVtbl
	BindDC uintptr
}

type ID2D1DCRenderTarget struct {
	ID2D1RenderTarget
}

func (obj *ID2D1DCRenderTarget) vtbl() *ID2D1DCRenderTargetVtbl {
	return (*ID2D1DCRenderTargetVtbl)(obj.unsafeVtbl)
}

func (obj *ID2D1DCRenderTarget) BindDC(
	hDC HDC,
	pSubRect *RECT) (
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().BindDC,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hDC),
		uintptr(unsafe.Pointer(pSubRect)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call BindDC: %#x", ret)
	}
	return
}

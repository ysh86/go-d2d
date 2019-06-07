// +build windows

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 2cd906a1-12e2-11dc-9fed-001143a055f9
var IID_ID2D1Geometry = GUID{0x2cd906a1, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1Geometry struct {
	ID2D1Resource
}

type vtblID2D1Geometry struct {
	vtblID2D1Resource
	GetBounds            uintptr
	GetWidenedBounds     uintptr
	StrokeContainsPoint  uintptr
	FillContainsPoint    uintptr
	CompareWithGeometry  uintptr
	Simplify             uintptr
	Tessellate           uintptr
	CombineWithGeometry  uintptr
	Outline              uintptr
	ComputeArea          uintptr
	ComputeLength        uintptr
	ComputePointAtLength uintptr
	Widen                uintptr
}

func (obj *ID2D1Geometry) vtbl() *vtblID2D1Geometry {
	return (*vtblID2D1Geometry)(obj.unsafeVtbl)
}

func (obj *ID2D1Geometry) GetBounds(
	worldTransform *D2D1_MATRIX_3X2_F) (
	bounds D2D1_RECT_F,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetBounds,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(unsafe.Pointer(&bounds)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call GetBounds: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) GetWidenedBounds(
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	bounds D2D1_RECT_F,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().GetWidenedBounds,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&bounds)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call GetWidenedBounds: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) StrokeContainsPoint(
	point D2D1_POINT_2F,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	contains bool,
	err error) {
	var containsWinbool int32
	var ret, _, _ = syscall.Syscall9(
		obj.vtbl().StrokeContainsPoint,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint64)(unsafe.Pointer(&point))),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&containsWinbool)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call StrokeContainsPoint: %#x", ret)
	}
	contains = (containsWinbool != 0)
	return
}

func (obj *ID2D1Geometry) FillContainsPoint(
	point D2D1_POINT_2F,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	contains bool,
	err error) {
	var containsWinbool int32
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().FillContainsPoint,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint64)(unsafe.Pointer(&point))),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&containsWinbool)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call FillContainsPoint: %#x", ret)
	}
	contains = (containsWinbool != 0)
	return
}

func (obj *ID2D1Geometry) CompareWithGeometry(
	inputGeometry *ID2D1Geometry,
	inputGeometryTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	relation D2D1_GEOMETRY_RELATION,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CompareWithGeometry,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(inputGeometry)),
		uintptr(unsafe.Pointer(inputGeometryTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&relation)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CompareWithGeometry: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) Simplify(
	simplificationOption D2D1_GEOMETRY_SIMPLIFICATION_OPTION,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink) (
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().Simplify,
		5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(simplificationOption),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(geometrySink)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Simplify: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) Tessellate(
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	tessellationSink *ID2D1TessellationSink) (
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().Tessellate,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(tessellationSink)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Tessellate: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) CombineWithGeometry(
	inputGeometry *ID2D1Geometry,
	combineMode D2D1_COMBINE_MODE,
	inputGeometryTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink) (
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().CombineWithGeometry,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(inputGeometry)),
		uintptr(combineMode),
		uintptr(unsafe.Pointer(inputGeometryTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(geometrySink)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call CombineWithGeometry: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) Outline(
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink) (
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().Outline,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(geometrySink)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Outline: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) ComputeArea(
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	area float32,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().ComputeArea,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&area)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call ComputeArea: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) ComputeLength(
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	length float32,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().ComputeLength,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&length)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call ComputeLength: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) ComputePointAtLength(
	length float32,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	point D2D1_POINT_2F,
	unitTangentVector D2D1_POINT_2F,
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().ComputePointAtLength,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint32)(unsafe.Pointer(&length))),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&point)),
		uintptr(unsafe.Pointer(&unitTangentVector)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call ComputePointAtLength: %#x", ret)
	}
	return
}

func (obj *ID2D1Geometry) Widen(
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink) (
	err error) {
	var ret, _, _ = syscall.Syscall6(
		obj.vtbl().Widen,
		6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(geometrySink)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Widen: %#x", ret)
	}
	return
}

// 2cd906a2-12e2-11dc-9fed-001143a055f9
var IID_ID2D1RectangleGeometry = GUID{0x2cd906a2, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1RectangleGeometry struct {
	ID2D1Geometry
}

type vtblID2D1RectangleGeometry struct {
	vtblID2D1Geometry
	GetRect uintptr
}

func (obj *ID2D1RectangleGeometry) vtbl() *vtblID2D1RectangleGeometry {
	return (*vtblID2D1RectangleGeometry)(obj.unsafeVtbl)
}

func (obj *ID2D1RectangleGeometry) GetRect() (
	rect D2D1_RECT_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetRect,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&rect)),
		0)
	return
}

// 2cd906a3-12e2-11dc-9fed-001143a055f9
var IID_ID2D1RoundedRectangleGeometry = GUID{0x2cd906a3, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1RoundedRectangleGeometry struct {
	ID2D1Geometry
}

type vtblID2D1RoundedRectangleGeometry struct {
	vtblID2D1Geometry
	GetRoundedRect uintptr
}

func (obj *ID2D1RoundedRectangleGeometry) vtbl() *vtblID2D1RoundedRectangleGeometry {
	return (*vtblID2D1RoundedRectangleGeometry)(obj.unsafeVtbl)
}

func (obj *ID2D1RoundedRectangleGeometry) GetRoundedRect() (
	roundedRect D2D1_ROUNDED_RECT) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetRoundedRect,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&roundedRect)),
		0)
	return
}

// 2cd906a4-12e2-11dc-9fed-001143a055f9
var IID_ID2D1EllipseGeometry = GUID{0x2cd906a4, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1EllipseGeometry struct {
	ID2D1Geometry
}

type vtblID2D1EllipseGeometry struct {
	vtblID2D1Geometry
	GetEllipse uintptr
}

func (obj *ID2D1EllipseGeometry) vtbl() *vtblID2D1EllipseGeometry {
	return (*vtblID2D1EllipseGeometry)(obj.unsafeVtbl)
}

func (obj *ID2D1EllipseGeometry) GetEllipse() (
	ellipse D2D1_ELLIPSE) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetEllipse,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&ellipse)),
		0)
	return
}

// 2cd906a6-12e2-11dc-9fed-001143a055f9
var IID_ID2D1GeometryGroup = GUID{0x2cd906a6, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1GeometryGroup struct {
	ID2D1Geometry
}

type vtblID2D1GeometryGroup struct {
	vtblID2D1Geometry
	GetFillMode            uintptr
	GetSourceGeometryCount uintptr
	GetSourceGeometries    uintptr
}

func (obj *ID2D1GeometryGroup) vtbl() *vtblID2D1GeometryGroup {
	return (*vtblID2D1GeometryGroup)(obj.unsafeVtbl)
}

func (obj *ID2D1GeometryGroup) GetFillMode() (
	result D2D1_FILL_MODE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetFillMode,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_FILL_MODE)(ret)
	return
}

func (obj *ID2D1GeometryGroup) GetSourceGeometryCount() (
	result uint32) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetSourceGeometryCount,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (uint32)(ret)
	return
}

func (obj *ID2D1GeometryGroup) GetSourceGeometries(
	geometries []*ID2D1Geometry) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetSourceGeometries,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(geometries[0]))),
		uintptr(len(geometries)))
	return
}

// 2cd906bb-12e2-11dc-9fed-001143a055f9
var IID_ID2D1TransformedGeometry = GUID{0x2cd906bb, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1TransformedGeometry struct {
	ID2D1Geometry
}

type vtblID2D1TransformedGeometry struct {
	vtblID2D1Geometry
	GetSourceGeometry uintptr
	GetTransform      uintptr
}

func (obj *ID2D1TransformedGeometry) vtbl() *vtblID2D1TransformedGeometry {
	return (*vtblID2D1TransformedGeometry)(obj.unsafeVtbl)
}

func (obj *ID2D1TransformedGeometry) GetSourceGeometry() (
	sourceGeometry *ID2D1Geometry) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetSourceGeometry,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&sourceGeometry)),
		0)
	return
}

func (obj *ID2D1TransformedGeometry) GetTransform() (
	transform D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetTransform,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&transform)),
		0)
	return
}

// 2cd906a5-12e2-11dc-9fed-001143a055f9
var IID_ID2D1PathGeometry = GUID{0x2cd906a5, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1PathGeometry struct {
	ID2D1Geometry
}

type vtblID2D1PathGeometry struct {
	vtblID2D1Geometry
	Open            uintptr
	Stream          uintptr
	GetSegmentCount uintptr
	GetFigureCount  uintptr
}

func (obj *ID2D1PathGeometry) vtbl() *vtblID2D1PathGeometry {
	return (*vtblID2D1PathGeometry)(obj.unsafeVtbl)
}

func (obj *ID2D1PathGeometry) Open() (
	geometrySink *ID2D1GeometrySink,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().Open,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&geometrySink)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Open: %#x", ret)
	}
	return
}

func (obj *ID2D1PathGeometry) Stream(
	geometrySink *ID2D1GeometrySink) (
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().Stream,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(geometrySink)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Stream: %#x", ret)
	}
	return
}

func (obj *ID2D1PathGeometry) GetSegmentCount() (
	count uint32,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetSegmentCount,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&count)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call GetSegmentCount: %#x", ret)
	}
	return
}

func (obj *ID2D1PathGeometry) GetFigureCount() (
	count uint32,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetFigureCount,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&count)),
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call GetFigureCount: %#x", ret)
	}
	return
}

// 2cd9069e-12e2-11dc-9fed-001143a055f9
var IID_ID2D1SimplifiedGeometrySink = GUID{0x2cd9069e, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1SimplifiedGeometrySink struct {
	IUnknown
}

type vtblID2D1SimplifiedGeometrySink struct {
	vtblIUnknown
	SetFillMode     uintptr
	SetSegmentFlags uintptr
	BeginFigure     uintptr
	AddLines        uintptr
	AddBeziers      uintptr
	EndFigure       uintptr
	Close           uintptr
}

func (obj *ID2D1SimplifiedGeometrySink) vtbl() *vtblID2D1SimplifiedGeometrySink {
	return (*vtblID2D1SimplifiedGeometrySink)(obj.unsafeVtbl)
}

func (obj *ID2D1SimplifiedGeometrySink) SetFillMode(
	fillMode D2D1_FILL_MODE) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetFillMode,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(fillMode),
		0)
	return
}

func (obj *ID2D1SimplifiedGeometrySink) SetSegmentFlags(
	vertexFlags D2D1_PATH_SEGMENT) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetSegmentFlags,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(vertexFlags),
		0)
	return
}

func (obj *ID2D1SimplifiedGeometrySink) BeginFigure(
	startPoint D2D1_POINT_2F,
	figureBegin D2D1_FIGURE_BEGIN) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().BeginFigure,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint64)(unsafe.Pointer(&startPoint))),
		uintptr(figureBegin))
	return
}

func (obj *ID2D1SimplifiedGeometrySink) AddLines(
	points []D2D1_POINT_2F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddLines,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(points[0]))),
		uintptr(len(points)))
	return
}

func (obj *ID2D1SimplifiedGeometrySink) AddBeziers(
	beziers []D2D1_BEZIER_SEGMENT) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddBeziers,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(beziers[0]))),
		uintptr(len(beziers)))
	return
}

func (obj *ID2D1SimplifiedGeometrySink) EndFigure(
	figureEnd D2D1_FIGURE_END) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().EndFigure,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(figureEnd),
		0)
	return
}

func (obj *ID2D1SimplifiedGeometrySink) Close() (
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().Close,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Close: %#x", ret)
	}
	return
}

// 2cd906c1-12e2-11dc-9fed-001143a055f9
var IID_ID2D1TessellationSink = GUID{0x2cd906c1, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1TessellationSink struct {
	IUnknown
}

type vtblID2D1TessellationSink struct {
	vtblIUnknown
	AddTriangles uintptr
	Close        uintptr
}

func (obj *ID2D1TessellationSink) vtbl() *vtblID2D1TessellationSink {
	return (*vtblID2D1TessellationSink)(obj.unsafeVtbl)
}

func (obj *ID2D1TessellationSink) AddTriangles(
	triangles []D2D1_TRIANGLE) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddTriangles,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(triangles[0]))),
		uintptr(len(triangles)))
	return
}

func (obj *ID2D1TessellationSink) Close() (
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().Close,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	if ret != S_OK {
		err = fmt.Errorf("Fail to call Close: %#x", ret)
	}
	return
}

// 2cd9069f-12e2-11dc-9fed-001143a055f9
var IID_ID2D1GeometrySink = GUID{0x2cd9069f, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1GeometrySink struct {
	ID2D1SimplifiedGeometrySink
}

type vtblID2D1GeometrySink struct {
	vtblID2D1SimplifiedGeometrySink
	AddLine             uintptr
	AddBezier           uintptr
	AddQuadraticBezier  uintptr
	AddQuadraticBeziers uintptr
	AddArc              uintptr
}

func (obj *ID2D1GeometrySink) vtbl() *vtblID2D1GeometrySink {
	return (*vtblID2D1GeometrySink)(obj.unsafeVtbl)
}

func (obj *ID2D1GeometrySink) AddLine(
	point D2D1_POINT_2F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddLine,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint64)(unsafe.Pointer(&point))),
		0)
	return
}

func (obj *ID2D1GeometrySink) AddBezier(
	bezier *D2D1_BEZIER_SEGMENT) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddBezier,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(bezier)),
		0)
	return
}

func (obj *ID2D1GeometrySink) AddQuadraticBezier(
	bezier *D2D1_QUADRATIC_BEZIER_SEGMENT) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddQuadraticBezier,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(bezier)),
		0)
	return
}

func (obj *ID2D1GeometrySink) AddQuadraticBeziers(
	beziers []D2D1_QUADRATIC_BEZIER_SEGMENT) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddQuadraticBeziers,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(beziers[0]))),
		uintptr(len(beziers)))
	return
}

func (obj *ID2D1GeometrySink) AddArc(
	arc *D2D1_ARC_SEGMENT) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().AddArc,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(arc)),
		0)
	return
}

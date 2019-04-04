// +build windows

package d2d

import (
	"syscall"
	"unsafe"
)

// 2cd9069d-12e2-11dc-9fed-001143a055f9
var IID_ID2D1StrokeStyle = GUID{0x2cd9069d, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1StrokeStyleVtbl struct {
	ID2D1ResourceVtbl
	GetStartCap    uintptr
	GetEndCap      uintptr
	GetDashCap     uintptr
	GetMiterLimit  uintptr
	GetLineJoin    uintptr
	GetDashOffset  uintptr
	GetDashStyle   uintptr
	GetDashesCount uintptr
	GetDashes      uintptr
}

type ID2D1StrokeStyle struct {
	vtbl *ID2D1StrokeStyleVtbl
}

func (obj *ID2D1StrokeStyle) GUID() *GUID {
	return &IID_ID2D1StrokeStyle
}

func (obj *ID2D1StrokeStyle) QueryInterface(
	riid *GUID) (
	dest unsafe.Pointer,
	err error) {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid)
}

func (obj *ID2D1StrokeStyle) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *ID2D1StrokeStyle) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}

func (obj *ID2D1StrokeStyle) Parent() *ID2D1Resource {
	return (*ID2D1Resource)(unsafe.Pointer(obj))
}

func (obj *ID2D1StrokeStyle) GetStartCap() (
	result D2D1_CAP_STYLE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetStartCap,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_CAP_STYLE)(ret)
	return
}

func (obj *ID2D1StrokeStyle) GetEndCap() (
	result D2D1_CAP_STYLE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetEndCap,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_CAP_STYLE)(ret)
	return
}

func (obj *ID2D1StrokeStyle) GetDashCap() (
	result D2D1_CAP_STYLE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetDashCap,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_CAP_STYLE)(ret)
	return
}

func (obj *ID2D1StrokeStyle) GetMiterLimit() (
	result float32) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetMiterLimit,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = *(*float32)(unsafe.Pointer(&ret))
	return
}

func (obj *ID2D1StrokeStyle) GetLineJoin() (
	result D2D1_LINE_JOIN) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetLineJoin,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_LINE_JOIN)(ret)
	return
}

func (obj *ID2D1StrokeStyle) GetDashOffset() (
	result float32) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetDashOffset,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = *(*float32)(unsafe.Pointer(&ret))
	return
}

func (obj *ID2D1StrokeStyle) GetDashStyle() (
	result D2D1_DASH_STYLE) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetDashStyle,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_DASH_STYLE)(ret)
	return
}

func (obj *ID2D1StrokeStyle) GetDashesCount() (
	result uint32) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl.GetDashesCount,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (uint32)(ret)
	return
}

func (obj *ID2D1StrokeStyle) GetDashes(
	dashes []float32) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl.GetDashes,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&(dashes[0]))),
		uintptr(len(dashes)))
	return
}

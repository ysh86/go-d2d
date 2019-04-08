// +build windows

package d2d

import (
	"syscall"
	"unsafe"
)

// 2cd906a8-12e2-11dc-9fed-001143a055f9
var IID_ID2D1Brush = GUID{0x2cd906a8, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1BrushVtbl struct {
	ID2D1ResourceVtbl
	SetOpacity   uintptr
	SetTransform uintptr
	GetOpacity   uintptr
	GetTransform uintptr
}

type ID2D1Brush struct {
	ID2D1Resource
}

func (obj *ID2D1Brush) vtbl() *ID2D1BrushVtbl {
	return (*ID2D1BrushVtbl)(obj.unsafeVtbl)
}

func (obj *ID2D1Brush) SetOpacity(
	opacity float32) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetOpacity,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(*(*uint32)(unsafe.Pointer(&opacity))),
		0)
	return
}

func (obj *ID2D1Brush) SetTransform(
	transform *D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetTransform,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(transform)),
		0)
	return
}

func (obj *ID2D1Brush) GetOpacity() (
	result float32) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetOpacity,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	ret32 := uint32(ret)
	result = *(*float32)(unsafe.Pointer(&ret32))
	return
}

func (obj *ID2D1Brush) GetTransform() (
	transform D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetTransform,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&transform)),
		0)
	return
}

// 2cd906a9-12e2-11dc-9fed-001143a055f9
var IID_ID2D1SolidColorBrush = GUID{0x2cd906a9, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1SolidColorBrushVtbl struct {
	ID2D1BrushVtbl
	SetColor uintptr
	GetColor uintptr
}

type ID2D1SolidColorBrush struct {
	ID2D1Brush
}

func (obj *ID2D1SolidColorBrush) vtbl() *ID2D1SolidColorBrushVtbl {
	return (*ID2D1SolidColorBrushVtbl)(obj.unsafeVtbl)
}

func (obj *ID2D1SolidColorBrush) SetColor(
	color *D2D1_COLOR_F) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetColor,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(color)),
		0)
	return
}

func (obj *ID2D1SolidColorBrush) GetColor() (
	result D2D1_COLOR_F) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().GetColor,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	result = (D2D1_COLOR_F)(ret)
	return
}

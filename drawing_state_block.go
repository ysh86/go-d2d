// +build windows

package d2d

import (
	"syscall"
	"unsafe"
)

// 28506e39-ebf6-46a1-bb47-fd85565ab957
var IID_ID2D1DrawingStateBlock = GUID{0x28506e39, 0xebf6, 0x46a1, [8]byte{0xbb, 0x47, 0xfd, 0x85, 0x56, 0x5a, 0xb9, 0x57}}

type ID2D1DrawingStateBlockVtbl struct {
	ID2D1ResourceVtbl
	GetDescription         uintptr
	SetDescription         uintptr
	SetTextRenderingParams uintptr
	GetTextRenderingParams uintptr
}

type ID2D1DrawingStateBlock struct {
	ID2D1Resource
}

func (obj *ID2D1DrawingStateBlock) vtbl() *ID2D1DrawingStateBlockVtbl {
	return (*ID2D1DrawingStateBlockVtbl)(obj.unsafeVtbl)
}

func (obj *ID2D1DrawingStateBlock) GetDescription() (
	stateDescription D2D1_DRAWING_STATE_DESCRIPTION) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetDescription,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&stateDescription)),
		0)
	return
}

func (obj *ID2D1DrawingStateBlock) SetDescription(
	stateDescription *D2D1_DRAWING_STATE_DESCRIPTION) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetDescription,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(stateDescription)),
		0)
	return
}

func (obj *ID2D1DrawingStateBlock) SetTextRenderingParams(
	textRenderingParams *IDWriteRenderingParams) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().SetTextRenderingParams,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(textRenderingParams)),
		0)
	return
}

func (obj *ID2D1DrawingStateBlock) GetTextRenderingParams() (
	textRenderingParams *IDWriteRenderingParams) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetTextRenderingParams,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&textRenderingParams)),
		0)
	return
}

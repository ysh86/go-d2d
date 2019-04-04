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
	vtbl *ID2D1DrawingStateBlockVtbl
}

func (obj *ID2D1DrawingStateBlock) GUID() *GUID {
	return &IID_ID2D1DrawingStateBlock
}

func (obj *ID2D1DrawingStateBlock) QueryInterface(
	riid *GUID) (
	dest unsafe.Pointer,
	err error) {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid)
}

func (obj *ID2D1DrawingStateBlock) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *ID2D1DrawingStateBlock) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}

func (obj *ID2D1DrawingStateBlock) Parent() *ID2D1Resource {
	return (*ID2D1Resource)(unsafe.Pointer(obj))
}

func (obj *ID2D1DrawingStateBlock) GetDescription() (
	stateDescription D2D1_DRAWING_STATE_DESCRIPTION) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl.GetDescription,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&stateDescription)),
		0)
	return
}

func (obj *ID2D1DrawingStateBlock) SetDescription(
	stateDescription *D2D1_DRAWING_STATE_DESCRIPTION) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl.SetDescription,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(stateDescription)),
		0)
	return
}

func (obj *ID2D1DrawingStateBlock) SetTextRenderingParams(
	//textRenderingParams *IDWriteRenderingParams) {
	textRenderingParams unsafe.Pointer) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl.SetTextRenderingParams,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(textRenderingParams)),
		0)
	return
}

func (obj *ID2D1DrawingStateBlock) GetTextRenderingParams() (
	//textRenderingParams *IDWriteRenderingParams) {
	textRenderingParams unsafe.Pointer) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl.GetTextRenderingParams,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&textRenderingParams)),
		0)
	return
}

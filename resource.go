// +build windows

package d2d

import (
	"syscall"
	"unsafe"
)

// 2cd90691-12e2-11dc-9fed-001143a055f9
var IID_ID2D1Resource = GUID{0x2cd90691, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1ResourceVtbl struct {
	IUnknownVtbl
	GetFactory uintptr
}

type ID2D1Resource struct {
	vtbl *ID2D1ResourceVtbl
}

func (obj *ID2D1Resource) GUID() *GUID {
	return &IID_ID2D1Resource
}

func (obj *ID2D1Resource) QueryInterface(
	riid *GUID) (
	dest unsafe.Pointer,
	err error) {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid)
}

func (obj *ID2D1Resource) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *ID2D1Resource) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}

func (obj *ID2D1Resource) Parent() *IUnknown {
	return (*IUnknown)(unsafe.Pointer(obj))
}

func (obj *ID2D1Resource) GetFactory() (
	factory *ID2D1Factory) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl.GetFactory,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&factory)),
		0)
	return
}

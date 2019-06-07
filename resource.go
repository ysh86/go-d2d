// +build windows

package d2d

import (
	"syscall"
	"unsafe"
)

// 2cd90691-12e2-11dc-9fed-001143a055f9
var IID_ID2D1Resource = GUID{0x2cd90691, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1Resource struct {
	IUnknown
}

type vtblID2D1Resource struct {
	vtblIUnknown
	GetFactory uintptr
}

func (obj *ID2D1Resource) vtbl() *vtblID2D1Resource {
	return (*vtblID2D1Resource)(obj.unsafeVtbl)
}

func (obj *ID2D1Resource) GetFactory() (
	factory *ID2D1Factory) {
	var _, _, _ = syscall.Syscall(
		obj.vtbl().GetFactory,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&factory)),
		0)
	return
}

// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type ComObjectPtr interface {
	GUID() *GUID
	RawPtr() uintptr
}
type ComObjectPtrPtr interface {
	ComObjectPtr
	SetRawPtr(uintptr)
}

// HRESULT values
const (
	S_OK = 0
)

var IID_IUnknown = GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type IUnknown struct {
	unsafeVtbl unsafe.Pointer
}

// TODO: いらない
func (obj *IUnknown) GUID() *GUID {
	return &IID_IUnknown
}

func (obj *IUnknown) vtbl() *IUnknownVtbl {
	return (*IUnknownVtbl)(obj.unsafeVtbl)
}

func (obj *IUnknown) QueryInterface(
	iid *GUID) (
	dest unsafe.Pointer,
	err error) {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().QueryInterface,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&dest)))
	if ret != S_OK {
		err = fmt.Errorf("Query interface error: %#x", ret)
	}
	return
}

func (obj *IUnknown) AddRef() uint32 {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *IUnknown) Release() uint32 {
	var ret, _, _ = syscall.Syscall(
		obj.vtbl().Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

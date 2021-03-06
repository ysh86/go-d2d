// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modd2d1               = syscall.NewLazyDLL("d2d1.dll")
	procD2D1CreateFactory = modd2d1.NewProc("D2D1CreateFactory")
)

func D2D1CreateFactory(
	factoryType D2D1_FACTORY_TYPE,
	factoryOption *D2D1_FACTORY_OPTIONS) (
	factory *ID2D1Factory,
	err error) {
	var ret, _, _ = procD2D1CreateFactory.Call(
		uintptr(factoryType),
		uintptr(unsafe.Pointer(&IID_ID2D1Factory)),
		uintptr(unsafe.Pointer(factoryOption)),
		uintptr(unsafe.Pointer(&factory)))
	if ret != S_OK {
		err = fmt.Errorf("Fail to create factory: %#x", ret)
	}
	return
}

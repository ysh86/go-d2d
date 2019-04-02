// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"testing"
)

func TestCreateFactory(t *testing.T) {
	f, err := D2D1CreateFactory(
		D2D1_FACTORY_TYPE_SINGLE_THREADED,
		&D2D1_FACTORY_OPTIONS{D2D1_DEBUG_LEVEL_NONE})
	if err != nil || f == nil {
		t.Errorf("Factory is nil")
	}
	defer f.Unk().Release()

	i, err := f.Unk().QueryInterface(&IID_IUnknown)
	if err != nil || i == nil {
		t.Errorf("IUnknown is nil")
	}
	unk := (*IUnknown)(i)
	defer unk.Release()

	ii, err := f.Unk().QueryInterface(&IID_ID2D1Factory)
	if err != nil || ii == nil {
		t.Errorf("ID2D1Factory is nil")
	}
	ff := (*ID2D1Factory)(ii)
	defer ff.Unk().Release()

	iii, err := f.Parent().QueryInterface(&IID_ID2D1GeometrySink)
	if err == nil || iii != nil {
		t.Errorf("ID2D1GeometrySink is NOT nil")
	}
}

func TestGetDesktopDpi(t *testing.T) {
	f, _ := D2D1CreateFactory(D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
	defer f.Unk().Release()
	x, y := f.GetDesktopDpi()
	if x == 0 || y == 0 {
		t.Errorf("Dpi is zero: %f, %f", x, y)
	}
	t.Logf("Dpi: %f, %f", x, y)
}

func TestReloadSystemMetrics(t *testing.T) {
	f, _ := D2D1CreateFactory(D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
	defer f.Unk().Release()
	err := f.ReloadSystemMetrics()
	if err != nil {
		t.Errorf("ReloadSystemMetrics() returns error")
	}
}

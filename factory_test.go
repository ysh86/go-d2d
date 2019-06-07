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
	defer f.Release()

	p, err := f.QueryInterface(&IID_IUnknown)
	if err != nil || p == nil {
		t.Errorf("IUnknown is nil")
	}
	unk := (*IUnknown)(p)
	defer unk.Release()

	pp, err := f.QueryInterface(&IID_ID2D1Factory)
	if err != nil || pp == nil {
		t.Errorf("ID2D1Factory is nil")
	}
	ff := (*ID2D1Factory)(pp)
	defer ff.Release()

	ppp, err := f.QueryInterface(&IID_ID2D1GeometrySink)
	if err == nil || ppp != nil {
		t.Errorf("ID2D1GeometrySink is NOT nil")
	}
	fff := ff
	fff.AddRef()
	defer fff.Release()
}

func TestGetDesktopDpi(t *testing.T) {
	f, _ := D2D1CreateFactory(D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
	defer f.Release()
	x, y := f.GetDesktopDpi()
	if x == 0 || y == 0 {
		t.Errorf("Dpi is zero: %f, %f", x, y)
	}
	t.Logf("Dpi: %f, %f", x, y)
}

func TestReloadSystemMetrics(t *testing.T) {
	f, _ := D2D1CreateFactory(D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
	defer f.Release()
	err := f.ReloadSystemMetrics()
	if err != nil {
		t.Errorf("ReloadSystemMetrics() returns error")
	}
}

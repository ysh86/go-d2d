// +build windows

package d2d

type HWND uintptr
type HDC uintptr
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type ID2D1Bitmap struct {
	IUnknown
}

type ID2D1Layer struct {
	IUnknown
}

type ID2D1Mesh struct {
	IUnknown
}

type ID2D1BitmapBrush struct {
	IUnknown
}

type ID2D1GradientStopCollection struct {
	IUnknown
}

type ID2D1LinearGradientBrush struct {
	IUnknown
}

type ID2D1RadialGradientBrush struct {
	IUnknown
}

type ID2D1BitmapRenderTarget struct {
	IUnknown
}

type IDXGISurface struct {
	IUnknown
}

type IDWriteTextFormat struct {
	IUnknown
}

type IDWriteTextLayout struct {
	IUnknown
}

type IDWriteRenderingParams struct {
	IUnknown
}

type IDWriteFontFace struct {
	IUnknown
}
type DWRITE_GLYPH_OFFSET struct {
	AdvanceOffset  float32
	AscenderOffset float32
}
type DWRITE_GLYPH_RUN struct {
	FontFace      *IDWriteFontFace
	FontEmSize    float32
	GlyphCount    uint32
	GlyphIndices  *uint16
	GlyphAdvances *float32
	GlyphOffsets  *DWRITE_GLYPH_OFFSET
	IsSideways    int32
	BidiLevel     uint32
}
type DWRITE_MEASURING_MODE uint32

const (
	DWRITE_MEASURING_MODE_NATURAL     = 0
	DWRITE_MEASURING_MODE_GDI_CLASSIC = 1
	DWRITE_MEASURING_MODE_GDI_NATURAL = 2
)

type IWICBitmap struct {
	IUnknown
}

type IWICBitmapSource struct {
	IUnknown
}

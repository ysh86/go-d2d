// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

type D2D1_FACTORY_OPTIONS struct {
	DebugLevel D2D1_DEBUG_LEVEL
}

type D2D_RECT_F struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

type D2D1_RECT_F D2D_RECT_F

type D2D_SIZE_F struct {
	Width  float32
	Height float32
}

type D2D1_SIZE_F D2D_SIZE_F

type D2D_SIZE_U struct {
	Width  uint32
	Height uint32
}

type D2D1_SIZE_U D2D_SIZE_U

type D2D1_ROUNDED_RECT struct {
	Rect    D2D1_RECT_F
	RadiusX float32
	RadiusY float32
}

type D2D1_ELLIPSE struct {
	Point   D2D1_POINT_2F
	RadiusX float32
	RadiusY float32
}

type D2D_POINT_2F struct {
	X float32
	Y float32
}

type D2D1_POINT_2F D2D_POINT_2F

//             A11 A12 0
// (x, y, 1) * A21 A22 0 = (A11x + A21y + A31, A12x + A22y + A32, 1)
//             A31 A32 1
type D2D_MATRIX_3X2_F struct {
	A11 float32
	A12 float32
	A21 float32
	A22 float32
	A31 float32
	A32 float32
}

type D2D1_MATRIX_3X2_F D2D_MATRIX_3X2_F

type D2D1_TRIANGLE struct {
	Point1 D2D1_POINT_2F
	Point2 D2D1_POINT_2F
	Point3 D2D1_POINT_2F
}

type D2D1_BEZIER_SEGMENT struct {
	Point1 D2D1_POINT_2F
	Point2 D2D1_POINT_2F
	Point3 D2D1_POINT_2F
}

type D2D1_QUADRATIC_BEZIER_SEGMENT struct {
	Point1 D2D1_POINT_2F
	Point2 D2D1_POINT_2F
}

type D2D1_ARC_SEGMENT struct {
	Point          D2D1_POINT_2F
	Size           D2D1_SIZE_F
	RotationAngle  float32
	SweepDirection D2D1_SWEEP_DIRECTION
	ArcSize        D2D1_ARC_SIZE
}

type D2D1_STROKE_STYLE_PROPERTIES struct {
	StartCap   D2D1_CAP_STYLE
	EndCap     D2D1_CAP_STYLE
	DashCap    D2D1_CAP_STYLE
	LineJoin   D2D1_LINE_JOIN
	MiterLimit float32
	DashStyle  D2D1_DASH_STYLE
	DashOffset float32
}

type D2D1_DRAWING_STATE_DESCRIPTION struct {
	AntialiasMode     D2D1_ANTIALIAS_MODE
	TextAntialiasMode D2D1_TEXT_ANTIALIAS_MODE
	Tag1              D2D1_TAG
	Tag2              D2D1_TAG
	Transform         D2D1_MATRIX_3X2_F
}

type D2D1_RENDER_TARGET_PROPERTIES struct {
	Type        D2D1_RENDER_TARGET_TYPE
	PixelFormat D2D1_PIXEL_FORMAT
	DpiX        float32
	DpiY        float32
	Usage       D2D1_RENDER_TARGET_USAGE
	MinLevel    D2D1_FEATURE_LEVEL
}

type D2D1_HWND_RENDER_TARGET_PROPERTIES struct {
	Hwnd           uintptr
	PixelSize      D2D1_SIZE_U
	PresentOptions D2D1_PRESENT_OPTIONS
}

type D2D1_PIXEL_FORMAT struct {
	Format    DXGI_FORMAT
	AlphaMode D2D1_ALPHA_MODE
}

type D3DCOLORVALUE struct {
	R float32
	G float32
	B float32
	A float32
}

type D2D_COLOR_F D3DCOLORVALUE
type D2D1_COLOR_F D2D_COLOR_F

type D2D1_BRUSH_PROPERTIES struct {
	Opacity   float32
	Transform D2D1_MATRIX_3X2_F
}

type D2D1_GRADIENT_STOP struct {
	Position float32
	Color    D2D1_COLOR_F
}

type D2D1_LINEAR_GRADIENT_BRUSH_PROPERTIES struct {
	StartPoint D2D1_POINT_2F
	EndPoint   D2D1_POINT_2F
}

type D2D1_RADIAL_GRADIENT_BRUSH_PROPERTIES struct {
	Center               D2D1_POINT_2F
	GradientOriginOffset D2D1_POINT_2F
	RadiusX              float32
	RadiusY              float32
}

type D2D1_BITMAP_PROPERTIES struct {
	PixelFormat D2D1_PIXEL_FORMAT
	DpiX        float32
	DpiY        float32
}

type D2D1_BITMAP_BRUSH_PROPERTIES struct {
	ExtendModeX       D2D1_EXTEND_MODE
	ExtendModeY       D2D1_EXTEND_MODE
	InterpolationMode D2D1_BITMAP_INTERPOLATION_MODE
}

type D2D1_LAYER_PARAMETERS struct {
	ContentBounds     D2D1_RECT_F
	GeometricMask     *ID2D1Geometry
	MaskAntialiasMode D2D1_ANTIALIAS_MODE
	MaskTransform     D2D1_MATRIX_3X2_F
	Opacity           float32
	OpacityBrush      *ID2D1Brush
	LayerOptions      D2D1_LAYER_OPTIONS
}

// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

type D2D1_FACTORY_TYPE uint32

const (
	D2D1_FACTORY_TYPE_SINGLE_THREADED D2D1_FACTORY_TYPE = 0
	D2D1_FACTORY_TYPE_MULTI_THREADED  D2D1_FACTORY_TYPE = 1
)

type D2D1_DEBUG_LEVEL uint32

const (
	D2D1_DEBUG_LEVEL_NONE        D2D1_DEBUG_LEVEL = 0
	D2D1_DEBUG_LEVEL_ERROR       D2D1_DEBUG_LEVEL = 1
	D2D1_DEBUG_LEVEL_WARNING     D2D1_DEBUG_LEVEL = 2
	D2D1_DEBUG_LEVEL_INFORMATION D2D1_DEBUG_LEVEL = 3
)

type D2D1_FILL_MODE uint32

const (
	D2D1_FILL_MODE_ALTERNATE D2D1_FILL_MODE = 0
	D2D1_FILL_MODE_WINDING   D2D1_FILL_MODE = 1
)

type D2D1_CAP_STYLE uint32

const (
	D2D1_CAP_STYLE_FLAT     D2D1_CAP_STYLE = 0
	D2D1_CAP_STYLE_SQUARE   D2D1_CAP_STYLE = 1
	D2D1_CAP_STYLE_ROUND    D2D1_CAP_STYLE = 2
	D2D1_CAP_STYLE_TRIANGLE D2D1_CAP_STYLE = 3
)

type D2D1_LINE_JOIN uint32

const (
	D2D1_LINE_JOIN_MITER          D2D1_LINE_JOIN = 0
	D2D1_LINE_JOIN_BEVEL          D2D1_LINE_JOIN = 1
	D2D1_LINE_JOIN_ROUND          D2D1_LINE_JOIN = 2
	D2D1_LINE_JOIN_MITER_OR_BEVEL D2D1_LINE_JOIN = 3
)

type D2D1_DASH_STYLE uint32

const (
	D2D1_DASH_STYLE_SOLID        D2D1_DASH_STYLE = 0
	D2D1_DASH_STYLE_DASH         D2D1_DASH_STYLE = 1
	D2D1_DASH_STYLE_DOT          D2D1_DASH_STYLE = 2
	D2D1_DASH_STYLE_DASH_DOT     D2D1_DASH_STYLE = 3
	D2D1_DASH_STYLE_DASH_DOT_DOT D2D1_DASH_STYLE = 4
	D2D1_DASH_STYLE_CUSTOM       D2D1_DASH_STYLE = 5
)

type D2D1_GEOMETRY_RELATION uint32

const (
	D2D1_GEOMETRY_RELATION_UNKNOWN      D2D1_GEOMETRY_RELATION = 0
	D2D1_GEOMETRY_RELATION_DISJOINT     D2D1_GEOMETRY_RELATION = 1
	D2D1_GEOMETRY_RELATION_IS_CONTAINED D2D1_GEOMETRY_RELATION = 2
	D2D1_GEOMETRY_RELATION_CONTAINS     D2D1_GEOMETRY_RELATION = 3
	D2D1_GEOMETRY_RELATION_OVERLAP      D2D1_GEOMETRY_RELATION = 4
)

type D2D1_GEOMETRY_SIMPLIFICATION_OPTION uint32

const (
	D2D1_GEOMETRY_SIMPLIFICATION_OPTION_CUBICS_AND_LINES D2D1_GEOMETRY_SIMPLIFICATION_OPTION = 0
	D2D1_GEOMETRY_SIMPLIFICATION_OPTION_LINES            D2D1_GEOMETRY_SIMPLIFICATION_OPTION = 1
)

type D2D1_COMBINE_MODE uint32

const (
	D2D1_COMBINE_MODE_UNION     D2D1_COMBINE_MODE = 0
	D2D1_COMBINE_MODE_INTERSECT D2D1_COMBINE_MODE = 1
	D2D1_COMBINE_MODE_XOR       D2D1_COMBINE_MODE = 2
	D2D1_COMBINE_MODE_EXCLUDE   D2D1_COMBINE_MODE = 3
)

type D2D1_PATH_SEGMENT uint32

const (
	D2D1_PATH_SEGMENT_NONE                  D2D1_PATH_SEGMENT = 0x00000000
	D2D1_PATH_SEGMENT_FORCE_UNSTROKED       D2D1_PATH_SEGMENT = 0x00000001
	D2D1_PATH_SEGMENT_FORCE_ROUND_LINE_JOIN D2D1_PATH_SEGMENT = 0x00000002
)

type D2D1_FIGURE_BEGIN uint32

const (
	D2D1_FIGURE_BEGIN_FILLED D2D1_FIGURE_BEGIN = 0
	D2D1_FIGURE_BEGIN_HOLLOW D2D1_FIGURE_BEGIN = 1
)

type D2D1_FIGURE_END uint32

const (
	D2D1_FIGURE_END_OPEN   D2D1_FIGURE_END = 0
	D2D1_FIGURE_END_CLOSED D2D1_FIGURE_END = 1
)

type D2D1_SWEEP_DIRECTION uint32

const (
	D2D1_SWEEP_DIRECTION_COUNTER_CLOCKWISE D2D1_SWEEP_DIRECTION = 0
	D2D1_SWEEP_DIRECTION_CLOCKWISE         D2D1_SWEEP_DIRECTION = 1
)

type D2D1_ARC_SIZE uint32

const (
	D2D1_ARC_SIZE_SMALL D2D1_ARC_SIZE = 0
	D2D1_ARC_SIZE_LARGE D2D1_ARC_SIZE = 1
)

type D2D1_ANTIALIAS_MODE uint32

const (
	D2D1_ANTIALIAS_MODE_PER_PRIMITIVE D2D1_ANTIALIAS_MODE = 0
	D2D1_ANTIALIAS_MODE_ALIASED       D2D1_ANTIALIAS_MODE = 1
)

type D2D1_TEXT_ANTIALIAS_MODE uint32

const (
	D2D1_TEXT_ANTIALIAS_MODE_DEFAULT   D2D1_TEXT_ANTIALIAS_MODE = 0
	D2D1_TEXT_ANTIALIAS_MODE_CLEARTYPE D2D1_TEXT_ANTIALIAS_MODE = 1
	D2D1_TEXT_ANTIALIAS_MODE_GRAYSCALE D2D1_TEXT_ANTIALIAS_MODE = 2
	D2D1_TEXT_ANTIALIAS_MODE_ALIASED   D2D1_TEXT_ANTIALIAS_MODE = 3
)

type D2D1_TAG uint64

type D2D1_RENDER_TARGET_TYPE uint32

const (
	D2D1_RENDER_TARGET_TYPE_DEFAULT  D2D1_RENDER_TARGET_TYPE = 0
	D2D1_RENDER_TARGET_TYPE_SOFTWARE D2D1_RENDER_TARGET_TYPE = 1
	D2D1_RENDER_TARGET_TYPE_HARDWARE D2D1_RENDER_TARGET_TYPE = 2
)

type D2D1_RENDER_TARGET_USAGE uint32

const (
	D2D1_RENDER_TARGET_USAGE_NONE                  D2D1_RENDER_TARGET_USAGE = 0x00000000
	D2D1_RENDER_TARGET_USAGE_FORCE_BITMAP_REMOTING D2D1_RENDER_TARGET_USAGE = 0x00000001
	D2D1_RENDER_TARGET_USAGE_GDI_COMPATIBLE        D2D1_RENDER_TARGET_USAGE = 0x00000002
)

const (
	D3D10_FEATURE_LEVEL_9_1  = 0x9100
	D3D10_FEATURE_LEVEL_10_0 = 0xa100
)

type D2D1_FEATURE_LEVEL uint32

const (
	D2D1_FEATURE_LEVEL_DEFAULT D2D1_FEATURE_LEVEL = 0
	D2D1_FEATURE_LEVEL_9       D2D1_FEATURE_LEVEL = D3D10_FEATURE_LEVEL_9_1
	D2D1_FEATURE_LEVEL_10      D2D1_FEATURE_LEVEL = D3D10_FEATURE_LEVEL_10_0
)

type D2D1_PRESENT_OPTIONS uint32

const (
	D2D1_PRESENT_OPTIONS_NONE            D2D1_PRESENT_OPTIONS = 0x00000000
	D2D1_PRESENT_OPTIONS_RETAIN_CONTENTS D2D1_PRESENT_OPTIONS = 0x00000001
	D2D1_PRESENT_OPTIONS_IMMEDIATELY     D2D1_PRESENT_OPTIONS = 0x00000002
)

type DXGI_FORMAT uint32

const (
	DXGI_FORMAT_UNKNOWN                    DXGI_FORMAT = 0
	DXGI_FORMAT_R32G32B32A32_TYPELESS      DXGI_FORMAT = 1
	DXGI_FORMAT_R32G32B32A32_FLOAT         DXGI_FORMAT = 2
	DXGI_FORMAT_R32G32B32A32_UINT          DXGI_FORMAT = 3
	DXGI_FORMAT_R32G32B32A32_SINT          DXGI_FORMAT = 4
	DXGI_FORMAT_R32G32B32_TYPELESS         DXGI_FORMAT = 5
	DXGI_FORMAT_R32G32B32_FLOAT            DXGI_FORMAT = 6
	DXGI_FORMAT_R32G32B32_UINT             DXGI_FORMAT = 7
	DXGI_FORMAT_R32G32B32_SINT             DXGI_FORMAT = 8
	DXGI_FORMAT_R16G16B16A16_TYPELESS      DXGI_FORMAT = 9
	DXGI_FORMAT_R16G16B16A16_FLOAT         DXGI_FORMAT = 10
	DXGI_FORMAT_R16G16B16A16_UNORM         DXGI_FORMAT = 11
	DXGI_FORMAT_R16G16B16A16_UINT          DXGI_FORMAT = 12
	DXGI_FORMAT_R16G16B16A16_SNORM         DXGI_FORMAT = 13
	DXGI_FORMAT_R16G16B16A16_SINT          DXGI_FORMAT = 14
	DXGI_FORMAT_R32G32_TYPELESS            DXGI_FORMAT = 15
	DXGI_FORMAT_R32G32_FLOAT               DXGI_FORMAT = 16
	DXGI_FORMAT_R32G32_UINT                DXGI_FORMAT = 17
	DXGI_FORMAT_R32G32_SINT                DXGI_FORMAT = 18
	DXGI_FORMAT_R32G8X24_TYPELESS          DXGI_FORMAT = 19
	DXGI_FORMAT_D32_FLOAT_S8X24_UINT       DXGI_FORMAT = 20
	DXGI_FORMAT_R32_FLOAT_X8X24_TYPELESS   DXGI_FORMAT = 21
	DXGI_FORMAT_X32_TYPELESS_G8X24_UINT    DXGI_FORMAT = 22
	DXGI_FORMAT_R10G10B10A2_TYPELESS       DXGI_FORMAT = 23
	DXGI_FORMAT_R10G10B10A2_UNORM          DXGI_FORMAT = 24
	DXGI_FORMAT_R10G10B10A2_UINT           DXGI_FORMAT = 25
	DXGI_FORMAT_R11G11B10_FLOAT            DXGI_FORMAT = 26
	DXGI_FORMAT_R8G8B8A8_TYPELESS          DXGI_FORMAT = 27
	DXGI_FORMAT_R8G8B8A8_UNORM             DXGI_FORMAT = 28
	DXGI_FORMAT_R8G8B8A8_UNORM_SRGB        DXGI_FORMAT = 29
	DXGI_FORMAT_R8G8B8A8_UINT              DXGI_FORMAT = 30
	DXGI_FORMAT_R8G8B8A8_SNORM             DXGI_FORMAT = 31
	DXGI_FORMAT_R8G8B8A8_SINT              DXGI_FORMAT = 32
	DXGI_FORMAT_R16G16_TYPELESS            DXGI_FORMAT = 33
	DXGI_FORMAT_R16G16_FLOAT               DXGI_FORMAT = 34
	DXGI_FORMAT_R16G16_UNORM               DXGI_FORMAT = 35
	DXGI_FORMAT_R16G16_UINT                DXGI_FORMAT = 36
	DXGI_FORMAT_R16G16_SNORM               DXGI_FORMAT = 37
	DXGI_FORMAT_R16G16_SINT                DXGI_FORMAT = 38
	DXGI_FORMAT_R32_TYPELESS               DXGI_FORMAT = 39
	DXGI_FORMAT_D32_FLOAT                  DXGI_FORMAT = 40
	DXGI_FORMAT_R32_FLOAT                  DXGI_FORMAT = 41
	DXGI_FORMAT_R32_UINT                   DXGI_FORMAT = 42
	DXGI_FORMAT_R32_SINT                   DXGI_FORMAT = 43
	DXGI_FORMAT_R24G8_TYPELESS             DXGI_FORMAT = 44
	DXGI_FORMAT_D24_UNORM_S8_UINT          DXGI_FORMAT = 45
	DXGI_FORMAT_R24_UNORM_X8_TYPELESS      DXGI_FORMAT = 46
	DXGI_FORMAT_X24_TYPELESS_G8_UINT       DXGI_FORMAT = 47
	DXGI_FORMAT_R8G8_TYPELESS              DXGI_FORMAT = 48
	DXGI_FORMAT_R8G8_UNORM                 DXGI_FORMAT = 49
	DXGI_FORMAT_R8G8_UINT                  DXGI_FORMAT = 50
	DXGI_FORMAT_R8G8_SNORM                 DXGI_FORMAT = 51
	DXGI_FORMAT_R8G8_SINT                  DXGI_FORMAT = 52
	DXGI_FORMAT_R16_TYPELESS               DXGI_FORMAT = 53
	DXGI_FORMAT_R16_FLOAT                  DXGI_FORMAT = 54
	DXGI_FORMAT_D16_UNORM                  DXGI_FORMAT = 55
	DXGI_FORMAT_R16_UNORM                  DXGI_FORMAT = 56
	DXGI_FORMAT_R16_UINT                   DXGI_FORMAT = 57
	DXGI_FORMAT_R16_SNORM                  DXGI_FORMAT = 58
	DXGI_FORMAT_R16_SINT                   DXGI_FORMAT = 59
	DXGI_FORMAT_R8_TYPELESS                DXGI_FORMAT = 60
	DXGI_FORMAT_R8_UNORM                   DXGI_FORMAT = 61
	DXGI_FORMAT_R8_UINT                    DXGI_FORMAT = 62
	DXGI_FORMAT_R8_SNORM                   DXGI_FORMAT = 63
	DXGI_FORMAT_R8_SINT                    DXGI_FORMAT = 64
	DXGI_FORMAT_A8_UNORM                   DXGI_FORMAT = 65
	DXGI_FORMAT_R1_UNORM                   DXGI_FORMAT = 66
	DXGI_FORMAT_R9G9B9E5_SHAREDEXP         DXGI_FORMAT = 67
	DXGI_FORMAT_R8G8_B8G8_UNORM            DXGI_FORMAT = 68
	DXGI_FORMAT_G8R8_G8B8_UNORM            DXGI_FORMAT = 69
	DXGI_FORMAT_BC1_TYPELESS               DXGI_FORMAT = 70
	DXGI_FORMAT_BC1_UNORM                  DXGI_FORMAT = 71
	DXGI_FORMAT_BC1_UNORM_SRGB             DXGI_FORMAT = 72
	DXGI_FORMAT_BC2_TYPELESS               DXGI_FORMAT = 73
	DXGI_FORMAT_BC2_UNORM                  DXGI_FORMAT = 74
	DXGI_FORMAT_BC2_UNORM_SRGB             DXGI_FORMAT = 75
	DXGI_FORMAT_BC3_TYPELESS               DXGI_FORMAT = 76
	DXGI_FORMAT_BC3_UNORM                  DXGI_FORMAT = 77
	DXGI_FORMAT_BC3_UNORM_SRGB             DXGI_FORMAT = 78
	DXGI_FORMAT_BC4_TYPELESS               DXGI_FORMAT = 79
	DXGI_FORMAT_BC4_UNORM                  DXGI_FORMAT = 80
	DXGI_FORMAT_BC4_SNORM                  DXGI_FORMAT = 81
	DXGI_FORMAT_BC5_TYPELESS               DXGI_FORMAT = 82
	DXGI_FORMAT_BC5_UNORM                  DXGI_FORMAT = 83
	DXGI_FORMAT_BC5_SNORM                  DXGI_FORMAT = 84
	DXGI_FORMAT_B5G6R5_UNORM               DXGI_FORMAT = 85
	DXGI_FORMAT_B5G5R5A1_UNORM             DXGI_FORMAT = 86
	DXGI_FORMAT_B8G8R8A8_UNORM             DXGI_FORMAT = 87
	DXGI_FORMAT_B8G8R8X8_UNORM             DXGI_FORMAT = 88
	DXGI_FORMAT_R10G10B10_XR_BIAS_A2_UNORM DXGI_FORMAT = 89
	DXGI_FORMAT_B8G8R8A8_TYPELESS          DXGI_FORMAT = 90
	DXGI_FORMAT_B8G8R8A8_UNORM_SRGB        DXGI_FORMAT = 91
	DXGI_FORMAT_B8G8R8X8_TYPELESS          DXGI_FORMAT = 92
	DXGI_FORMAT_B8G8R8X8_UNORM_SRGB        DXGI_FORMAT = 93
	DXGI_FORMAT_BC6H_TYPELESS              DXGI_FORMAT = 94
	DXGI_FORMAT_BC6H_UF16                  DXGI_FORMAT = 95
	DXGI_FORMAT_BC6H_SF16                  DXGI_FORMAT = 96
	DXGI_FORMAT_BC7_TYPELESS               DXGI_FORMAT = 97
	DXGI_FORMAT_BC7_UNORM                  DXGI_FORMAT = 98
	DXGI_FORMAT_BC7_UNORM_SRGB             DXGI_FORMAT = 99
	DXGI_FORMAT_FORCE_UINT                 DXGI_FORMAT = 0xffffffff
)

type D2D1_ALPHA_MODE uint32

const (
	D2D1_ALPHA_MODE_UNKNOWN       D2D1_ALPHA_MODE = 0
	D2D1_ALPHA_MODE_PREMULTIPLIED D2D1_ALPHA_MODE = 1
	D2D1_ALPHA_MODE_STRAIGHT      D2D1_ALPHA_MODE = 2
	D2D1_ALPHA_MODE_IGNORE        D2D1_ALPHA_MODE = 3
)

type D2D1_GAMMA uint32

const (
	D2D1_GAMMA_2_2         = 0
	D2D1_GAMMA_1_0         = 1
	D2D1_GAMMA_FORCE_DWORD = 0xffffffff
)

type D2D1_EXTEND_MODE uint32

const (
	D2D1_EXTEND_MODE_CLAMP       = 0
	D2D1_EXTEND_MODE_WRAP        = 1
	D2D1_EXTEND_MODE_MIRROR      = 2
	D2D1_EXTEND_MODE_FORCE_DWORD = 0xffffffff
)

type D2D1_COMPATIBLE_RENDER_TARGET_OPTIONS uint32

const (
	D2D1_COMPATIBLE_RENDER_TARGET_OPTIONS_NONE           = 0x00000000
	D2D1_COMPATIBLE_RENDER_TARGET_OPTIONS_GDI_COMPATIBLE = 0x00000001
	D2D1_COMPATIBLE_RENDER_TARGET_OPTIONS_FORCE_DWORD    = 0xffffffff
)

type D2D1_DRAW_TEXT_OPTIONS uint32

const (
	D2D1_DRAW_TEXT_OPTIONS_NO_SNAP     = 0x00000001
	D2D1_DRAW_TEXT_OPTIONS_CLIP        = 0x00000002
	D2D1_DRAW_TEXT_OPTIONS_NONE        = 0x00000000
	D2D1_DRAW_TEXT_OPTIONS_FORCE_DWORD = 0xffffffff
)

type D2D1_WINDOW_STATE uint32

const (
	D2D1_WINDOW_STATE_NONE        = 0x0000000
	D2D1_WINDOW_STATE_OCCLUDED    = 0x0000001
	D2D1_WINDOW_STATE_FORCE_DWORD = 0xffffffff
)

type D2D1_BITMAP_INTERPOLATION_MODE uint32

const (
	D2D1_BITMAP_INTERPOLATION_MODE_NEAREST_NEIGHBOR = 0
	D2D1_BITMAP_INTERPOLATION_MODE_LINEAR           = 1
	D2D1_BITMAP_INTERPOLATION_MODE_FORCE_DWORD      = 0xffffffff
)

type D2D1_OPACITY_MASK_CONTENT uint32

const (
	D2D1_OPACITY_MASK_CONTENT_GRAPHICS            = 0
	D2D1_OPACITY_MASK_CONTENT_TEXT_NATURAL        = 1
	D2D1_OPACITY_MASK_CONTENT_TEXT_GDI_COMPATIBLE = 2
	D2D1_OPACITY_MASK_CONTENT_FORCE_DWORD         = 0xffffffff
)

type D2D1_LAYER_OPTIONS uint32

const (
	D2D1_LAYER_OPTIONS_NONE                     = 0x00000000
	D2D1_LAYER_OPTIONS_INITIALIZE_FOR_CLEARTYPE = 0x00000001
	D2D1_LAYER_OPTIONS_FORCE_DWORD              = 0xffffffff
)

cat << EOS > ../../geometry.go
// +build windows

package d2d

import (
    "fmt"
    "syscall"
    "unsafe"
)

EOS
../../gen.exe interface ID2D1Geometry >> ../../geometry.go
../../gen.exe interface ID2D1RectangleGeometry >> ../../geometry.go
../../gen.exe interface ID2D1RoundedRectangleGeometry >> ../../geometry.go
../../gen.exe interface ID2D1EllipseGeometry >> ../../geometry.go
../../gen.exe interface ID2D1GeometryGroup >> ../../geometry.go
../../gen.exe interface ID2D1TransformedGeometry >> ../../geometry.go
../../gen.exe interface ID2D1PathGeometry >> ../../geometry.go
../../gen.exe interface ID2D1SimplifiedGeometrySink >> ../../geometry.go
../../gen.exe interface ID2D1TessellationSink >> ../../geometry.go
../../gen.exe interface ID2D1GeometrySink >> ../../geometry.go
gofmt -w ../../geometry.go


cat << EOS > ../../render_target.go
// +build windows

package d2d

import (
    "fmt"
    "syscall"
    "unsafe"
)

EOS
../../gen.exe interface ID2D1RenderTarget >> ../../render_target.go
../../gen.exe interface ID2D1HwndRenderTarget >> ../../render_target.go
../../gen.exe interface ID2D1DCRenderTarget >> ../../render_target.go
gofmt -w ../../render_target.go

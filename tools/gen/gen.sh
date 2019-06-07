####################################
cat << EOS > ../../brush.go
// +build windows

package d2d

import (
    "syscall"
    "unsafe"
)

EOS
./gen.exe interface ID2D1Brush >> ../../brush.go
./gen.exe interface ID2D1SolidColorBrush >> ../../brush.go
gofmt -w ../../brush.go


####################################
cat << EOS > ../../drawing_state_block.go
// +build windows

package d2d

import (
    "syscall"
    "unsafe"
)

EOS
./gen.exe interface ID2D1DrawingStateBlock >> ../../drawing_state_block.go
gofmt -w ../../drawing_state_block.go


####################################
cat << EOS > ../../factory.go
// +build windows

package d2d

import (
    "fmt"
    "syscall"
    "unsafe"
)

EOS
./gen.exe interface ID2D1Factory >> ../../factory.go
gofmt -w ../../factory.go


####################################
cat << EOS > ../../geometry.go
// +build windows

package d2d

import (
    "fmt"
    "syscall"
    "unsafe"
)

EOS
./gen.exe interface ID2D1Geometry >> ../../geometry.go
./gen.exe interface ID2D1RectangleGeometry >> ../../geometry.go
./gen.exe interface ID2D1RoundedRectangleGeometry >> ../../geometry.go
./gen.exe interface ID2D1EllipseGeometry >> ../../geometry.go
./gen.exe interface ID2D1GeometryGroup >> ../../geometry.go
./gen.exe interface ID2D1TransformedGeometry >> ../../geometry.go
./gen.exe interface ID2D1PathGeometry >> ../../geometry.go
./gen.exe interface ID2D1SimplifiedGeometrySink >> ../../geometry.go
./gen.exe interface ID2D1TessellationSink >> ../../geometry.go
./gen.exe interface ID2D1GeometrySink >> ../../geometry.go
gofmt -w ../../geometry.go


####################################
cat << EOS > ../../render_target.go
// +build windows

package d2d

import (
    "fmt"
    "syscall"
    "unsafe"
)

EOS
./gen.exe interface ID2D1RenderTarget >> ../../render_target.go
./gen.exe interface ID2D1HwndRenderTarget >> ../../render_target.go
./gen.exe interface ID2D1DCRenderTarget >> ../../render_target.go
gofmt -w ../../render_target.go


####################################
cat << EOS > ../../resource.go
// +build windows

package d2d

import (
    "syscall"
    "unsafe"
)

EOS
./gen.exe interface ID2D1Resource >> ../../resource.go
gofmt -w ../../resource.go


####################################
cat << EOS > ../../stroke_style.go
// +build windows

package d2d

import (
    "syscall"
    "unsafe"
)

EOS
./gen.exe interface ID2D1StrokeStyle >> ../../stroke_style.go
gofmt -w ../../stroke_style.go

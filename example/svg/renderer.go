package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows"

	d2d "github.com/ysh86/go-d2d"
	"github.com/ysh86/gui"
	"github.com/ysh86/svg"
)

type svgRenderer struct {
	logger *log.Logger

	factory *d2d.ID2D1Factory

	// doc
	svg           *svg.Root
	geometryGroup *d2d.ID2D1GeometryGroup

	dpiX float32
	dpiY float32

	renderTarget        *d2d.ID2D1HwndRenderTarget
	lightSlateGrayBrush *d2d.ID2D1Brush
	cornflowerBlueBrush *d2d.ID2D1Brush
}

func NewSVGRenderer(svg *svg.Root) (*svgRenderer, error) {
	return &svgRenderer{svg: svg}, nil
}

func (r *svgRenderer) enableLog() error {
	r.logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	if r.logger != nil {
		r.logger.Print("start logging")
	}

	return nil
}

// Init calls methods for instantiating drawing resources
func (r *svgRenderer) Init() error {
	err := gui.CoInitializeEx(0, gui.COINIT_APARTMENTTHREADED|gui.COINIT_DISABLE_OLE1DDE)
	if err != nil {
		return fmt.Errorf("CoInitializeEx")
	}

	// Initialize device-indpendent resources, such
	// as the Direct2D factory.
	err = r.createDeviceIndependentResources()
	if err != nil {
		return err
	}

	// Because the CreateWindow function takes its size in pixels,
	// obtain the system DPI and use it to scale the window size.
	r.dpiX, r.dpiY = r.factory.GetDesktopDpi()

	return nil
}

// Deinit releases resources
func (r *svgRenderer) Deinit() {
	r.discardDeviceResources()

	if r.geometryGroup != nil {
		r.geometryGroup.Release()
		r.geometryGroup = nil
	}
	r.svg = nil

	if r.factory != nil {
		r.factory.Release()
		r.factory = nil
	}

	gui.CoUninitialize()
}

func (r *svgRenderer) Dpi() (float32, float32) {
	return r.dpiX, r.dpiY
}

func (r *svgRenderer) Update(width, height uint32) error {
	if r.renderTarget != nil {
		r.renderTarget.Resize(
			&d2d.D2D1_SIZE_U{Width: width, Height: height})
	}

	return nil
}

func (r *svgRenderer) Draw(nativeWindow uintptr) error {
	err := r.createDeviceResources(nativeWindow)
	if err != nil {
		return err
	}

	r.renderTarget.BeginDraw()

	white := d2d.D2D1_COLOR_F{R: 1, G: 1, B: 1, A: 1}
	//black := d2d.D2D1_COLOR_F{R: 0, G: 0, B: 0, A: 1}

	for _, g := range r.svg.Groups {
		// SVG
		// a c e   x   ax + cy + e
		// b d f . y = bx + dy + f
		// 0 0 1   1   0  + 0  + 1
		//
		// D2D
		//             A11 A12 0
		// (x, y, 1) * A21 A22 0 = (A11x + A21y + A31, A12x + A22y + A32, 1)
		//             A31 A32 1
		matrix := d2d.D2D1_MATRIX_3X2_F{
			A11: g.Transform.A,
			A12: g.Transform.B,
			A21: g.Transform.C,
			A22: g.Transform.D,
			A31: g.Transform.E,
			A32: g.Transform.F,
		}
		r.renderTarget.SetTransform(&matrix)
		r.renderTarget.Clear(&white) // TODO: 本当は none & layer 追加の方がいいか
		/*
			for _, gg := range g.Groups {
				stroke := d2d.D2D1_COLOR_F{
					R: float32(gg.Stroke.R) / 255.,
					G: float32(gg.Stroke.G) / 255.,
					B: float32(gg.Stroke.B) / 255.,
					A: float32(gg.Stroke.A) / 255.,
				}
				strokeBrush, err := r.renderTarget.CreateSolidColorBrush(
					&stroke,
					nil)
				if err != nil {
					r.discardDeviceResources()
					panic(err)
				}
				defer strokeBrush.Release()

				for _, p := range gg.Paths {
					sink, err := r.geometryGroup.Open()
					if err != nil {
						r.discardDeviceResources()
						panic(err)
					}
					for _, c := range p.D {
					}
					// TODO: gg.Fill
					sink.Close()

					r.renderTarget.DrawGeometry(
						&(r.geometryGroup.ID2D1Geometry),
						&(strokeBrush.ID2D1Brush),
						gg.StrokeWidth,
						nil)

				}
			}
		*/
		gg := g.Groups[0]
		stroke := d2d.D2D1_COLOR_F{
			R: float32(gg.Stroke.R) / 255.,
			G: float32(gg.Stroke.G) / 255.,
			B: float32(gg.Stroke.B) / 255.,
			A: float32(gg.Stroke.A) / 255.,
		}
		strokeBrush, err := r.renderTarget.CreateSolidColorBrush(
			&stroke,
			nil)
		if err != nil {
			r.discardDeviceResources()
			panic(err)
		}
		defer strokeBrush.Release()
		r.renderTarget.DrawGeometry(
			&(r.geometryGroup.ID2D1Geometry),
			&(strokeBrush.ID2D1Brush),
			gg.StrokeWidth,
			nil)

		_, _, err = r.renderTarget.EndDraw()
		if err != nil {
			r.discardDeviceResources()
		}
		return nil
	}

	size := r.renderTarget.GetSize()

	// Draw a grid background.
	width := int(size.Width)
	height := int(size.Height)
	for x := 0; x < width; x += 10 {
		r.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{X: float32(x), Y: 0.0},
			d2d.D2D1_POINT_2F{X: float32(x), Y: size.Height},
			r.lightSlateGrayBrush,
			0.5,
			nil)
	}
	for y := 0; y < height; y += 10 {
		r.renderTarget.DrawLine(
			d2d.D2D1_POINT_2F{X: 0.0, Y: float32(y)},
			d2d.D2D1_POINT_2F{X: size.Width, Y: float32(y)},
			r.lightSlateGrayBrush,
			0.5,
			nil)
	}

	// Draw two rectangles.
	rectangle1 := d2d.D2D1_RECT_F{
		Left:   size.Width/2.0 - 50.0,
		Top:    size.Height/2.0 - 50.0,
		Right:  size.Width/2.0 + 50.0,
		Bottom: size.Height/2.0 + 50.0,
	}
	rectangle2 := d2d.D2D1_RECT_F{
		Left:   size.Width/2.0 - 100.0,
		Top:    size.Height/2.0 - 100.0,
		Right:  size.Width/2.0 + 100.0,
		Bottom: size.Height/2.0 + 100.0,
	}
	// Draw a filled rectangle.
	r.renderTarget.FillRectangle(
		&rectangle1,
		r.lightSlateGrayBrush)
	// Draw the outline of a rectangle.
	r.renderTarget.DrawRectangle(
		&rectangle2,
		r.cornflowerBlueBrush,
		1,
		nil)

	_, _, err = r.renderTarget.EndDraw()
	if err != nil {
		r.discardDeviceResources()
	}

	return nil
}

// private methods

func (r *svgRenderer) createDeviceIndependentResources() (err error) {
	r.factory, err = d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)

	if r.svg == nil {
		return
	}

	var lineBuf *bytes.Buffer
	if r.logger != nil {
		lineBuf = new(bytes.Buffer)
	}

	gs := make([]*d2d.ID2D1Geometry, 0, len(r.svg.Groups))
	for _, g := range r.svg.Groups {
		ggs := make([]*d2d.ID2D1Geometry, 0, len(g.Groups))
		for _, gg := range g.Groups {
			ps := make([]*d2d.ID2D1Geometry, 0, len(gg.Paths))
			for _, p := range gg.Paths {
				if len(p.D) < 1 {
					continue
				}
				path, err := r.factory.CreatePathGeometry()
				if err != nil {
					panic(err)
				}
				sink, err := path.Open()
				if err != nil {
					panic(err)
				}

				// TODO: gg.Fill
				sink.SetFillMode(d2d.D2D1_FILL_MODE_WINDING)

				// SVG M
				pre := p.D[0]
				if pre.Command != "m" && pre.Command != "M" {
					continue
				}
				if len(pre.Points) < 1 {
					continue
				}
				prePoint := d2d.D2D1_POINT_2F{
					X: pre.Points[0].X,
					Y: pre.Points[0].Y,
				}
				sink.BeginFigure(prePoint, d2d.D2D1_FIGURE_BEGIN_FILLED)
				if r.logger != nil {
					lineBuf.Reset()
					lineBuf.WriteString(fmt.Sprintf("begin M:%v, ", prePoint))
				}
				// lines
				for _, svgPoint := range pre.Points[1:] {
					point := d2d.D2D1_POINT_2F{
						X: svgPoint.X,
						Y: svgPoint.Y,
					}

					if pre.Command == "m" {
						point.X += prePoint.X
						point.Y += prePoint.Y
					}
					prePoint = point

					sink.AddLine(point)
					if r.logger != nil {
						lineBuf.WriteString(fmt.Sprintf("M:%v, ", prePoint))
					}
				}

				// commands
				preP1 := prePoint
				for i, c := range p.D[1:] {
					switch c.Command {
					case "m", "M":
						panic(fmt.Errorf("Not implemented: %s", c.Command))
					case "l", "L":
						for _, svgPoint := range c.Points {
							point := d2d.D2D1_POINT_2F{
								X: svgPoint.X,
								Y: svgPoint.Y,
							}

							if c.Command == "l" {
								point.X += prePoint.X
								point.Y += prePoint.Y
							}
							prePoint = point
							preP1 = point

							sink.AddLine(point)
							if r.logger != nil {
								lineBuf.WriteString(fmt.Sprintf("L:%v, ", prePoint))
							}
						}
					case "s", "S":
						for j, svgPoint := range c.Points {
							if j%2 == 1 {
								p1 := d2d.D2D1_POINT_2F{
									X: preP1.X,
									Y: preP1.Y,
								}
								p2 := d2d.D2D1_POINT_2F{
									X: c.Points[j-1].X,
									Y: c.Points[j-1].Y,
								}
								point := d2d.D2D1_POINT_2F{
									X: svgPoint.X,
									Y: svgPoint.Y,
								}

								if c.Command == "s" {
									p2.X += prePoint.X
									p2.Y += prePoint.Y
									point.X += prePoint.X
									point.Y += prePoint.Y
								}
								prePoint = point
								preP1.X = -(p2.X - point.X) + point.X
								preP1.Y = -(p2.Y - point.Y) + point.Y

								sink.AddBezier(&d2d.D2D1_BEZIER_SEGMENT{
									Point1: p1,
									Point2: p2,
									Point3: point})
								if r.logger != nil {
									lineBuf.WriteString(fmt.Sprintf("S:%v, ", prePoint))
								}
							}
						}
					case "c", "C":
						for j, svgPoint := range c.Points {
							if j%3 == 2 {
								p1 := d2d.D2D1_POINT_2F{
									X: c.Points[j-2].X,
									Y: c.Points[j-2].Y,
								}
								p2 := d2d.D2D1_POINT_2F{
									X: c.Points[j-1].X,
									Y: c.Points[j-1].Y,
								}
								point := d2d.D2D1_POINT_2F{
									X: svgPoint.X,
									Y: svgPoint.Y,
								}

								if c.Command == "c" {
									p1.X += prePoint.X
									p1.Y += prePoint.Y
									p2.X += prePoint.X
									p2.Y += prePoint.Y
									point.X += prePoint.X
									point.Y += prePoint.Y
								}
								prePoint = point
								preP1.X = -(p2.X - point.X) + point.X
								preP1.Y = -(p2.Y - point.Y) + point.Y

								sink.AddBezier(&d2d.D2D1_BEZIER_SEGMENT{
									Point1: p1,
									Point2: p2,
									Point3: point})
								if r.logger != nil {
									lineBuf.WriteString(fmt.Sprintf("C:%v, ", prePoint))
								}
							}
						}
					case "h", "H":
						for _, svgPoint := range c.Points {
							point := d2d.D2D1_POINT_2F{
								X: svgPoint.X,
								Y: prePoint.Y,
							}

							if c.Command == "h" {
								point.X += prePoint.X
							}
							prePoint = point
							preP1 = point

							sink.AddLine(point)
							if r.logger != nil {
								lineBuf.WriteString(fmt.Sprintf("H:%v, ", prePoint))
							}
						}
					case "v", "V":
						for _, svgPoint := range c.Points {
							point := d2d.D2D1_POINT_2F{
								X: prePoint.X,
								Y: svgPoint.Y,
							}

							if c.Command == "v" {
								point.Y += prePoint.Y
							}
							prePoint = point
							preP1 = point

							sink.AddLine(point)
							if r.logger != nil {
								lineBuf.WriteString(fmt.Sprintf("V:%v, ", prePoint))
							}
						}
					case "z", "Z":
						sink.EndFigure(d2d.D2D1_FIGURE_END_CLOSED)
						if r.logger != nil {
							lineBuf.WriteString("close")
						}
						if i != len(p.D[1:])-1 {
							panic(fmt.Errorf("Not implemented: %s", c.Command))
						}
					default:
						panic(fmt.Errorf("Unknown: %s", c.Command))
					}
					pre = c
				}
				// no Z = open
				if pre.Command != "z" && pre.Command != "Z" {
					sink.EndFigure(d2d.D2D1_FIGURE_END_OPEN)
					if r.logger != nil {
						lineBuf.WriteString("open")
					}
				}

				sink.Close()
				if r.logger != nil {
					r.logger.Print(lineBuf.String())
				}
				ps = append(ps, &(path.ID2D1Geometry))
			}
			group, err := r.factory.CreateGeometryGroup(
				d2d.D2D1_FILL_MODE_WINDING,
				ps,
			)
			if err != nil {
				panic(err)
			}
			ggs = append(ggs, &(group.ID2D1Geometry))
		}
		group, err := r.factory.CreateGeometryGroup(
			d2d.D2D1_FILL_MODE_WINDING,
			ggs,
		)
		if err != nil {
			panic(err)
		}
		gs = append(gs, &(group.ID2D1Geometry))
	}
	r.geometryGroup, err = r.factory.CreateGeometryGroup(
		d2d.D2D1_FILL_MODE_WINDING,
		gs,
	)
	if err != nil {
		panic(err)
	}

	return
}

func (r *svgRenderer) createDeviceResources(nativeWindow uintptr) error {
	if r.renderTarget != nil {
		return nil // already created
	}

	hwnd := windows.Handle(nativeWindow)

	var rc gui.Rect
	err := gui.GetClientRect(hwnd, &rc)
	if err != nil {
		return err
	}
	r.renderTarget, err = r.factory.CreateHwndRenderTarget(
		d2d.RenderTargetProperties(),
		&d2d.D2D1_HWND_RENDER_TARGET_PROPERTIES{
			Hwnd: uintptr(hwnd),
			PixelSize: d2d.D2D1_SIZE_U{
				Width:  uint32(rc.Right - rc.Left),
				Height: uint32(rc.Bottom - rc.Top)},
			PresentOptions: d2d.D2D1_PRESENT_OPTIONS_NONE,
		})
	if err != nil {
		return err
	}

	lightSlateGray := d2d.D2D1_COLOR_F{R: 0x77 / 255., G: 0x88 / 255., B: 0x99 / 255., A: 1}
	lightSlateGrayBrush, err := r.renderTarget.CreateSolidColorBrush(
		&lightSlateGray,
		nil)
	if err != nil {
		return err
	}
	r.lightSlateGrayBrush = &(lightSlateGrayBrush.ID2D1Brush)

	cornflowerBlue := d2d.D2D1_COLOR_F{R: 0x64 / 255., G: 0x95 / 255., B: 0xED / 255., A: 1}
	cornflowerBlueBrush, err := r.renderTarget.CreateSolidColorBrush(
		&cornflowerBlue,
		nil)
	if err != nil {
		return err
	}
	r.cornflowerBlueBrush = &(cornflowerBlueBrush.ID2D1Brush)

	return nil
}

func (r *svgRenderer) discardDeviceResources() {
	if r.renderTarget != nil {
		r.renderTarget.Release()
		r.renderTarget = nil
	}
	if r.lightSlateGrayBrush != nil {
		r.lightSlateGrayBrush.Release()
		r.lightSlateGrayBrush = nil
	}
	if r.cornflowerBlueBrush != nil {
		r.cornflowerBlueBrush.Release()
		r.cornflowerBlueBrush = nil
	}
}

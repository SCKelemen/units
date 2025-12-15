package units

import "fmt"

// ═══════════════════════════════════════════════════════════════
//  Absolute Length Conversions
// ═══════════════════════════════════════════════════════════════

// Conversion ratios per CSS Values spec:
// The anchor unit is the pixel (px).
// 1in = 2.54cm = 96px
// 1cm = 96px/2.54 ≈ 37.795275591px
// 1mm = 1/10cm
// 1Q = 1/40cm
// 1pt = 1/72in
// 1pc = 12pt

const (
	pxPerInch = 96.0
	cmPerInch = 2.54
	pxPerCm   = pxPerInch / cmPerInch // ≈ 37.795275591
	pxPerMm   = pxPerCm / 10.0
	pxPerQ    = pxPerCm / 40.0
	pxPerPt   = pxPerInch / 72.0
	pxPerPc   = pxPerPt * 12.0
)

// ToPx converts an absolute length to pixels.
// Returns error if the length unit is relative (requires context).
//
// Example:
//
//	length := units.In(1)
//	px, err := length.ToPx()  // Returns units.Px(96)
func (l Length) ToPx() (Length, error) {
	if !l.IsAbsolute() {
		return Length{}, fmt.Errorf("cannot convert %s to pixels: unit is not absolute", l.Unit)
	}

	var pxValue float64
	switch l.Unit {
	case Pixel:
		pxValue = l.Value
	case Inch:
		pxValue = l.Value * pxPerInch
	case Centimeter:
		pxValue = l.Value * pxPerCm
	case Millimeter:
		pxValue = l.Value * pxPerMm
	case QuarterMillimeter:
		pxValue = l.Value * pxPerQ
	case Point:
		pxValue = l.Value * pxPerPt
	case Pica:
		pxValue = l.Value * pxPerPc
	default:
		return Length{}, fmt.Errorf("unknown absolute unit: %s", l.Unit)
	}

	return Px(pxValue), nil
}

// To converts an absolute length to another absolute unit.
// Returns error if either unit is relative.
//
// Example:
//
//	length := units.In(1)
//	cm, err := length.To(units.Centimeter)  // Returns units.Cm(2.54)
func (l Length) To(targetUnit LengthUnit) (Length, error) {
	// First convert to pixels
	px, err := l.ToPx()
	if err != nil {
		return Length{}, err
	}

	// Then convert from pixels to target unit
	var targetValue float64
	switch targetUnit {
	case Pixel:
		targetValue = px.Value
	case Inch:
		targetValue = px.Value / pxPerInch
	case Centimeter:
		targetValue = px.Value / pxPerCm
	case Millimeter:
		targetValue = px.Value / pxPerMm
	case QuarterMillimeter:
		targetValue = px.Value / pxPerQ
	case Point:
		targetValue = px.Value / pxPerPt
	case Pica:
		targetValue = px.Value / pxPerPc
	default:
		return Length{}, fmt.Errorf("cannot convert to %s: unit is not absolute", targetUnit)
	}

	return Length{Value: targetValue, Unit: targetUnit}, nil
}

// ═══════════════════════════════════════════════════════════════
//  Context-Aware Conversions
// ═══════════════════════════════════════════════════════════════

// Context provides the necessary information to resolve relative lengths
// to absolute values.
type Context struct {
	// Font metrics for font-relative units
	FontSize         float64 // For em, ch, ex, etc. (in px)
	RootFontSize     float64 // For rem, rch, rex, etc. (in px)
	XHeight          float64 // For ex (in px)
	CapHeight        float64 // For cap (in px)
	ChWidth          float64 // For ch - width of "0" glyph (in px)
	IcWidth          float64 // For ic - width of "水" glyph (in px)
	LineHeight       float64 // For lh (in px)
	RootLineHeight   float64 // For rlh (in px)

	// Viewport dimensions for viewport-relative units
	ViewportWidth  float64 // For vw, vi, etc. (in px)
	ViewportHeight float64 // For vh, vb, etc. (in px)

	// Container dimensions for container-relative units
	ContainerWidth  float64 // For cqw, cqi (in px)
	ContainerHeight float64 // For cqh, cqb (in px)
}

// Resolve converts any length (absolute or relative) to pixels using the provided context.
//
// Example:
//
//	ctx := units.Context{
//	    FontSize: 16.0,
//	    ViewportWidth: 1920.0,
//	}
//	length := units.Em(2)
//	px, err := length.Resolve(ctx)  // Returns units.Px(32)
func (l Length) Resolve(ctx Context) (Length, error) {
	// Absolute units can be converted directly
	if l.IsAbsolute() {
		return l.ToPx()
	}

	var pxValue float64

	// Font-relative units
	switch l.Unit {
	case EmUnit:
		if ctx.FontSize == 0 {
			return Length{}, fmt.Errorf("cannot resolve em: FontSize not set in context")
		}
		pxValue = l.Value * ctx.FontSize

	case RemUnit:
		if ctx.RootFontSize == 0 {
			return Length{}, fmt.Errorf("cannot resolve rem: RootFontSize not set in context")
		}
		pxValue = l.Value * ctx.RootFontSize

	case ExUnit:
		if ctx.XHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve ex: XHeight not set in context")
		}
		pxValue = l.Value * ctx.XHeight

	case CapUnit:
		if ctx.CapHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve cap: CapHeight not set in context")
		}
		pxValue = l.Value * ctx.CapHeight

	case ChUnit:
		if ctx.ChWidth == 0 {
			return Length{}, fmt.Errorf("cannot resolve ch: ChWidth not set in context")
		}
		pxValue = l.Value * ctx.ChWidth

	case IcUnit:
		if ctx.IcWidth == 0 {
			return Length{}, fmt.Errorf("cannot resolve ic: IcWidth not set in context")
		}
		pxValue = l.Value * ctx.IcWidth

	case LhUnit:
		if ctx.LineHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve lh: LineHeight not set in context")
		}
		pxValue = l.Value * ctx.LineHeight

	case RlhUnit:
		if ctx.RootLineHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve rlh: RootLineHeight not set in context")
		}
		pxValue = l.Value * ctx.RootLineHeight

	// Viewport-relative units
	case VwUnit:
		if ctx.ViewportWidth == 0 {
			return Length{}, fmt.Errorf("cannot resolve vw: ViewportWidth not set in context")
		}
		pxValue = l.Value * ctx.ViewportWidth / 100.0

	case VhUnit:
		if ctx.ViewportHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve vh: ViewportHeight not set in context")
		}
		pxValue = l.Value * ctx.ViewportHeight / 100.0

	case VminUnit:
		if ctx.ViewportWidth == 0 || ctx.ViewportHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve vmin: ViewportWidth and ViewportHeight not set in context")
		}
		min := ctx.ViewportWidth
		if ctx.ViewportHeight < min {
			min = ctx.ViewportHeight
		}
		pxValue = l.Value * min / 100.0

	case VmaxUnit:
		if ctx.ViewportWidth == 0 || ctx.ViewportHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve vmax: ViewportWidth and ViewportHeight not set in context")
		}
		max := ctx.ViewportWidth
		if ctx.ViewportHeight > max {
			max = ctx.ViewportHeight
		}
		pxValue = l.Value * max / 100.0

	// Container-relative units
	case CqwUnit:
		if ctx.ContainerWidth == 0 {
			return Length{}, fmt.Errorf("cannot resolve cqw: ContainerWidth not set in context")
		}
		pxValue = l.Value * ctx.ContainerWidth / 100.0

	case CqhUnit:
		if ctx.ContainerHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve cqh: ContainerHeight not set in context")
		}
		pxValue = l.Value * ctx.ContainerHeight / 100.0

	case CqminUnit:
		if ctx.ContainerWidth == 0 || ctx.ContainerHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve cqmin: ContainerWidth and ContainerHeight not set in context")
		}
		min := ctx.ContainerWidth
		if ctx.ContainerHeight < min {
			min = ctx.ContainerHeight
		}
		pxValue = l.Value * min / 100.0

	case CqmaxUnit:
		if ctx.ContainerWidth == 0 || ctx.ContainerHeight == 0 {
			return Length{}, fmt.Errorf("cannot resolve cqmax: ContainerWidth and ContainerHeight not set in context")
		}
		max := ctx.ContainerWidth
		if ctx.ContainerHeight > max {
			max = ctx.ContainerHeight
		}
		pxValue = l.Value * max / 100.0

	default:
		return Length{}, fmt.Errorf("unsupported unit for resolution: %s", l.Unit)
	}

	return Px(pxValue), nil
}

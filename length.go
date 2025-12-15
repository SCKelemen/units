// Package units implements the CSS Values and Units Module Level 4 specification.
//
// This package provides type-safe value types for CSS units including:
//   - Absolute lengths (px, cm, mm, in, pt, pc)
//   - Font-relative lengths (em, rem, ch, ex, cap, ic, lh)
//   - Viewport-relative lengths (vw, vh, vmin, vmax, vb, vi)
//   - Container-relative lengths (cqw, cqh, cqi, cqb, cqmin, cqmax)
//
// Originally implemented in github.com/SCKelemen/layout and extracted as a
// standalone package for reuse across layout engines, text rendering, and
// other CSS-based projects.
//
// References:
//   - CSS Values and Units Module Level 4: https://www.w3.org/TR/css-values-4/
//   - MDN Web Docs - CSS values and units: https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Values_and_Units
//   - web.dev - Learn CSS - Sizing Units: https://web.dev/learn/css/sizing/
//
// See also:
//   - CSS Text Module Level 3: https://www.w3.org/TR/css-text-3/
//   - CSS Writing Modes Level 4: https://www.w3.org/TR/css-writing-modes-4/
//   - CSS Logical Properties: https://www.w3.org/TR/css-logical-1/
package units

import "fmt"

// ═══════════════════════════════════════════════════════════════
//  Core Length Type
// ═══════════════════════════════════════════════════════════════

// Length represents a CSS length value with an explicit unit.
//
// Per CSS Values spec, lengths can be:
//   - Absolute: px, cm, mm, Q, in, pt, pc
//   - Font-relative: em, rem, ex, rex, cap, rcap, ch, rch, ic, ric, lh, rlh
//   - Viewport-relative: vw, vh, vmin, vmax, vb, vi, svw, svh, lvw, lvh, dvw, dvh
//   - Container-relative: cqw, cqh, cqi, cqb, cqmin, cqmax
//
// Example:
//
//	width := units.Px(400)    // 400 pixels
//	margin := units.Em(1.5)   // 1.5em
//	height := units.Vh(100)   // 100vh (full viewport height)
type Length struct {
	Value float64
	Unit  LengthUnit
}

// LengthUnit specifies the unit of measurement for a length value.
type LengthUnit string

// ═══════════════════════════════════════════════════════════════
//  Absolute Length Units
// ═══════════════════════════════════════════════════════════════

// Absolute length units are fixed in relation to each other and anchored
// to some physical measurement. They are mainly useful when the output
// environment is known.
//
// The anchor unit is the pixel (px). The reference pixel is the visual
// angle of one pixel on a device with a pixel density of 96dpi and a
// distance from the reader of an arm's length.
const (
	// Base unit names (uppercase to avoid conflict with constructor functions)

	// PX - Absolute length, anchor unit
	// 1px = 1/96 of 1 inch
	// Reference: https://www.w3.org/TR/css-values-4/#absolute-lengths
	PX LengthUnit = "px"

	// CM - Absolute length
	// 1cm = 96px/2.54
	// Reference: https://www.w3.org/TR/css-values-4/#absolute-lengths
	CM LengthUnit = "cm"

	// MM - Absolute length
	// 1mm = 1/10 of 1cm
	// Reference: https://www.w3.org/TR/css-values-4/#absolute-lengths
	MM LengthUnit = "mm"

	// QQ - Absolute length (quarter-millimeter)
	// 1Q = 1/40 of 1cm
	// Reference: https://www.w3.org/TR/css-values-4/#absolute-lengths
	QQ LengthUnit = "Q"

	// IN - Absolute length
	// 1in = 2.54cm = 96px
	// Reference: https://www.w3.org/TR/css-values-4/#absolute-lengths
	IN LengthUnit = "in"

	// PT - Absolute length (typography)
	// 1pt = 1/72 of 1in
	// Reference: https://www.w3.org/TR/css-values-4/#absolute-lengths
	PT LengthUnit = "pt"

	// PC - Absolute length (typography)
	// 1pc = 1/6 of 1in = 12pt
	// Reference: https://www.w3.org/TR/css-values-4/#absolute-lengths
	PC LengthUnit = "pc"

	// Convenience aliases (long-form names only, short names conflict with constructors)
	Pixel             LengthUnit = PX
	Centimeter        LengthUnit = CM
	Millimeter        LengthUnit = MM
	QuarterMillimeter LengthUnit = QQ
	Inch              LengthUnit = IN
	Point             LengthUnit = PT
	Pica              LengthUnit = PC
)

// ═══════════════════════════════════════════════════════════════
//  Font-Relative Length Units
// ═══════════════════════════════════════════════════════════════

// Font-relative lengths define length values in terms of the font metrics
// of the element on which they are used (or, for the "root" variants, the
// root element).
const (
	// Base unit names (uppercase to avoid conflict with constructor functions)

	// EM - Font-relative length
	// Equal to the computed value of font-size on the element
	EM LengthUnit = "em"

	// REM - Font-relative length
	// Equal to the computed value of font-size on the root element
	REM LengthUnit = "rem"

	// EX - Font-relative length
	// Equal to the x-height of the font (height of lowercase 'x')
	EX LengthUnit = "ex"

	// REX - Font-relative length
	// Equal to the x-height of the root element's font
	REX LengthUnit = "rex"

	// CAP - Font-relative length
	// Equal to the cap-height of the font (height of capital letters)
	CAP LengthUnit = "cap"

	// RCAP - Font-relative length
	// Equal to the cap-height of the root element's font
	RCAP LengthUnit = "rcap"

	// CH - Font-relative length
	// Equal to the advance measure of the "0" (ZERO, U+0030) glyph
	// Used for monospace width calculations (terminal cells)
	CH LengthUnit = "ch"

	// RCH - Font-relative length
	// Equal to the advance measure of "0" in the root element's font
	RCH LengthUnit = "rch"

	// IC - Font-relative length
	// Equal to the advance measure of the "水" (CJK water ideograph, U+6C34) glyph
	// Used for ideographic character width
	IC LengthUnit = "ic"

	// RIC - Font-relative length
	// Equal to the advance measure of "水" in the root element's font
	RIC LengthUnit = "ric"

	// LH - Font-relative length
	// Equal to the computed value of line-height on the element
	LH LengthUnit = "lh"

	// RLH - Font-relative length
	// Equal to the computed value of line-height on the root element
	RLH LengthUnit = "rlh"

	// Convenience aliases (Unit suffix only, short names conflict with constructors)
	EmUnit   LengthUnit = EM
	RemUnit  LengthUnit = REM
	ExUnit   LengthUnit = EX
	RexUnit  LengthUnit = REX
	CapUnit  LengthUnit = CAP
	RcapUnit LengthUnit = RCAP
	ChUnit   LengthUnit = CH
	RchUnit  LengthUnit = RCH
	IcUnit   LengthUnit = IC
	RicUnit  LengthUnit = RIC
	LhUnit   LengthUnit = LH
	RlhUnit  LengthUnit = RLH
)

// ═══════════════════════════════════════════════════════════════
//  Viewport-Relative Length Units
// ═══════════════════════════════════════════════════════════════

// Viewport-percentage lengths define length values relative to the size
// of the initial containing block (viewport).
const (
	// Base unit names (uppercase to avoid conflict with constructor functions)

	// VW - Viewport width
	// Equal to 1% of viewport width
	VW LengthUnit = "vw"

	// VH - Viewport height
	// Equal to 1% of viewport height
	VH LengthUnit = "vh"

	// VMIN - Viewport minimum
	// Equal to the smaller of vw or vh
	VMIN LengthUnit = "vmin"

	// VMAX - Viewport maximum
	// Equal to the larger of vw or vh
	VMAX LengthUnit = "vmax"

	// VB - Viewport block size
	// Equal to 1% of viewport size in block axis
	VB LengthUnit = "vb"

	// VI - Viewport inline size
	// Equal to 1% of viewport size in inline axis
	VI LengthUnit = "vi"

	// Small viewport units (smallest possible viewport)
	SVW LengthUnit = "svw"
	SVH LengthUnit = "svh"
	SVB LengthUnit = "svb"
	SVI LengthUnit = "svi"

	// Large viewport units (largest possible viewport)
	LVW LengthUnit = "lvw"
	LVH LengthUnit = "lvh"
	LVB LengthUnit = "lvb"
	LVI LengthUnit = "lvi"

	// Dynamic viewport units (current viewport, accounting for dynamic UI)
	DVW LengthUnit = "dvw"
	DVH LengthUnit = "dvh"
	DVB LengthUnit = "dvb"
	DVI LengthUnit = "dvi"

	// Convenience aliases (Unit suffix only, short names conflict with constructors)
	VwUnit   LengthUnit = VW
	VhUnit   LengthUnit = VH
	VminUnit LengthUnit = VMIN
	VmaxUnit LengthUnit = VMAX
	VbUnit   LengthUnit = VB
	ViUnit   LengthUnit = VI
	SvwUnit  LengthUnit = SVW
	SvhUnit  LengthUnit = SVH
	SvbUnit  LengthUnit = SVB
	SviUnit  LengthUnit = SVI
	LvwUnit  LengthUnit = LVW
	LvhUnit  LengthUnit = LVH
	LvbUnit  LengthUnit = LVB
	LviUnit  LengthUnit = LVI
	DvwUnit  LengthUnit = DVW
	DvhUnit  LengthUnit = DVH
	DvbUnit  LengthUnit = DVB
	DviUnit  LengthUnit = DVI
)

// ═══════════════════════════════════════════════════════════════
//  Container-Relative Length Units
// ═══════════════════════════════════════════════════════════════

// Container query length units specify lengths relative to the dimensions
// of a query container.
const (
	// Base unit names (uppercase to avoid conflict with constructor functions)

	// CQW - Container query width
	// Equal to 1% of query container's width
	CQW LengthUnit = "cqw"

	// CQH - Container query height
	// Equal to 1% of query container's height
	CQH LengthUnit = "cqh"

	// CQI - Container query inline size
	// Equal to 1% of query container's inline size
	CQI LengthUnit = "cqi"

	// CQB - Container query block size
	// Equal to 1% of query container's block size
	CQB LengthUnit = "cqb"

	// CQMIN - Container query minimum
	// Equal to the smaller of cqi or cqb
	CQMIN LengthUnit = "cqmin"

	// CQMAX - Container query maximum
	// Equal to the larger of cqi or cqb
	CQMAX LengthUnit = "cqmax"

	// Convenience aliases (Unit suffix only, short names conflict with constructors)
	CqwUnit   LengthUnit = CQW
	CqhUnit   LengthUnit = CQH
	CqiUnit   LengthUnit = CQI
	CqbUnit   LengthUnit = CQB
	CqminUnit LengthUnit = CQMIN
	CqmaxUnit LengthUnit = CQMAX
)

// ═══════════════════════════════════════════════════════════════
//  Constructor Functions - Absolute Units
// ═══════════════════════════════════════════════════════════════

// Px creates a length in pixels (px).
func Px(value float64) Length {
	return Length{Value: value, Unit: PX}
}

// Cm creates a length in centimeters (cm).
func Cm(value float64) Length {
	return Length{Value: value, Unit: CM}
}

// Mm creates a length in millimeters (mm).
func Mm(value float64) Length {
	return Length{Value: value, Unit: MM}
}

// Q creates a length in quarter-millimeters (Q).
func Q(value float64) Length {
	return Length{Value: value, Unit: QQ}
}

// In creates a length in inches (in).
func In(value float64) Length {
	return Length{Value: value, Unit: IN}
}

// Pt creates a length in points (pt).
func Pt(value float64) Length {
	return Length{Value: value, Unit: PT}
}

// Pc creates a length in picas (pc).
func Pc(value float64) Length {
	return Length{Value: value, Unit: PC}
}

// ═══════════════════════════════════════════════════════════════
//  Constructor Functions - Font-Relative Units
// ═══════════════════════════════════════════════════════════════

// Em creates a length in em units.
func Em(value float64) Length {
	return Length{Value: value, Unit: EM}
}

// Rem creates a length in rem units.
func Rem(value float64) Length {
	return Length{Value: value, Unit: REM}
}

// Ex creates a length in ex units.
func Ex(value float64) Length {
	return Length{Value: value, Unit: EX}
}

// Rex creates a length in rex units.
func Rex(value float64) Length {
	return Length{Value: value, Unit: REX}
}

// Cap creates a length in cap units.
func Cap(value float64) Length {
	return Length{Value: value, Unit: CAP}
}

// Rcap creates a length in rcap units.
func Rcap(value float64) Length {
	return Length{Value: value, Unit: RCAP}
}

// Ch creates a length in ch units (character width, typically "0" glyph).
func Ch(value float64) Length {
	return Length{Value: value, Unit: CH}
}

// Rch creates a length in rch units.
func Rch(value float64) Length {
	return Length{Value: value, Unit: RCH}
}

// Ic creates a length in ic units (ideographic character width).
func Ic(value float64) Length {
	return Length{Value: value, Unit: IC}
}

// Ric creates a length in ric units.
func Ric(value float64) Length {
	return Length{Value: value, Unit: RIC}
}

// Lh creates a length in lh units (line height).
func Lh(value float64) Length {
	return Length{Value: value, Unit: LH}
}

// Rlh creates a length in rlh units.
func Rlh(value float64) Length {
	return Length{Value: value, Unit: RLH}
}

// ═══════════════════════════════════════════════════════════════
//  Constructor Functions - Viewport-Relative Units
// ═══════════════════════════════════════════════════════════════

// Vw creates a length in vw units (viewport width percentage).
func Vw(value float64) Length {
	return Length{Value: value, Unit: VW}
}

// Vh creates a length in vh units (viewport height percentage).
func Vh(value float64) Length {
	return Length{Value: value, Unit: VH}
}

// Vmin creates a length in vmin units.
func Vmin(value float64) Length {
	return Length{Value: value, Unit: VMIN}
}

// Vmax creates a length in vmax units.
func Vmax(value float64) Length {
	return Length{Value: value, Unit: VMAX}
}

// Vb creates a length in vb units (viewport block size).
func Vb(value float64) Length {
	return Length{Value: value, Unit: VB}
}

// Vi creates a length in vi units (viewport inline size).
func Vi(value float64) Length {
	return Length{Value: value, Unit: VI}
}

// ═══════════════════════════════════════════════════════════════
//  Constructor Functions - Container-Relative Units
// ═══════════════════════════════════════════════════════════════

// Cqw creates a length in cqw units (container query width).
func Cqw(value float64) Length {
	return Length{Value: value, Unit: CQW}
}

// Cqh creates a length in cqh units (container query height).
func Cqh(value float64) Length {
	return Length{Value: value, Unit: CQH}
}

// Cqi creates a length in cqi units (container query inline size).
func Cqi(value float64) Length {
	return Length{Value: value, Unit: CQI}
}

// Cqb creates a length in cqb units (container query block size).
func Cqb(value float64) Length {
	return Length{Value: value, Unit: CQB}
}

// Cqmin creates a length in cqmin units.
func Cqmin(value float64) Length {
	return Length{Value: value, Unit: CQMIN}
}

// Cqmax creates a length in cqmax units.
func Cqmax(value float64) Length {
	return Length{Value: value, Unit: CQMAX}
}

// ═══════════════════════════════════════════════════════════════
//  Length Operations
// ═══════════════════════════════════════════════════════════════

// String returns a CSS-compatible string representation.
func (l Length) String() string {
	return fmt.Sprintf("%.2f%s", l.Value, l.Unit)
}

// IsZero returns true if the length value is zero.
func (l Length) IsZero() bool {
	return l.Value == 0
}

// IsAbsolute returns true if the unit is an absolute length unit.
func (l Length) IsAbsolute() bool {
	switch l.Unit {
	case PX, CM, MM, QQ, IN, PT, PC:
		return true
	default:
		return false
	}
}

// IsFontRelative returns true if the unit is font-relative.
func (l Length) IsFontRelative() bool {
	switch l.Unit {
	case EM, REM, EX, REX, CAP, RCAP, CH, RCH, IC, RIC, LH, RLH:
		return true
	default:
		return false
	}
}

// IsViewportRelative returns true if the unit is viewport-relative.
func (l Length) IsViewportRelative() bool {
	switch l.Unit {
	case VW, VH, VMIN, VMAX, VB, VI, SVW, SVH, SVB, SVI, LVW, LVH, LVB, LVI, DVW, DVH, DVB, DVI:
		return true
	default:
		return false
	}
}

// IsContainerRelative returns true if the unit is container-relative.
func (l Length) IsContainerRelative() bool {
	switch l.Unit {
	case CQW, CQH, CQI, CQB, CQMIN, CQMAX:
		return true
	default:
		return false
	}
}

// Add adds two lengths with the same unit.
// Panics if units don't match.
func (l Length) Add(other Length) Length {
	if l.Unit != other.Unit {
		panic(fmt.Sprintf("cannot add lengths with different units: %s + %s", l.Unit, other.Unit))
	}
	return Length{Value: l.Value + other.Value, Unit: l.Unit}
}

// Sub subtracts two lengths with the same unit.
// Panics if units don't match.
func (l Length) Sub(other Length) Length {
	if l.Unit != other.Unit {
		panic(fmt.Sprintf("cannot subtract lengths with different units: %s - %s", l.Unit, other.Unit))
	}
	return Length{Value: l.Value - other.Value, Unit: l.Unit}
}

// Mul multiplies a length by a scalar.
func (l Length) Mul(scalar float64) Length {
	return Length{Value: l.Value * scalar, Unit: l.Unit}
}

// Div divides a length by a scalar.
func (l Length) Div(scalar float64) Length {
	return Length{Value: l.Value / scalar, Unit: l.Unit}
}

// LessThan returns true if this length is less than another.
// Panics if units don't match.
func (l Length) LessThan(other Length) bool {
	if l.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare lengths with different units: %s < %s", l.Unit, other.Unit))
	}
	return l.Value < other.Value
}

// GreaterThan returns true if this length is greater than another.
// Panics if units don't match.
func (l Length) GreaterThan(other Length) bool {
	if l.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare lengths with different units: %s > %s", l.Unit, other.Unit))
	}
	return l.Value > other.Value
}

// Raw returns the raw float64 value, discarding unit information.
// Use this when interfacing with APIs that need raw numbers.
func (l Length) Raw() float64 {
	return l.Value
}

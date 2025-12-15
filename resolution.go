package units

import "fmt"

// ═══════════════════════════════════════════════════════════════
//  Resolution Units - CSS Values Level 4 Section 6.4
// ═══════════════════════════════════════════════════════════════

// Resolution represents a CSS resolution value.
//
// Resolution values are used in CSS media queries and image properties
// to describe the density of pixels. CSS supports dots per inch (dpi),
// dots per centimeter (dpcm), and dots per pixel unit (dppx).
//
// References:
//   - CSS Values Level 4 - Resolution: https://www.w3.org/TR/css-values-4/#resolution
//   - MDN - CSS resolution: https://developer.mozilla.org/en-US/docs/Web/CSS/resolution
//   - web.dev - Responsive images: https://web.dev/learn/design/responsive-images/
type Resolution struct {
	Value float64
	Unit  ResolutionUnit
}

// ResolutionUnit represents a CSS resolution unit type.
type ResolutionUnit string

// CSS resolution units per CSS Values Level 4.
//
// Relationship:
//   - 1dppx = 96dpi
//   - 1dpcm = 96dpi / 2.54 ≈ 37.795dpi
//   - 1in = 2.54cm
const (
	DotsPerInch      ResolutionUnit = "dpi"  // Dots per inch
	DotsPerCentimeter ResolutionUnit = "dpcm" // Dots per centimeter
	DotsPerPixel     ResolutionUnit = "dppx" // Dots per pixel unit
)

// ═══════════════════════════════════════════════════════════════
//  Constructors
// ═══════════════════════════════════════════════════════════════

// Dpi creates a resolution value in dots per inch.
//
// Example:
//
//	res := units.Dpi(96)  // Standard screen resolution
func Dpi(value float64) Resolution {
	return Resolution{Value: value, Unit: DotsPerInch}
}

// Dpcm creates a resolution value in dots per centimeter.
//
// Example:
//
//	res := units.Dpcm(37.795)  // Equivalent to 96dpi
func Dpcm(value float64) Resolution {
	return Resolution{Value: value, Unit: DotsPerCentimeter}
}

// Dppx creates a resolution value in dots per pixel unit.
//
// Example:
//
//	res := units.Dppx(2)  // 2x "Retina" display
func Dppx(value float64) Resolution {
	return Resolution{Value: value, Unit: DotsPerPixel}
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the resolution.
func (r Resolution) String() string {
	return fmt.Sprintf("%.2f%s", r.Value, r.Unit)
}

// IsZero returns true if the resolution value is zero.
func (r Resolution) IsZero() bool {
	return r.Value == 0
}

// Raw returns the raw numeric value of the resolution.
func (r Resolution) Raw() float64 {
	return r.Value
}

// ═══════════════════════════════════════════════════════════════
//  Arithmetic Operations
// ═══════════════════════════════════════════════════════════════

// Add adds two resolution values with the same unit.
// Panics if the units are different.
func (r Resolution) Add(other Resolution) Resolution {
	if r.Unit != other.Unit {
		panic(fmt.Sprintf("cannot add resolution values with different units: %s + %s", r.Unit, other.Unit))
	}
	return Resolution{Value: r.Value + other.Value, Unit: r.Unit}
}

// Sub subtracts another resolution value with the same unit.
// Panics if the units are different.
func (r Resolution) Sub(other Resolution) Resolution {
	if r.Unit != other.Unit {
		panic(fmt.Sprintf("cannot subtract resolution values with different units: %s - %s", r.Unit, other.Unit))
	}
	return Resolution{Value: r.Value - other.Value, Unit: r.Unit}
}

// Mul multiplies the resolution by a scalar value.
func (r Resolution) Mul(scalar float64) Resolution {
	return Resolution{Value: r.Value * scalar, Unit: r.Unit}
}

// Div divides the resolution by a scalar value.
func (r Resolution) Div(scalar float64) Resolution {
	return Resolution{Value: r.Value / scalar, Unit: r.Unit}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// LessThan returns true if this resolution is less than the other.
// Panics if the units are different.
func (r Resolution) LessThan(other Resolution) bool {
	if r.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare resolution values with different units: %s < %s", r.Unit, other.Unit))
	}
	return r.Value < other.Value
}

// GreaterThan returns true if this resolution is greater than the other.
// Panics if the units are different.
func (r Resolution) GreaterThan(other Resolution) bool {
	if r.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare resolution values with different units: %s > %s", r.Unit, other.Unit))
	}
	return r.Value > other.Value
}

// ═══════════════════════════════════════════════════════════════
//  Conversions
// ═══════════════════════════════════════════════════════════════

const (
	dpiPerDppx       = 96.0
	cmPerInchRes     = 2.54
	dpiPerDpcm       = dpiPerDppx / cmPerInchRes // ≈ 37.795275591
)

// ToDpi converts any resolution value to dots per inch.
//
// Example:
//
//	res := units.Dppx(2)
//	dpi := res.ToDpi()  // Returns units.Dpi(192)
func (r Resolution) ToDpi() Resolution {
	switch r.Unit {
	case DotsPerInch:
		return r
	case DotsPerCentimeter:
		return Dpi(r.Value * cmPerInchRes)
	case DotsPerPixel:
		return Dpi(r.Value * dpiPerDppx)
	default:
		panic(fmt.Sprintf("unknown resolution unit: %s", r.Unit))
	}
}

// ToDpcm converts any resolution value to dots per centimeter.
//
// Example:
//
//	res := units.Dpi(96)
//	dpcm := res.ToDpcm()  // Returns units.Dpcm(37.795)
func (r Resolution) ToDpcm() Resolution {
	switch r.Unit {
	case DotsPerInch:
		return Dpcm(r.Value / cmPerInchRes)
	case DotsPerCentimeter:
		return r
	case DotsPerPixel:
		return Dpcm(r.Value * dpiPerDpcm)
	default:
		panic(fmt.Sprintf("unknown resolution unit: %s", r.Unit))
	}
}

// ToDppx converts any resolution value to dots per pixel unit.
//
// Example:
//
//	res := units.Dpi(192)
//	dppx := res.ToDppx()  // Returns units.Dppx(2)
func (r Resolution) ToDppx() Resolution {
	switch r.Unit {
	case DotsPerInch:
		return Dppx(r.Value / dpiPerDppx)
	case DotsPerCentimeter:
		return Dppx(r.Value / dpiPerDpcm)
	case DotsPerPixel:
		return r
	default:
		panic(fmt.Sprintf("unknown resolution unit: %s", r.Unit))
	}
}

// To converts the resolution to another unit.
//
// Example:
//
//	res := units.Dpi(96)
//	dppx := res.To(units.DotsPerPixel)  // Returns units.Dppx(1)
func (r Resolution) To(targetUnit ResolutionUnit) Resolution {
	switch targetUnit {
	case DotsPerInch:
		return r.ToDpi()
	case DotsPerCentimeter:
		return r.ToDpcm()
	case DotsPerPixel:
		return r.ToDppx()
	default:
		panic(fmt.Sprintf("unknown target resolution unit: %s", targetUnit))
	}
}

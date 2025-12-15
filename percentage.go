package units

import "fmt"

// ═══════════════════════════════════════════════════════════════
//  Percentage - CSS Values Level 4 Section 5.5
// ═══════════════════════════════════════════════════════════════

// Percentage represents a CSS percentage value.
//
// Percentage values are always relative to another quantity. The meaning
// of 100% depends on the context in which the percentage is used.
// For example, width: 50% means 50% of the containing block's width.
//
// References:
//   - CSS Values Level 4 - Percentages: https://www.w3.org/TR/css-values-4/#percentages
//   - MDN - CSS percentage: https://developer.mozilla.org/en-US/docs/Web/CSS/percentage
type Percentage struct {
	Value float64
}

// ═══════════════════════════════════════════════════════════════
//  Constructor
// ═══════════════════════════════════════════════════════════════

// Percent creates a percentage value.
//
// Example:
//
//	width := units.Percent(50)  // 50%
func Percent(value float64) Percentage {
	return Percentage{Value: value}
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the percentage.
func (p Percentage) String() string {
	return fmt.Sprintf("%.2f%%", p.Value)
}

// IsZero returns true if the percentage value is zero.
func (p Percentage) IsZero() bool {
	return p.Value == 0
}

// Raw returns the raw numeric value of the percentage.
func (p Percentage) Raw() float64 {
	return p.Value
}

// Fraction returns the percentage as a fraction (e.g., 50% -> 0.5).
func (p Percentage) Fraction() float64 {
	return p.Value / 100.0
}

// Of calculates the percentage of a given value.
//
// Example:
//
//	pct := units.Percent(50)
//	result := pct.Of(200)  // Returns 100 (50% of 200)
func (p Percentage) Of(value float64) float64 {
	return value * p.Fraction()
}

// ═══════════════════════════════════════════════════════════════
//  Arithmetic Operations
// ═══════════════════════════════════════════════════════════════

// Add adds two percentage values.
func (p Percentage) Add(other Percentage) Percentage {
	return Percentage{Value: p.Value + other.Value}
}

// Sub subtracts another percentage value.
func (p Percentage) Sub(other Percentage) Percentage {
	return Percentage{Value: p.Value - other.Value}
}

// Mul multiplies the percentage by a scalar value.
func (p Percentage) Mul(scalar float64) Percentage {
	return Percentage{Value: p.Value * scalar}
}

// Div divides the percentage by a scalar value.
func (p Percentage) Div(scalar float64) Percentage {
	return Percentage{Value: p.Value / scalar}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// LessThan returns true if this percentage is less than the other.
func (p Percentage) LessThan(other Percentage) bool {
	return p.Value < other.Value
}

// GreaterThan returns true if this percentage is greater than the other.
func (p Percentage) GreaterThan(other Percentage) bool {
	return p.Value > other.Value
}

// Equals returns true if this percentage equals the other.
func (p Percentage) Equals(other Percentage) bool {
	return p.Value == other.Value
}

// Clamp returns a percentage clamped to the given range.
//
// Example:
//
//	pct := units.Percent(150)
//	clamped := pct.Clamp(0, 100)  // Returns units.Percent(100)
func (p Percentage) Clamp(min, max float64) Percentage {
	value := p.Value
	if value < min {
		value = min
	}
	if value > max {
		value = max
	}
	return Percentage{Value: value}
}

package units

import (
	"fmt"
	"math"
)

// ═══════════════════════════════════════════════════════════════
//  Number - CSS Values Level 4 Section 5.3
// ═══════════════════════════════════════════════════════════════

// Number represents a CSS number value.
//
// Number values are dimensionless real numbers, possibly with a fractional
// component. They're used in properties like line-height, opacity, flex-grow,
// and as multipliers in various calculations.
//
// References:
//   - CSS Values Level 4 - Numbers: https://www.w3.org/TR/css-values-4/#numbers
//   - MDN - CSS number: https://developer.mozilla.org/en-US/docs/Web/CSS/number
type Number struct {
	Value float64
}

// ═══════════════════════════════════════════════════════════════
//  Constructor
// ═══════════════════════════════════════════════════════════════

// Num creates a number value.
//
// Example:
//
//	lineHeight := units.Num(1.5)  // line-height: 1.5
//	opacity := units.Num(0.8)     // opacity: 0.8
func Num(value float64) Number {
	return Number{Value: value}
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the number.
func (n Number) String() string {
	// Use appropriate precision based on value
	if n.Value == float64(int64(n.Value)) {
		return fmt.Sprintf("%.0f", n.Value)
	}
	return fmt.Sprintf("%.6g", n.Value)
}

// IsZero returns true if the number value is zero.
func (n Number) IsZero() bool {
	return n.Value == 0
}

// IsPositive returns true if the number is greater than zero.
func (n Number) IsPositive() bool {
	return n.Value > 0
}

// IsNegative returns true if the number is less than zero.
func (n Number) IsNegative() bool {
	return n.Value < 0
}

// Raw returns the raw numeric value.
func (n Number) Raw() float64 {
	return n.Value
}

// Abs returns the absolute value of the number.
func (n Number) Abs() Number {
	return Number{Value: math.Abs(n.Value)}
}

// ═══════════════════════════════════════════════════════════════
//  Arithmetic Operations
// ═══════════════════════════════════════════════════════════════

// Add adds two numbers.
func (n Number) Add(other Number) Number {
	return Number{Value: n.Value + other.Value}
}

// Sub subtracts another number.
func (n Number) Sub(other Number) Number {
	return Number{Value: n.Value - other.Value}
}

// Mul multiplies two numbers.
func (n Number) Mul(other Number) Number {
	return Number{Value: n.Value * other.Value}
}

// Div divides by another number.
func (n Number) Div(other Number) Number {
	return Number{Value: n.Value / other.Value}
}

// Pow raises the number to the given power.
func (n Number) Pow(exponent float64) Number {
	return Number{Value: math.Pow(n.Value, exponent)}
}

// Sqrt returns the square root of the number.
func (n Number) Sqrt() Number {
	return Number{Value: math.Sqrt(n.Value)}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// LessThan returns true if this number is less than the other.
func (n Number) LessThan(other Number) bool {
	return n.Value < other.Value
}

// GreaterThan returns true if this number is greater than the other.
func (n Number) GreaterThan(other Number) bool {
	return n.Value > other.Value
}

// Equals returns true if this number equals the other.
func (n Number) Equals(other Number) bool {
	return n.Value == other.Value
}

// ═══════════════════════════════════════════════════════════════
//  Utility Methods
// ═══════════════════════════════════════════════════════════════

// Clamp returns a number clamped to the given range.
//
// Example:
//
//	num := units.Num(1.8)
//	clamped := num.Clamp(0, 1)  // Returns units.Num(1)
func (n Number) Clamp(minValue, maxValue float64) Number {
	value := n.Value
	if value < minValue {
		value = minValue
	}
	if value > maxValue {
		value = maxValue
	}
	return Number{Value: value}
}

// Round returns the number rounded to the nearest integer.
func (n Number) Round() Number {
	return Number{Value: math.Round(n.Value)}
}

// Floor returns the largest integer less than or equal to the number.
func (n Number) Floor() Number {
	return Number{Value: math.Floor(n.Value)}
}

// Ceil returns the smallest integer greater than or equal to the number.
func (n Number) Ceil() Number {
	return Number{Value: math.Ceil(n.Value)}
}

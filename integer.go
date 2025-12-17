package units

import (
	"fmt"
	"math"
)

// ═══════════════════════════════════════════════════════════════
//  Integer - CSS Values Level 4 Section 5.2
// ═══════════════════════════════════════════════════════════════

// Integer represents a CSS integer value.
//
// Integer values are whole numbers (no fractional component). They're used
// in properties like z-index, column-count, grid-row-start, and anywhere
// a count or index is needed.
//
// References:
//   - CSS Values Level 4 - Integers: https://www.w3.org/TR/css-values-4/#integers
//   - MDN - CSS integer: https://developer.mozilla.org/en-US/docs/Web/CSS/integer
type Integer struct {
	Value int64
}

// ═══════════════════════════════════════════════════════════════
//  Constructor
// ═══════════════════════════════════════════════════════════════

// Int creates an integer value.
//
// Example:
//
//	zIndex := units.Int(10)      // z-index: 10
//	columns := units.Int(3)      // column-count: 3
func Int(value int64) Integer {
	return Integer{Value: value}
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the integer.
func (i Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

// IsZero returns true if the integer value is zero.
func (i Integer) IsZero() bool {
	return i.Value == 0
}

// IsPositive returns true if the integer is greater than zero.
func (i Integer) IsPositive() bool {
	return i.Value > 0
}

// IsNegative returns true if the integer is less than zero.
func (i Integer) IsNegative() bool {
	return i.Value < 0
}

// IsEven returns true if the integer is even.
func (i Integer) IsEven() bool {
	return i.Value%2 == 0
}

// IsOdd returns true if the integer is odd.
func (i Integer) IsOdd() bool {
	return i.Value%2 != 0
}

// Raw returns the raw integer value.
func (i Integer) Raw() int64 {
	return i.Value
}

// Float returns the integer as a float64.
func (i Integer) Float() float64 {
	return float64(i.Value)
}

// Abs returns the absolute value of the integer.
func (i Integer) Abs() Integer {
	if i.Value < 0 {
		return Integer{Value: -i.Value}
	}
	return i
}

// ═══════════════════════════════════════════════════════════════
//  Arithmetic Operations
// ═══════════════════════════════════════════════════════════════

// Add adds two integers.
func (i Integer) Add(other Integer) Integer {
	return Integer{Value: i.Value + other.Value}
}

// Sub subtracts another integer.
func (i Integer) Sub(other Integer) Integer {
	return Integer{Value: i.Value - other.Value}
}

// Mul multiplies two integers.
func (i Integer) Mul(other Integer) Integer {
	return Integer{Value: i.Value * other.Value}
}

// Div divides by another integer (integer division, truncates toward zero).
func (i Integer) Div(other Integer) Integer {
	return Integer{Value: i.Value / other.Value}
}

// Mod returns the remainder of division (modulo operation).
func (i Integer) Mod(other Integer) Integer {
	return Integer{Value: i.Value % other.Value}
}

// Pow raises the integer to the given integer power.
func (i Integer) Pow(exponent int64) Integer {
	return Integer{Value: int64(math.Pow(float64(i.Value), float64(exponent)))}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// LessThan returns true if this integer is less than the other.
func (i Integer) LessThan(other Integer) bool {
	return i.Value < other.Value
}

// GreaterThan returns true if this integer is greater than the other.
func (i Integer) GreaterThan(other Integer) bool {
	return i.Value > other.Value
}

// Equals returns true if this integer equals the other.
func (i Integer) Equals(other Integer) bool {
	return i.Value == other.Value
}

// ═══════════════════════════════════════════════════════════════
//  Utility Methods
// ═══════════════════════════════════════════════════════════════

// Clamp returns an integer clamped to the given range.
//
// Example:
//
//	num := units.Int(150)
//	clamped := num.Clamp(0, 100)  // Returns units.Int(100)
func (i Integer) Clamp(min, max int64) Integer {
	value := i.Value
	if value < min {
		value = min
	}
	if value > max {
		value = max
	}
	return Integer{Value: value}
}

// Minimum returns the smaller of two integers.
func (i Integer) Minimum(other Integer) Integer {
	if i.Value < other.Value {
		return i
	}
	return other
}

// Maximum returns the larger of two integers.
func (i Integer) Maximum(other Integer) Integer {
	if i.Value > other.Value {
		return i
	}
	return other
}

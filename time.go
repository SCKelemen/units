package units

import "fmt"

// ═══════════════════════════════════════════════════════════════
//  Time Units - CSS Values Level 4 Section 6.2
// ═══════════════════════════════════════════════════════════════

// Time represents a CSS time value.
//
// Time values are used in CSS animations, transitions, and other
// time-based properties. CSS supports seconds and milliseconds.
//
// References:
//   - CSS Values Level 4 - Time: https://www.w3.org/TR/css-values-4/#time
//   - MDN - CSS time: https://developer.mozilla.org/en-US/docs/Web/CSS/time
//   - web.dev - CSS animations: https://web.dev/learn/css/animations/
type Time struct {
	Value float64
	Unit  TimeUnit
}

// TimeUnit represents a CSS time unit type.
type TimeUnit string

// CSS time units per CSS Values Level 4.
//
// Relationship:
//   - 1s = 1000ms
const (
	Second      TimeUnit = "s"  // Seconds
	Millisecond TimeUnit = "ms" // Milliseconds
)

// ═══════════════════════════════════════════════════════════════
//  Constructors
// ═══════════════════════════════════════════════════════════════

// Sec creates a time value in seconds.
//
// Example:
//
//	duration := units.Sec(2.5)  // 2.5 seconds
func Sec(value float64) Time {
	return Time{Value: value, Unit: Second}
}

// Ms creates a time value in milliseconds.
//
// Example:
//
//	duration := units.Ms(500)  // 500 milliseconds
func Ms(value float64) Time {
	return Time{Value: value, Unit: Millisecond}
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the time.
func (t Time) String() string {
	return fmt.Sprintf("%.2f%s", t.Value, t.Unit)
}

// IsZero returns true if the time value is zero.
func (t Time) IsZero() bool {
	return t.Value == 0
}

// Raw returns the raw numeric value of the time.
func (t Time) Raw() float64 {
	return t.Value
}

// ═══════════════════════════════════════════════════════════════
//  Arithmetic Operations
// ═══════════════════════════════════════════════════════════════

// Add adds two time values with the same unit.
// Panics if the units are different.
func (t Time) Add(other Time) Time {
	if t.Unit != other.Unit {
		panic(fmt.Sprintf("cannot add time values with different units: %s + %s", t.Unit, other.Unit))
	}
	return Time{Value: t.Value + other.Value, Unit: t.Unit}
}

// Sub subtracts another time value with the same unit.
// Panics if the units are different.
func (t Time) Sub(other Time) Time {
	if t.Unit != other.Unit {
		panic(fmt.Sprintf("cannot subtract time values with different units: %s - %s", t.Unit, other.Unit))
	}
	return Time{Value: t.Value - other.Value, Unit: t.Unit}
}

// Mul multiplies the time by a scalar value.
func (t Time) Mul(scalar float64) Time {
	return Time{Value: t.Value * scalar, Unit: t.Unit}
}

// Div divides the time by a scalar value.
func (t Time) Div(scalar float64) Time {
	return Time{Value: t.Value / scalar, Unit: t.Unit}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// LessThan returns true if this time is less than the other.
// Panics if the units are different.
func (t Time) LessThan(other Time) bool {
	if t.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare time values with different units: %s < %s", t.Unit, other.Unit))
	}
	return t.Value < other.Value
}

// GreaterThan returns true if this time is greater than the other.
// Panics if the units are different.
func (t Time) GreaterThan(other Time) bool {
	if t.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare time values with different units: %s > %s", t.Unit, other.Unit))
	}
	return t.Value > other.Value
}

// ═══════════════════════════════════════════════════════════════
//  Conversions
// ═══════════════════════════════════════════════════════════════

const msPerSecond = 1000.0

// ToSec converts any time value to seconds.
//
// Example:
//
//	time := units.Ms(2500)
//	sec := time.ToSec()  // Returns units.Sec(2.5)
func (t Time) ToSec() Time {
	switch t.Unit {
	case Second:
		return t
	case Millisecond:
		return Sec(t.Value / msPerSecond)
	default:
		panic(fmt.Sprintf("unknown time unit: %s", t.Unit))
	}
}

// ToMs converts any time value to milliseconds.
//
// Example:
//
//	time := units.Sec(2.5)
//	ms := time.ToMs()  // Returns units.Ms(2500)
func (t Time) ToMs() Time {
	switch t.Unit {
	case Second:
		return Ms(t.Value * msPerSecond)
	case Millisecond:
		return t
	default:
		panic(fmt.Sprintf("unknown time unit: %s", t.Unit))
	}
}

// To converts the time to another unit.
//
// Example:
//
//	time := units.Sec(1.5)
//	ms := time.To(units.Millisecond)  // Returns units.Ms(1500)
func (t Time) To(targetUnit TimeUnit) Time {
	switch targetUnit {
	case Second:
		return t.ToSec()
	case Millisecond:
		return t.ToMs()
	default:
		panic(fmt.Sprintf("unknown target time unit: %s", targetUnit))
	}
}

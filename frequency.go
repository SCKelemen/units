package units

import "fmt"

// ═══════════════════════════════════════════════════════════════
//  Frequency Units - CSS Values Level 4 Section 6.3
// ═══════════════════════════════════════════════════════════════

// Frequency represents a CSS frequency value.
//
// Frequency values are used in CSS for audio properties and
// other frequency-based features. CSS supports hertz and kilohertz.
//
// References:
//   - CSS Values Level 4 - Frequency: https://www.w3.org/TR/css-values-4/#frequency
//   - MDN - CSS frequency: https://developer.mozilla.org/en-US/docs/Web/CSS/frequency
type Frequency struct {
	Value float64
	Unit  FrequencyUnit
}

// FrequencyUnit represents a CSS frequency unit type.
type FrequencyUnit string

// CSS frequency units per CSS Values Level 4.
//
// Relationship:
//   - 1kHz = 1000Hz
const (
	Hertz     FrequencyUnit = "Hz"  // Hertz
	Kilohertz FrequencyUnit = "kHz" // Kilohertz
)

// ═══════════════════════════════════════════════════════════════
//  Constructors
// ═══════════════════════════════════════════════════════════════

// Hz creates a frequency value in hertz.
//
// Example:
//
//	freq := units.Hz(440)  // 440 Hz (A4 note)
func Hz(value float64) Frequency {
	return Frequency{Value: value, Unit: Hertz}
}

// KHz creates a frequency value in kilohertz.
//
// Example:
//
//	freq := units.KHz(20)  // 20 kHz
func KHz(value float64) Frequency {
	return Frequency{Value: value, Unit: Kilohertz}
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the frequency.
func (f Frequency) String() string {
	return fmt.Sprintf("%.2f%s", f.Value, f.Unit)
}

// IsZero returns true if the frequency value is zero.
func (f Frequency) IsZero() bool {
	return f.Value == 0
}

// Raw returns the raw numeric value of the frequency.
func (f Frequency) Raw() float64 {
	return f.Value
}

// ═══════════════════════════════════════════════════════════════
//  Arithmetic Operations
// ═══════════════════════════════════════════════════════════════

// Add adds two frequency values with the same unit.
// Panics if the units are different.
func (f Frequency) Add(other Frequency) Frequency {
	if f.Unit != other.Unit {
		panic(fmt.Sprintf("cannot add frequency values with different units: %s + %s", f.Unit, other.Unit))
	}
	return Frequency{Value: f.Value + other.Value, Unit: f.Unit}
}

// Sub subtracts another frequency value with the same unit.
// Panics if the units are different.
func (f Frequency) Sub(other Frequency) Frequency {
	if f.Unit != other.Unit {
		panic(fmt.Sprintf("cannot subtract frequency values with different units: %s - %s", f.Unit, other.Unit))
	}
	return Frequency{Value: f.Value - other.Value, Unit: f.Unit}
}

// Mul multiplies the frequency by a scalar value.
func (f Frequency) Mul(scalar float64) Frequency {
	return Frequency{Value: f.Value * scalar, Unit: f.Unit}
}

// Div divides the frequency by a scalar value.
func (f Frequency) Div(scalar float64) Frequency {
	return Frequency{Value: f.Value / scalar, Unit: f.Unit}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// LessThan returns true if this frequency is less than the other.
// Panics if the units are different.
func (f Frequency) LessThan(other Frequency) bool {
	if f.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare frequency values with different units: %s < %s", f.Unit, other.Unit))
	}
	return f.Value < other.Value
}

// GreaterThan returns true if this frequency is greater than the other.
// Panics if the units are different.
func (f Frequency) GreaterThan(other Frequency) bool {
	if f.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare frequency values with different units: %s > %s", f.Unit, other.Unit))
	}
	return f.Value > other.Value
}

// ═══════════════════════════════════════════════════════════════
//  Conversions
// ═══════════════════════════════════════════════════════════════

const hzPerKHz = 1000.0

// ToHz converts any frequency value to hertz.
//
// Example:
//
//	freq := units.KHz(2.5)
//	hz := freq.ToHz()  // Returns units.Hz(2500)
func (f Frequency) ToHz() Frequency {
	switch f.Unit {
	case Hertz:
		return f
	case Kilohertz:
		return Hz(f.Value * hzPerKHz)
	default:
		panic(fmt.Sprintf("unknown frequency unit: %s", f.Unit))
	}
}

// ToKHz converts any frequency value to kilohertz.
//
// Example:
//
//	freq := units.Hz(2500)
//	khz := freq.ToKHz()  // Returns units.KHz(2.5)
func (f Frequency) ToKHz() Frequency {
	switch f.Unit {
	case Hertz:
		return KHz(f.Value / hzPerKHz)
	case Kilohertz:
		return f
	default:
		panic(fmt.Sprintf("unknown frequency unit: %s", f.Unit))
	}
}

// To converts the frequency to another unit.
//
// Example:
//
//	freq := units.Hz(1500)
//	khz := freq.To(units.Kilohertz)  // Returns units.KHz(1.5)
func (f Frequency) To(targetUnit FrequencyUnit) Frequency {
	switch targetUnit {
	case Hertz:
		return f.ToHz()
	case Kilohertz:
		return f.ToKHz()
	default:
		panic(fmt.Sprintf("unknown target frequency unit: %s", targetUnit))
	}
}

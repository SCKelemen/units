package units

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ═══════════════════════════════════════════════════════════════
//  Angle Units - CSS Values Level 4 Section 6.1
// ═══════════════════════════════════════════════════════════════

// Angle represents a CSS angle value.
//
// Angles are used in various CSS properties like transforms, gradients,
// and animations. CSS supports four angle units: degrees, gradians,
// radians, and turns.
//
// References:
//   - CSS Values Level 4 - Angles: https://www.w3.org/TR/css-values-4/#angles
//   - MDN - CSS angle: https://developer.mozilla.org/en-US/docs/Web/CSS/angle
//   - web.dev - CSS transforms: https://web.dev/learn/css/transforms/
type Angle struct {
	Value float64
	Unit  AngleUnit
}

// AngleUnit represents a CSS angle unit type.
type AngleUnit string

// CSS angle units per CSS Values Level 4.
//
// Relationship:
//   - 1turn = 360deg = 400grad = 2π rad
//   - 1deg = 1/360 turn
//   - 1grad = 1/400 turn
//   - 1rad = 1/(2π) turn
const (
	Degree  AngleUnit = "deg"  // Degrees: 360deg = full circle
	Gradian AngleUnit = "grad" // Gradians: 400grad = full circle
	Radian  AngleUnit = "rad"  // Radians: 2π rad = full circle
	Turn    AngleUnit = "turn" // Turns: 1turn = full circle
)

// ═══════════════════════════════════════════════════════════════
//  Constructors
// ═══════════════════════════════════════════════════════════════

// Deg creates an angle value in degrees.
//
// Example:
//
//	angle := units.Deg(45)  // 45 degrees
func Deg(value float64) Angle {
	return Angle{Value: value, Unit: Degree}
}

// Grad creates an angle value in gradians.
//
// Example:
//
//	angle := units.Grad(200)  // 200 gradians (half circle)
func Grad(value float64) Angle {
	return Angle{Value: value, Unit: Gradian}
}

// Rad creates an angle value in radians.
//
// Example:
//
//	angle := units.Rad(math.Pi)  // π radians (half circle)
func Rad(value float64) Angle {
	return Angle{Value: value, Unit: Radian}
}

// Turns creates an angle value in turns.
//
// Example:
//
//	angle := units.Turns(0.25)  // 0.25 turns (quarter circle)
func Turns(value float64) Angle {
	return Angle{Value: value, Unit: Turn}
}

// ParseAngle parses a CSS angle string into an Angle.
//
// Supported formats:
//   - "<number>deg" (degrees)
//   - "<number>grad" (gradians)
//   - "<number>rad" (radians)
//   - "<number>turn" (turns)
//   - "<number>" (defaults to degrees)
//
// The parser is case-insensitive and accepts optional whitespace.
//
// Examples:
//
//	angle, _ := units.ParseAngle("180deg")
//	angle, _ := units.ParseAngle("0.5turn")
//	angle, _ := units.ParseAngle("3.14159rad")
//	angle, _ := units.ParseAngle("200grad")
//	angle, _ := units.ParseAngle("90") // defaults to degrees
func ParseAngle(input string) (Angle, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return Angle{}, fmt.Errorf("invalid angle %q: empty value", input)
	}

	// Split "<number><unit>" where unit is optional trailing letters.
	// This also allows optional whitespace between number and unit.
	unitStart := len(s)
	for unitStart > 0 {
		ch := s[unitStart-1]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			unitStart--
			continue
		}
		break
	}

	numberPart := strings.TrimSpace(s[:unitStart])
	unitPart := strings.ToLower(strings.TrimSpace(s[unitStart:]))
	if numberPart == "" {
		return Angle{}, fmt.Errorf("invalid angle %q: missing numeric value", input)
	}

	value, err := strconv.ParseFloat(numberPart, 64)
	if err != nil {
		return Angle{}, fmt.Errorf("invalid angle %q: %w", input, err)
	}
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return Angle{}, fmt.Errorf("invalid angle %q: value must be finite", input)
	}

	switch unitPart {
	case "":
		return Deg(value), nil
	case string(Degree):
		return Deg(value), nil
	case string(Gradian):
		return Grad(value), nil
	case string(Radian):
		return Rad(value), nil
	case string(Turn):
		return Turns(value), nil
	default:
		return Angle{}, fmt.Errorf("invalid angle %q: unsupported unit %q", input, unitPart)
	}
}

// MustParseAngle parses a CSS angle string and panics if parsing fails.
func MustParseAngle(input string) Angle {
	angle, err := ParseAngle(input)
	if err != nil {
		panic(err)
	}
	return angle
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the angle.
func (a Angle) String() string {
	return fmt.Sprintf("%.2f%s", a.Value, a.Unit)
}

// IsZero returns true if the angle value is zero.
func (a Angle) IsZero() bool {
	return a.Value == 0
}

// Raw returns the raw numeric value of the angle.
func (a Angle) Raw() float64 {
	return a.Value
}

// ═══════════════════════════════════════════════════════════════
//  Arithmetic Operations
// ═══════════════════════════════════════════════════════════════

// Add adds two angles with the same unit.
// Panics if the units are different.
func (a Angle) Add(other Angle) Angle {
	if a.Unit != other.Unit {
		panic(fmt.Sprintf("cannot add angles with different units: %s + %s", a.Unit, other.Unit))
	}
	return Angle{Value: a.Value + other.Value, Unit: a.Unit}
}

// Sub subtracts another angle with the same unit.
// Panics if the units are different.
func (a Angle) Sub(other Angle) Angle {
	if a.Unit != other.Unit {
		panic(fmt.Sprintf("cannot subtract angles with different units: %s - %s", a.Unit, other.Unit))
	}
	return Angle{Value: a.Value - other.Value, Unit: a.Unit}
}

// Mul multiplies the angle by a scalar value.
func (a Angle) Mul(scalar float64) Angle {
	return Angle{Value: a.Value * scalar, Unit: a.Unit}
}

// Div divides the angle by a scalar value.
func (a Angle) Div(scalar float64) Angle {
	return Angle{Value: a.Value / scalar, Unit: a.Unit}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// LessThan returns true if this angle is less than the other.
// Panics if the units are different.
func (a Angle) LessThan(other Angle) bool {
	if a.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare angles with different units: %s < %s", a.Unit, other.Unit))
	}
	return a.Value < other.Value
}

// GreaterThan returns true if this angle is greater than the other.
// Panics if the units are different.
func (a Angle) GreaterThan(other Angle) bool {
	if a.Unit != other.Unit {
		panic(fmt.Sprintf("cannot compare angles with different units: %s > %s", a.Unit, other.Unit))
	}
	return a.Value > other.Value
}

// ═══════════════════════════════════════════════════════════════
//  Conversions
// ═══════════════════════════════════════════════════════════════

// Conversion constants
const (
	degreesPerTurn   = 360.0
	gradiansPerTurn  = 400.0
	radiansPerTurn   = 2 * math.Pi
	degreesPerRadian = 180.0 / math.Pi
	radiansPerDegree = math.Pi / 180.0
)

// ToDeg converts any angle to degrees.
//
// Example:
//
//	angle := units.Turns(0.5)
//	deg := angle.ToDeg()  // Returns units.Deg(180)
func (a Angle) ToDeg() Angle {
	var degValue float64
	switch a.Unit {
	case Degree:
		degValue = a.Value
	case Gradian:
		degValue = a.Value * degreesPerTurn / gradiansPerTurn
	case Radian:
		degValue = a.Value * degreesPerRadian
	case Turn:
		degValue = a.Value * degreesPerTurn
	default:
		panic(fmt.Sprintf("unknown angle unit: %s", a.Unit))
	}
	return Deg(degValue)
}

// ToRad converts any angle to radians.
//
// Example:
//
//	angle := units.Deg(180)
//	rad := angle.ToRad()  // Returns units.Rad(π)
func (a Angle) ToRad() Angle {
	var radValue float64
	switch a.Unit {
	case Degree:
		radValue = a.Value * radiansPerDegree
	case Gradian:
		radValue = a.Value * radiansPerTurn / gradiansPerTurn
	case Radian:
		radValue = a.Value
	case Turn:
		radValue = a.Value * radiansPerTurn
	default:
		panic(fmt.Sprintf("unknown angle unit: %s", a.Unit))
	}
	return Rad(radValue)
}

// ToGrad converts any angle to gradians.
//
// Example:
//
//	angle := units.Deg(180)
//	grad := angle.ToGrad()  // Returns units.Grad(200)
func (a Angle) ToGrad() Angle {
	var gradValue float64
	switch a.Unit {
	case Degree:
		gradValue = a.Value * gradiansPerTurn / degreesPerTurn
	case Gradian:
		gradValue = a.Value
	case Radian:
		gradValue = a.Value * gradiansPerTurn / radiansPerTurn
	case Turn:
		gradValue = a.Value * gradiansPerTurn
	default:
		panic(fmt.Sprintf("unknown angle unit: %s", a.Unit))
	}
	return Grad(gradValue)
}

// ToTurns converts any angle to turns.
//
// Example:
//
//	angle := units.Deg(180)
//	turns := angle.ToTurns()  // Returns units.Turns(0.5)
func (a Angle) ToTurns() Angle {
	var turnValue float64
	switch a.Unit {
	case Degree:
		turnValue = a.Value / degreesPerTurn
	case Gradian:
		turnValue = a.Value / gradiansPerTurn
	case Radian:
		turnValue = a.Value / radiansPerTurn
	case Turn:
		turnValue = a.Value
	default:
		panic(fmt.Sprintf("unknown angle unit: %s", a.Unit))
	}
	return Turns(turnValue)
}

// To converts the angle to another unit.
//
// Example:
//
//	angle := units.Deg(90)
//	rad := angle.To(units.Radian)  // Returns units.Rad(π/2)
func (a Angle) To(targetUnit AngleUnit) Angle {
	switch targetUnit {
	case Degree:
		return a.ToDeg()
	case Gradian:
		return a.ToGrad()
	case Radian:
		return a.ToRad()
	case Turn:
		return a.ToTurns()
	default:
		panic(fmt.Sprintf("unknown target angle unit: %s", targetUnit))
	}
}

// ═══════════════════════════════════════════════════════════════
//  Normalization
// ═══════════════════════════════════════════════════════════════

// Normalize returns an equivalent angle in the range [0, full circle).
// This is useful for angles that may have wrapped around multiple times.
//
// Example:
//
//	angle := units.Deg(450)
//	norm := angle.Normalize()  // Returns units.Deg(90)
func (a Angle) Normalize() Angle {
	var fullCircle float64
	switch a.Unit {
	case Degree:
		fullCircle = degreesPerTurn
	case Gradian:
		fullCircle = gradiansPerTurn
	case Radian:
		fullCircle = radiansPerTurn
	case Turn:
		fullCircle = 1.0
	default:
		panic(fmt.Sprintf("unknown angle unit: %s", a.Unit))
	}

	normalized := math.Mod(a.Value, fullCircle)
	if normalized < 0 {
		normalized += fullCircle
	}

	return Angle{Value: normalized, Unit: a.Unit}
}

package units

import (
	"fmt"
	"math"
)

// ═══════════════════════════════════════════════════════════════
//  Ratio - CSS Values Level 4 Section 5.7
// ═══════════════════════════════════════════════════════════════

// Ratio represents a CSS ratio value.
//
// Ratio values consist of two positive numbers separated by a slash (/),
// representing the ratio of the first number to the second. Common uses
// include aspect-ratio property for maintaining proportions.
//
// If the second number is omitted, it defaults to 1.
//
// References:
//   - CSS Values Level 4 - Ratios: https://www.w3.org/TR/css-values-4/#ratios
//   - MDN - CSS ratio: https://developer.mozilla.org/en-US/docs/Web/CSS/ratio
//   - MDN - aspect-ratio: https://developer.mozilla.org/en-US/docs/Web/CSS/aspect-ratio
type Ratio struct {
	First  float64
	Second float64
}

// ═══════════════════════════════════════════════════════════════
//  Constructors
// ═══════════════════════════════════════════════════════════════

// NewRatio creates a ratio value from two numbers.
//
// Example:
//
//	aspectRatio := units.NewRatio(16, 9)  // 16:9 aspect ratio
//	square := units.NewRatio(1, 1)        // 1:1 aspect ratio
func NewRatio(first, second float64) Ratio {
	if first <= 0 || second <= 0 {
		panic(fmt.Sprintf("ratio values must be positive: got %f/%f", first, second))
	}
	return Ratio{First: first, Second: second}
}

// RatioFrom creates a ratio from a single number (equivalent to n/1).
//
// Example:
//
//	ratio := units.RatioFrom(2)  // 2/1 ratio
func RatioFrom(value float64) Ratio {
	if value <= 0 {
		panic(fmt.Sprintf("ratio value must be positive: got %f", value))
	}
	return Ratio{First: value, Second: 1}
}

// ═══════════════════════════════════════════════════════════════
//  Methods
// ═══════════════════════════════════════════════════════════════

// String returns the string representation of the ratio.
func (r Ratio) String() string {
	// If second is 1, can use simplified notation
	if r.Second == 1.0 {
		return fmt.Sprintf("%.6g", r.First)
	}
	return fmt.Sprintf("%.6g / %.6g", r.First, r.Second)
}

// Value returns the ratio as a decimal number (first / second).
//
// Example:
//
//	ratio := units.NewRatio(16, 9)
//	value := ratio.Value()  // Returns ~1.778
func (r Ratio) Value() float64 {
	return r.First / r.Second
}

// IsSquare returns true if the ratio is 1:1.
func (r Ratio) IsSquare() bool {
	return r.First == r.Second
}

// IsWide returns true if the ratio is wider than it is tall (first > second).
func (r Ratio) IsWide() bool {
	return r.First > r.Second
}

// IsTall returns true if the ratio is taller than it is wide (first < second).
func (r Ratio) IsTall() bool {
	return r.First < r.Second
}

// Inverse returns the inverse of the ratio (swaps first and second).
//
// Example:
//
//	ratio := units.NewRatio(16, 9)
//	inverse := ratio.Inverse()  // Returns 9/16
func (r Ratio) Inverse() Ratio {
	return Ratio{First: r.Second, Second: r.First}
}

// Simplify returns a simplified version of the ratio by dividing both
// numbers by their greatest common divisor.
//
// Example:
//
//	ratio := units.NewRatio(16, 8)
//	simplified := ratio.Simplify()  // Returns 2/1
func (r Ratio) Simplify() Ratio {
	gcd := greatestCommonDivisor(r.First, r.Second)
	return Ratio{
		First:  r.First / gcd,
		Second: r.Second / gcd,
	}
}

// ═══════════════════════════════════════════════════════════════
//  Comparison Operations
// ═══════════════════════════════════════════════════════════════

// Equals returns true if two ratios are equal.
func (r Ratio) Equals(other Ratio) bool {
	// Compare as decimal values
	return math.Abs(r.Value()-other.Value()) < 0.0001
}

// LessThan returns true if this ratio is less than the other.
func (r Ratio) LessThan(other Ratio) bool {
	return r.Value() < other.Value()
}

// GreaterThan returns true if this ratio is greater than the other.
func (r Ratio) GreaterThan(other Ratio) bool {
	return r.Value() > other.Value()
}

// ═══════════════════════════════════════════════════════════════
//  Utility Methods
// ═══════════════════════════════════════════════════════════════

// ApplyToWidth calculates the height for a given width based on this ratio.
//
// Example:
//
//	ratio := units.NewRatio(16, 9)  // 16:9
//	height := ratio.ApplyToWidth(1920)  // Returns 1080
func (r Ratio) ApplyToWidth(width float64) float64 {
	return width * r.Second / r.First
}

// ApplyToHeight calculates the width for a given height based on this ratio.
//
// Example:
//
//	ratio := units.NewRatio(16, 9)  // 16:9
//	width := ratio.ApplyToHeight(1080)  // Returns 1920
func (r Ratio) ApplyToHeight(height float64) float64 {
	return height * r.First / r.Second
}

// FitWidth returns the dimensions that fit within maxWidth while maintaining the ratio.
func (r Ratio) FitWidth(maxWidth float64) (width, height float64) {
	return maxWidth, r.ApplyToWidth(maxWidth)
}

// FitHeight returns the dimensions that fit within maxHeight while maintaining the ratio.
func (r Ratio) FitHeight(maxHeight float64) (width, height float64) {
	return r.ApplyToHeight(maxHeight), maxHeight
}

// ═══════════════════════════════════════════════════════════════
//  Helper Functions
// ═══════════════════════════════════════════════════════════════

// greatestCommonDivisor calculates the GCD of two numbers using Euclid's algorithm.
func greatestCommonDivisor(a, b float64) float64 {
	// Convert to integers for GCD calculation
	aInt := int64(math.Round(a * 1000))
	bInt := int64(math.Round(b * 1000))

	for bInt != 0 {
		aInt, bInt = bInt, aInt%bInt
	}

	return float64(aInt) / 1000.0
}

// ═══════════════════════════════════════════════════════════════
//  Common Ratios
// ═══════════════════════════════════════════════════════════════

var (
	// Common aspect ratios
	Ratio16x9  = NewRatio(16, 9)  // Standard widescreen
	Ratio16x10 = NewRatio(16, 10) // Computer display
	Ratio4x3   = NewRatio(4, 3)   // Traditional TV/monitor
	Ratio3x2   = NewRatio(3, 2)   // Classic 35mm film
	Ratio21x9  = NewRatio(21, 9)  // Ultrawide
	Ratio1x1   = NewRatio(1, 1)   // Square
	GoldenRatio = NewRatio(1.618, 1) // Golden ratio (φ)
)

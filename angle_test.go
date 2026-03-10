package units

import (
	"math"
	"testing"
)

func TestAngleConstructors(t *testing.T) {
	tests := []struct {
		name     string
		angle    Angle
		expected string
		unit     AngleUnit
	}{
		{"Degrees", Deg(90), "90.00deg", Degree},
		{"Gradians", Grad(100), "100.00grad", Gradian},
		{"Radians", Rad(math.Pi), "3.14rad", Radian},
		{"Turns", Turns(0.25), "0.25turn", Turn},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.angle.Unit != tt.unit {
				t.Errorf("Unit = %s, want %s", tt.angle.Unit, tt.unit)
			}
			if tt.angle.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.angle.String(), tt.expected)
			}
		})
	}
}

func TestAngleConversions(t *testing.T) {
	tests := []struct {
		name     string
		angle    Angle
		expected Angle
	}{
		// Degrees to other units
		{"360deg to turns", Deg(360).ToTurns(), Turns(1.0)},
		{"180deg to rad", Deg(180).ToRad(), Rad(math.Pi)},
		{"180deg to grad", Deg(180).ToGrad(), Grad(200)},

		// Turns to other units
		{"0.5turn to deg", Turns(0.5).ToDeg(), Deg(180)},
		{"1turn to deg", Turns(1).ToDeg(), Deg(360)},

		// Radians to other units
		{"π rad to deg", Rad(math.Pi).ToDeg(), Deg(180)},
		{"2π rad to deg", Rad(2 * math.Pi).ToDeg(), Deg(360)},

		// Gradians to other units
		{"200grad to deg", Grad(200).ToDeg(), Deg(180)},
		{"400grad to deg", Grad(400).ToDeg(), Deg(360)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.angle.Value, tt.expected.Value) {
				t.Errorf("Value = %.6f%s, want %.6f%s",
					tt.angle.Value, tt.angle.Unit, tt.expected.Value, tt.expected.Unit)
			}
			if tt.angle.Unit != tt.expected.Unit {
				t.Errorf("Unit = %s, want %s", tt.angle.Unit, tt.expected.Unit)
			}
		})
	}
}

func TestAngleNormalize(t *testing.T) {
	tests := []struct {
		name     string
		angle    Angle
		expected Angle
	}{
		{"450deg normalizes to 90deg", Deg(450), Deg(90)},
		{"720deg normalizes to 0deg", Deg(720), Deg(0)},
		{"-90deg normalizes to 270deg", Deg(-90), Deg(270)},
		{"1.5turns normalizes to 0.5turns", Turns(1.5), Turns(0.5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.angle.Normalize()
			if !almostEqual(result.Value, tt.expected.Value) {
				t.Errorf("Normalize() = %.6f%s, want %.6f%s",
					result.Value, result.Unit, tt.expected.Value, tt.expected.Unit)
			}
		})
	}
}

func TestAngleArithmetic(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result := Deg(90).Add(Deg(45))
		if result.Value != 135.0 {
			t.Errorf("Add() = %.2f, want 135.00", result.Value)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		result := Deg(180).Sub(Deg(90))
		if result.Value != 90.0 {
			t.Errorf("Sub() = %.2f, want 90.00", result.Value)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		result := Deg(45).Mul(2.0)
		if result.Value != 90.0 {
			t.Errorf("Mul() = %.2f, want 90.00", result.Value)
		}
	})

	t.Run("Div", func(t *testing.T) {
		result := Deg(180).Div(2.0)
		if result.Value != 90.0 {
			t.Errorf("Div() = %.2f, want 90.00", result.Value)
		}
	})
}

func TestAngleComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Angle
		b           Angle
		lessThan    bool
		greaterThan bool
	}{
		{"45deg < 90deg", Deg(45), Deg(90), true, false},
		{"180deg > 90deg", Deg(180), Deg(90), false, true},
		{"0.25turn < 0.5turn", Turns(0.25), Turns(0.5), true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.a.LessThan(tt.b) != tt.lessThan {
				t.Errorf("LessThan() = %v, want %v", tt.a.LessThan(tt.b), tt.lessThan)
			}
			if tt.a.GreaterThan(tt.b) != tt.greaterThan {
				t.Errorf("GreaterThan() = %v, want %v", tt.a.GreaterThan(tt.b), tt.greaterThan)
			}
		})
	}
}

func TestAngleIsZero(t *testing.T) {
	tests := []struct {
		name     string
		angle    Angle
		expected bool
	}{
		{"Zero degrees", Deg(0), true},
		{"Non-zero degrees", Deg(45), false},
		{"Zero turns", Turns(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.angle.IsZero() != tt.expected {
				t.Errorf("IsZero() = %v, want %v", tt.angle.IsZero(), tt.expected)
			}
		})
	}
}

func TestParseAngle(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Angle
		wantError bool
	}{
		{"Degree unit", "180deg", Deg(180), false},
		{"Gradian unit", "200grad", Grad(200), false},
		{"Radian unit", "3.141592653589793rad", Rad(math.Pi), false},
		{"Turn unit", "0.5turn", Turns(0.5), false},
		{"Bare number defaults to deg", "90", Deg(90), false},
		{"Whitespace", "  45deg  ", Deg(45), false},
		{"Space between number and unit", "180 deg", Deg(180), false},
		{"Case insensitive unit", "0.5TURN", Turns(0.5), false},
		{"Scientific notation", "1e2deg", Deg(100), false},
		{"Empty", "", Angle{}, true},
		{"Missing number", "deg", Angle{}, true},
		{"Unsupported unit", "10foo", Angle{}, true},
		{"Non-finite", "NaNdeg", Angle{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAngle(tt.input)
			if tt.wantError {
				if err == nil {
					t.Fatalf("ParseAngle(%q) expected error, got nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseAngle(%q) unexpected error: %v", tt.input, err)
			}
			if got.Unit != tt.want.Unit {
				t.Fatalf("ParseAngle(%q) unit = %s, want %s", tt.input, got.Unit, tt.want.Unit)
			}
			if !almostEqual(got.Value, tt.want.Value) {
				t.Fatalf("ParseAngle(%q) value = %.12f, want %.12f", tt.input, got.Value, tt.want.Value)
			}
		})
	}
}

func TestMustParseAngle(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		a := MustParseAngle("90deg")
		if a.Unit != Degree || !almostEqual(a.Value, 90) {
			t.Fatalf("MustParseAngle(valid) = %+v, want 90deg", a)
		}
	})

	t.Run("Invalid panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("MustParseAngle(invalid) expected panic")
			}
		}()
		_ = MustParseAngle("invalid")
	})
}

func BenchmarkAngleToDeg(b *testing.B) {
	angle := Rad(math.Pi)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = angle.ToDeg()
	}
}

func BenchmarkAngleNormalize(b *testing.B) {
	angle := Deg(450)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = angle.Normalize()
	}
}

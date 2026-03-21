package units

import (
	"testing"
)

func TestLengthConstructors(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		expected string
		unit     LengthUnit
	}{
		{"Pixels", Px(100), "100.00px", PX},
		{"Centimeters", Cm(2.54), "2.54cm", CM},
		{"Millimeters", Mm(25.4), "25.40mm", MM},
		{"Inches", In(1), "1.00in", IN},
		{"Points", Pt(12), "12.00pt", PT},
		{"Picas", Pc(1), "1.00pc", PC},
		{"Em", Em(1.5), "1.50em", EM},
		{"Rem", Rem(2), "2.00rem", REM},
		{"Ch", Ch(40), "40.00ch", CH},
		{"Viewport width", Vw(100), "100.00vw", VW},
		{"Viewport height", Vh(50), "50.00vh", VH},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.length.Unit != tt.unit {
				t.Errorf("Unit = %s, want %s", tt.length.Unit, tt.unit)
			}
			if tt.length.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.length.String(), tt.expected)
			}
		})
	}
}

func TestLengthIsZero(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		expected bool
	}{
		{"Zero pixels", Px(0), true},
		{"Zero em", Em(0), true},
		{"Non-zero pixels", Px(100), false},
		{"Non-zero em", Em(1.5), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.length.IsZero() != tt.expected {
				t.Errorf("IsZero() = %v, want %v", tt.length.IsZero(), tt.expected)
			}
		})
	}
}

func TestLengthIsAbsolute(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		expected bool
	}{
		{"Pixels", Px(100), true},
		{"Centimeters", Cm(2.54), true},
		{"Inches", In(1), true},
		{"Points", Pt(12), true},
		{"Em (font-relative)", Em(1.5), false},
		{"Viewport width", Vw(100), false},
		{"Container width", Cqw(50), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.length.IsAbsolute() != tt.expected {
				t.Errorf("IsAbsolute() = %v, want %v", tt.length.IsAbsolute(), tt.expected)
			}
		})
	}
}

func TestLengthIsFontRelative(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		expected bool
	}{
		{"Em", Em(1.5), true},
		{"Rem", Rem(2), true},
		{"Ch", Ch(40), true},
		{"Ex", Ex(1), true},
		{"Pixels", Px(100), false},
		{"Viewport width", Vw(100), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.length.IsFontRelative() != tt.expected {
				t.Errorf("IsFontRelative() = %v, want %v", tt.length.IsFontRelative(), tt.expected)
			}
		})
	}
}

func TestLengthIsViewportRelative(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		expected bool
	}{
		{"Viewport width", Vw(100), true},
		{"Viewport height", Vh(50), true},
		{"Vmin", Vmin(10), true},
		{"Vmax", Vmax(90), true},
		{"Pixels", Px(100), false},
		{"Em", Em(1.5), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.length.IsViewportRelative() != tt.expected {
				t.Errorf("IsViewportRelative() = %v, want %v", tt.length.IsViewportRelative(), tt.expected)
			}
		})
	}
}

func TestLengthIsContainerRelative(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		expected bool
	}{
		{"Container width", Cqw(100), true},
		{"Container height", Cqh(50), true},
		{"Container inline", Cqi(75), true},
		{"Pixels", Px(100), false},
		{"Em", Em(1.5), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.length.IsContainerRelative() != tt.expected {
				t.Errorf("IsContainerRelative() = %v, want %v", tt.length.IsContainerRelative(), tt.expected)
			}
		})
	}
}

func TestLengthAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        Length
		b        Length
		expected Length
	}{
		{"Pixels", Px(100), Px(50), Px(150)},
		{"Em", Em(1.5), Em(0.5), Em(2.0)},
		{"Ch", Ch(40), Ch(10), Ch(50)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a.Add(tt.b)
			if result.Value != tt.expected.Value || result.Unit != tt.expected.Unit {
				t.Errorf("Add() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestLengthAddPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Add() did not panic with different units")
		}
	}()

	_ = Px(100).Add(Em(1.5))
}

func TestLengthSub(t *testing.T) {
	tests := []struct {
		name     string
		a        Length
		b        Length
		expected Length
	}{
		{"Pixels", Px(100), Px(30), Px(70)},
		{"Em", Em(2.0), Em(0.5), Em(1.5)},
		{"Ch", Ch(50), Ch(10), Ch(40)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a.Sub(tt.b)
			if result.Value != tt.expected.Value || result.Unit != tt.expected.Unit {
				t.Errorf("Sub() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestLengthMul(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		scalar   float64
		expected Length
	}{
		{"Pixels", Px(100), 2.0, Px(200)},
		{"Em", Em(1.5), 3.0, Em(4.5)},
		{"Ch", Ch(40), 0.5, Ch(20)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.length.Mul(tt.scalar)
			if result.Value != tt.expected.Value || result.Unit != tt.expected.Unit {
				t.Errorf("Mul() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestLengthDiv(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		scalar   float64
		expected Length
	}{
		{"Pixels", Px(100), 2.0, Px(50)},
		{"Em", Em(3.0), 2.0, Em(1.5)},
		{"Ch", Ch(80), 4.0, Ch(20)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.length.Div(tt.scalar)
			if result.Value != tt.expected.Value || result.Unit != tt.expected.Unit {
				t.Errorf("Div() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestLengthComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Length
		b           Length
		lessThan    bool
		greaterThan bool
	}{
		{"100px < 200px", Px(100), Px(200), true, false},
		{"200px > 100px", Px(200), Px(100), false, true},
		{"1.5em < 2em", Em(1.5), Em(2.0), true, false},
		{"2em > 1.5em", Em(2.0), Em(1.5), false, true},
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

func TestParseLength(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Length
		wantErr bool
	}{
		{name: "px", input: "24px", want: Px(24)},
		{name: "unitless defaults to px", input: "24", want: Px(24)},
		{name: "uppercase Q unit", input: "10Q", want: Q(10)},
		{name: "font relative", input: "1.5rem", want: Rem(1.5)},
		{name: "viewport relative", input: "-0.25vh", want: Vh(-0.25)},
		{name: "container relative", input: "4cqi", want: Cqi(4)},
		{name: "dynamic viewport", input: "2dvw", want: Length{Value: 2, Unit: DVW}},
		{name: "small viewport", input: "3svh", want: Length{Value: 3, Unit: SVH}},
		{name: "invalid empty", input: "", wantErr: true},
		{name: "invalid unit", input: "12foo", wantErr: true},
		{name: "missing number", input: "px", wantErr: true},
		{name: "percentage unsupported", input: "50%", wantErr: true},
		{name: "nan unsupported", input: "NaNpx", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLength(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseLength(%q) expected error, got nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseLength(%q) unexpected error: %v", tt.input, err)
			}
			if got.Unit != tt.want.Unit {
				t.Fatalf("ParseLength(%q) unit = %s, want %s", tt.input, got.Unit, tt.want.Unit)
			}
			if got.Value != tt.want.Value {
				t.Fatalf("ParseLength(%q) value = %.12f, want %.12f", tt.input, got.Value, tt.want.Value)
			}
		})
	}
}

func TestMustParseLength(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		l := MustParseLength("90px")
		if l.Unit != PX || l.Value != 90 {
			t.Fatalf("MustParseLength(valid) = %+v, want 90px", l)
		}
	})

	t.Run("invalid panics", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Fatal("MustParseLength(invalid) expected panic")
			}
		}()
		_ = MustParseLength("invalid")
	})
}

func TestLengthComparisonPanic(t *testing.T) {
	t.Run("LessThan panics with different units", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("LessThan() did not panic with different units")
			}
		}()
		_ = Px(100).LessThan(Em(1.5))
	})

	t.Run("GreaterThan panics with different units", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("GreaterThan() did not panic with different units")
			}
		}()
		_ = Px(100).GreaterThan(Em(1.5))
	})
}

func TestLengthRaw(t *testing.T) {
	tests := []struct {
		name     string
		length   Length
		expected float64
	}{
		{"Pixels", Px(123.45), 123.45},
		{"Em", Em(1.5), 1.5},
		{"Ch", Ch(40), 40.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.length.Raw() != tt.expected {
				t.Errorf("Raw() = %f, want %f", tt.length.Raw(), tt.expected)
			}
		})
	}
}

package units

import "testing"

func TestNumberConstructor(t *testing.T) {
	tests := []struct {
		name     string
		num      Number
		expected string
		value    float64
	}{
		{"1.5", Num(1.5), "1.5", 1.5},
		{"0.8", Num(0.8), "0.8", 0.8},
		{"2", Num(2), "2", 2.0},
		{"0", Num(0), "0", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.num.Value != tt.value {
				t.Errorf("Value = %f, want %f", tt.num.Value, tt.value)
			}
			if tt.num.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.num.String(), tt.expected)
			}
		})
	}
}

func TestNumberArithmetic(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result := Num(1.5).Add(Num(2.5))
		if result.Value != 4.0 {
			t.Errorf("Add() = %f, want 4.0", result.Value)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		result := Num(5.0).Sub(Num(2.0))
		if result.Value != 3.0 {
			t.Errorf("Sub() = %f, want 3.0", result.Value)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		result := Num(2.5).Mul(Num(4.0))
		if result.Value != 10.0 {
			t.Errorf("Mul() = %f, want 10.0", result.Value)
		}
	})

	t.Run("Div", func(t *testing.T) {
		result := Num(10.0).Div(Num(2.0))
		if result.Value != 5.0 {
			t.Errorf("Div() = %f, want 5.0", result.Value)
		}
	})

	t.Run("Pow", func(t *testing.T) {
		result := Num(2.0).Pow(3.0)
		if result.Value != 8.0 {
			t.Errorf("Pow(3) = %f, want 8.0", result.Value)
		}
	})

	t.Run("Sqrt", func(t *testing.T) {
		result := Num(9.0).Sqrt()
		if result.Value != 3.0 {
			t.Errorf("Sqrt() = %f, want 3.0", result.Value)
		}
	})
}

func TestNumberComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Number
		b           Number
		lessThan    bool
		greaterThan bool
		equals      bool
	}{
		{"1.5 < 2.0", Num(1.5), Num(2.0), true, false, false},
		{"3.0 > 2.0", Num(3.0), Num(2.0), false, true, false},
		{"2.0 = 2.0", Num(2.0), Num(2.0), false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.a.LessThan(tt.b) != tt.lessThan {
				t.Errorf("LessThan() = %v, want %v", tt.a.LessThan(tt.b), tt.lessThan)
			}
			if tt.a.GreaterThan(tt.b) != tt.greaterThan {
				t.Errorf("GreaterThan() = %v, want %v", tt.a.GreaterThan(tt.b), tt.greaterThan)
			}
			if tt.a.Equals(tt.b) != tt.equals {
				t.Errorf("Equals() = %v, want %v", tt.a.Equals(tt.b), tt.equals)
			}
		})
	}
}

func TestNumberSign(t *testing.T) {
	tests := []struct {
		name       string
		num        Number
		isPositive bool
		isNegative bool
		isZero     bool
	}{
		{"Positive", Num(5.0), true, false, false},
		{"Negative", Num(-5.0), false, true, false},
		{"Zero", Num(0), false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.num.IsPositive() != tt.isPositive {
				t.Errorf("IsPositive() = %v, want %v", tt.num.IsPositive(), tt.isPositive)
			}
			if tt.num.IsNegative() != tt.isNegative {
				t.Errorf("IsNegative() = %v, want %v", tt.num.IsNegative(), tt.isNegative)
			}
			if tt.num.IsZero() != tt.isZero {
				t.Errorf("IsZero() = %v, want %v", tt.num.IsZero(), tt.isZero)
			}
		})
	}
}

func TestNumberAbs(t *testing.T) {
	tests := []struct {
		name     string
		num      Number
		expected float64
	}{
		{"Positive", Num(5.0), 5.0},
		{"Negative", Num(-5.0), 5.0},
		{"Zero", Num(0), 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.num.Abs()
			if result.Value != tt.expected {
				t.Errorf("Abs() = %f, want %f", result.Value, tt.expected)
			}
		})
	}
}

func TestNumberRounding(t *testing.T) {
	t.Run("Round", func(t *testing.T) {
		tests := []struct {
			value    float64
			expected float64
		}{
			{1.4, 1.0},
			{1.5, 2.0},
			{1.6, 2.0},
			{-1.4, -1.0},
			{-1.5, -2.0},
		}
		for _, tt := range tests {
			result := Num(tt.value).Round()
			if result.Value != tt.expected {
				t.Errorf("Round(%f) = %f, want %f", tt.value, result.Value, tt.expected)
			}
		}
	})

	t.Run("Floor", func(t *testing.T) {
		result := Num(1.9).Floor()
		if result.Value != 1.0 {
			t.Errorf("Floor() = %f, want 1.0", result.Value)
		}
	})

	t.Run("Ceil", func(t *testing.T) {
		result := Num(1.1).Ceil()
		if result.Value != 2.0 {
			t.Errorf("Ceil() = %f, want 2.0", result.Value)
		}
	})
}

func TestNumberClamp(t *testing.T) {
	tests := []struct {
		name     string
		num      Number
		min      float64
		max      float64
		expected float64
	}{
		{"Clamp above max", Num(1.8), 0, 1, 1.0},
		{"Clamp below min", Num(-0.5), 0, 1, 0.0},
		{"Within range", Num(0.5), 0, 1, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.num.Clamp(tt.min, tt.max)
			if result.Value != tt.expected {
				t.Errorf("Clamp(%f, %f) = %f, want %f", tt.min, tt.max, result.Value, tt.expected)
			}
		})
	}
}

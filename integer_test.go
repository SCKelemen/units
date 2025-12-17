package units

import "testing"

func TestIntegerConstructor(t *testing.T) {
	tests := []struct {
		name     string
		num      Integer
		expected string
		value    int64
	}{
		{"10", Int(10), "10", 10},
		{"0", Int(0), "0", 0},
		{"-5", Int(-5), "-5", -5},
		{"999", Int(999), "999", 999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.num.Value != tt.value {
				t.Errorf("Value = %d, want %d", tt.num.Value, tt.value)
			}
			if tt.num.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.num.String(), tt.expected)
			}
		})
	}
}

func TestIntegerArithmetic(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result := Int(10).Add(Int(5))
		if result.Value != 15 {
			t.Errorf("Add() = %d, want 15", result.Value)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		result := Int(10).Sub(Int(3))
		if result.Value != 7 {
			t.Errorf("Sub() = %d, want 7", result.Value)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		result := Int(5).Mul(Int(3))
		if result.Value != 15 {
			t.Errorf("Mul() = %d, want 15", result.Value)
		}
	})

	t.Run("Div", func(t *testing.T) {
		result := Int(10).Div(Int(3))
		if result.Value != 3 {
			t.Errorf("Div() = %d, want 3", result.Value)
		}
	})

	t.Run("Mod", func(t *testing.T) {
		result := Int(10).Mod(Int(3))
		if result.Value != 1 {
			t.Errorf("Mod() = %d, want 1", result.Value)
		}
	})

	t.Run("Pow", func(t *testing.T) {
		result := Int(2).Pow(3)
		if result.Value != 8 {
			t.Errorf("Pow(3) = %d, want 8", result.Value)
		}
	})
}

func TestIntegerComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Integer
		b           Integer
		lessThan    bool
		greaterThan bool
		equals      bool
	}{
		{"5 < 10", Int(5), Int(10), true, false, false},
		{"15 > 10", Int(15), Int(10), false, true, false},
		{"10 = 10", Int(10), Int(10), false, false, true},
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

func TestIntegerSign(t *testing.T) {
	tests := []struct {
		name       string
		num        Integer
		isPositive bool
		isNegative bool
		isZero     bool
	}{
		{"Positive", Int(10), true, false, false},
		{"Negative", Int(-10), false, true, false},
		{"Zero", Int(0), false, false, true},
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

func TestIntegerParity(t *testing.T) {
	tests := []struct {
		name   string
		num    Integer
		isEven bool
		isOdd  bool
	}{
		{"Even", Int(10), true, false},
		{"Odd", Int(11), false, true},
		{"Zero is even", Int(0), true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.num.IsEven() != tt.isEven {
				t.Errorf("IsEven() = %v, want %v", tt.num.IsEven(), tt.isEven)
			}
			if tt.num.IsOdd() != tt.isOdd {
				t.Errorf("IsOdd() = %v, want %v", tt.num.IsOdd(), tt.isOdd)
			}
		})
	}
}

func TestIntegerAbs(t *testing.T) {
	tests := []struct {
		name     string
		num      Integer
		expected int64
	}{
		{"Positive", Int(10), 10},
		{"Negative", Int(-10), 10},
		{"Zero", Int(0), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.num.Abs()
			if result.Value != tt.expected {
				t.Errorf("Abs() = %d, want %d", result.Value, tt.expected)
			}
		})
	}
}

func TestIntegerMinMax(t *testing.T) {
	t.Run("Minimum", func(t *testing.T) {
		result := Int(10).Minimum(Int(5))
		if result.Value != 5 {
			t.Errorf("Minimum() = %d, want 5", result.Value)
		}
	})

	t.Run("Maximum", func(t *testing.T) {
		result := Int(10).Maximum(Int(5))
		if result.Value != 10 {
			t.Errorf("Maximum() = %d, want 10", result.Value)
		}
	})
}

func TestIntegerClamp(t *testing.T) {
	tests := []struct {
		name     string
		num      Integer
		min      int64
		max      int64
		expected int64
	}{
		{"Clamp above max", Int(150), 0, 100, 100},
		{"Clamp below min", Int(-10), 0, 100, 0},
		{"Within range", Int(50), 0, 100, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.num.Clamp(tt.min, tt.max)
			if result.Value != tt.expected {
				t.Errorf("Clamp(%d, %d) = %d, want %d", tt.min, tt.max, result.Value, tt.expected)
			}
		})
	}
}

func TestIntegerFloat(t *testing.T) {
	num := Int(10)
	f := num.Float()
	if f != 10.0 {
		t.Errorf("Float() = %f, want 10.0", f)
	}
}

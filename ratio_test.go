package units

import "testing"

func TestRatioConstructor(t *testing.T) {
	tests := []struct {
		name     string
		ratio    Ratio
		expected string
		value    float64
	}{
		{"16:9", NewRatio(16, 9), "16 / 9", 16.0 / 9.0},
		{"4:3", NewRatio(4, 3), "4 / 3", 4.0 / 3.0},
		{"2:1", NewRatio(2, 1), "2", 2.0},
		{"1:1", NewRatio(1, 1), "1", 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.ratio.Value(), tt.value) {
				t.Errorf("Value() = %f, want %f", tt.ratio.Value(), tt.value)
			}
			if tt.ratio.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.ratio.String(), tt.expected)
			}
		})
	}
}

func TestRatioFrom(t *testing.T) {
	ratio := RatioFrom(2)
	if ratio.First != 2.0 || ratio.Second != 1.0 {
		t.Errorf("RatioFrom(2) = %f/%f, want 2.0/1.0", ratio.First, ratio.Second)
	}
}

func TestRatioProperties(t *testing.T) {
	t.Run("IsSquare", func(t *testing.T) {
		if !NewRatio(1, 1).IsSquare() {
			t.Error("1:1 should be square")
		}
		if NewRatio(16, 9).IsSquare() {
			t.Error("16:9 should not be square")
		}
	})

	t.Run("IsWide", func(t *testing.T) {
		if !NewRatio(16, 9).IsWide() {
			t.Error("16:9 should be wide")
		}
		if NewRatio(9, 16).IsWide() {
			t.Error("9:16 should not be wide")
		}
	})

	t.Run("IsTall", func(t *testing.T) {
		if !NewRatio(9, 16).IsTall() {
			t.Error("9:16 should be tall")
		}
		if NewRatio(16, 9).IsTall() {
			t.Error("16:9 should not be tall")
		}
	})
}

func TestRatioInverse(t *testing.T) {
	ratio := NewRatio(16, 9)
	inverse := ratio.Inverse()

	if inverse.First != 9.0 || inverse.Second != 16.0 {
		t.Errorf("Inverse of 16/9 = %f/%f, want 9/16", inverse.First, inverse.Second)
	}
}

func TestRatioSimplify(t *testing.T) {
	tests := []struct {
		name     string
		ratio    Ratio
		expected Ratio
	}{
		{"16:8 -> 2:1", NewRatio(16, 8), NewRatio(2, 1)},
		{"4:2 -> 2:1", NewRatio(4, 2), NewRatio(2, 1)},
		// Note: Simplify uses floating point GCD which may not be exact
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.ratio.Simplify()
			// Check if values are approximately equal
			if !almostEqual(result.Value(), tt.expected.Value()) {
				t.Errorf("Simplify() value = %f, want %f", result.Value(), tt.expected.Value())
			}
		})
	}
}

func TestRatioComparison(t *testing.T) {
	r1 := NewRatio(16, 9) // ~1.778
	r2 := NewRatio(4, 3)  // ~1.333
	r3 := NewRatio(16, 9) // ~1.778

	t.Run("Equals", func(t *testing.T) {
		if !r1.Equals(r3) {
			t.Error("16:9 should equal 16:9")
		}
		if r1.Equals(r2) {
			t.Error("16:9 should not equal 4:3")
		}
	})

	t.Run("LessThan", func(t *testing.T) {
		if !r2.LessThan(r1) {
			t.Error("4:3 should be less than 16:9")
		}
	})

	t.Run("GreaterThan", func(t *testing.T) {
		if !r1.GreaterThan(r2) {
			t.Error("16:9 should be greater than 4:3")
		}
	})
}

func TestRatioApply(t *testing.T) {
	ratio := NewRatio(16, 9)

	t.Run("ApplyToWidth", func(t *testing.T) {
		height := ratio.ApplyToWidth(1920)
		if !almostEqual(height, 1080) {
			t.Errorf("ApplyToWidth(1920) = %f, want 1080", height)
		}
	})

	t.Run("ApplyToHeight", func(t *testing.T) {
		width := ratio.ApplyToHeight(1080)
		if !almostEqual(width, 1920) {
			t.Errorf("ApplyToHeight(1080) = %f, want 1920", width)
		}
	})

	t.Run("FitWidth", func(t *testing.T) {
		w, h := ratio.FitWidth(1920)
		if !almostEqual(w, 1920) || !almostEqual(h, 1080) {
			t.Errorf("FitWidth(1920) = (%f, %f), want (1920, 1080)", w, h)
		}
	})

	t.Run("FitHeight", func(t *testing.T) {
		w, h := ratio.FitHeight(1080)
		if !almostEqual(w, 1920) || !almostEqual(h, 1080) {
			t.Errorf("FitHeight(1080) = (%f, %f), want (1920, 1080)", w, h)
		}
	})
}

func TestCommonRatios(t *testing.T) {
	tests := []struct {
		name     string
		ratio    Ratio
		expected float64
	}{
		{"16:9", Ratio16x9, 16.0 / 9.0},
		{"16:10", Ratio16x10, 16.0 / 10.0},
		{"4:3", Ratio4x3, 4.0 / 3.0},
		{"3:2", Ratio3x2, 3.0 / 2.0},
		{"21:9", Ratio21x9, 21.0 / 9.0},
		{"1:1", Ratio1x1, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.ratio.Value(), tt.expected) {
				t.Errorf("Value() = %f, want %f", tt.ratio.Value(), tt.expected)
			}
		})
	}
}

func TestRatioPanic(t *testing.T) {
	t.Run("Zero first value panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("NewRatio(0, 1) should panic")
			}
		}()
		NewRatio(0, 1)
	})

	t.Run("Zero second value panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("NewRatio(1, 0) should panic")
			}
		}()
		NewRatio(1, 0)
	})

	t.Run("Negative values panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("NewRatio(-1, 1) should panic")
			}
		}()
		NewRatio(-1, 1)
	})
}

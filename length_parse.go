package units

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var lengthUnitAliases = map[string]LengthUnit{
	"px":    PX,
	"cm":    CM,
	"mm":    MM,
	"q":     QQ,
	"in":    IN,
	"pt":    PT,
	"pc":    PC,
	"em":    EM,
	"rem":   REM,
	"ex":    EX,
	"rex":   REX,
	"cap":   CAP,
	"rcap":  RCAP,
	"ch":    CH,
	"rch":   RCH,
	"ic":    IC,
	"ric":   RIC,
	"lh":    LH,
	"rlh":   RLH,
	"vw":    VW,
	"vh":    VH,
	"vmin":  VMIN,
	"vmax":  VMAX,
	"vb":    VB,
	"vi":    VI,
	"svw":   SVW,
	"svh":   SVH,
	"svb":   SVB,
	"svi":   SVI,
	"lvw":   LVW,
	"lvh":   LVH,
	"lvb":   LVB,
	"lvi":   LVI,
	"dvw":   DVW,
	"dvh":   DVH,
	"dvb":   DVB,
	"dvi":   DVI,
	"cqw":   CQW,
	"cqh":   CQH,
	"cqi":   CQI,
	"cqb":   CQB,
	"cqmin": CQMIN,
	"cqmax": CQMAX,
}

// ParseLength parses a CSS length string into a Length.
//
// Supported formats:
//   - "<number><unit>" where unit is any known CSS length unit
//   - "<number>" (defaults to px)
//
// The parser is case-insensitive and accepts optional whitespace.
func ParseLength(input string) (Length, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return Length{}, fmt.Errorf("invalid length %q: empty value", input)
	}

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
		return Length{}, fmt.Errorf("invalid length %q: missing numeric value", input)
	}

	value, err := strconv.ParseFloat(numberPart, 64)
	if err != nil {
		return Length{}, fmt.Errorf("invalid length %q: %w", input, err)
	}
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return Length{}, fmt.Errorf("invalid length %q: value must be finite", input)
	}

	if unitPart == "" {
		return Px(value), nil
	}
	unit, ok := lengthUnitAliases[unitPart]
	if !ok {
		return Length{}, fmt.Errorf("invalid length %q: unsupported unit %q", input, unitPart)
	}
	return Length{Value: value, Unit: unit}, nil
}

// MustParseLength parses a CSS length string and panics if parsing fails.
func MustParseLength(input string) Length {
	length, err := ParseLength(input)
	if err != nil {
		panic(err)
	}
	return length
}

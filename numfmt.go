package numfmt

import (
	"strconv"
	"strings"
	"math"
)

func FloatCommas(f float64, prec int) string {
	intPart, floatPart := math.Modf(f)
	if math.Abs(intPart) > 0 {
		intStr := IntCommas(int(intPart))
		floatStr := strconv.FormatFloat(floatPart, 'f', prec, 64)
		if frac := strings.Index(floatStr, "."); frac != -1 {
			floatStr = floatStr[frac:]
		}
		return intStr + floatStr
	} else {
		return strconv.FormatFloat(floatPart, 'f', prec, 64)
	}
}

func IntCommas(n int) string {
	var sign string
	if n < 0 {
		sign = "-"
		n *= -1
	}
	digits := strconv.Itoa(n)
	pad := (3 - len(digits) % 3) % 3
	groups := []string{sign + digits[0:3-pad]} // special case first group to include sign
	for i, end := 3, 6; i < len(digits)+pad; i, end = end, end + 3 {
		groups = append(groups, digits[i-pad:end-pad])
	}
	return strings.Join(groups, ",")
}
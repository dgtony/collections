package math

import (
	"math"

	d "github.com/shopspring/decimal"
)

// Exact rounding implementation.
// Beware very expensive computation!
func RoundFixedDecimal(x float64, precision int) float64 {
	rounded, _ := d.NewFromFloat(x).Round(int32(precision)).Float64()
	return rounded
}

// Implementation was found in old versions of CockroachDB.
//
// Beware! This works correctly for banker’s rounding,
// but uses some undefined behavior of Go. The conversion
// of v (a float64) to uint64 is not well defined and
// works differently on amd64 and arm.
func RoundFixed(x float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))

	if pow == 0 {
		// Rounding to so many digits on the left that we're underflowing.
		// Avoid a NaN below.
		return 0
	}
	if math.Abs(x*pow) > 1e17 {
		// Rounding touches decimals below float precision; the operation
		// is a no-op.
		return x
	}

	v, frac := math.Modf(x * pow)
	// The following computation implements unbiased rounding, also
	// called bankers' rounding. It ensures that values that fall
	// exactly between two integers get equal chance to be rounded up or
	// down.
	if x > 0.0 {
		if frac > 0.5 || (frac == 0.5 && uint64(v)%2 != 0) {
			v += 1.0
		}
	} else {
		if frac < -0.5 || (frac == -0.5 && uint64(v)%2 != 0) {
			v -= 1.0
		}
	}

	return v / pow
}

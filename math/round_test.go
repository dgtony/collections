package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFixed(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		expected  float64
	}{
		{1.23456789, -1, 0},
		{1.23456789, 0, 1},
		{1.23456789, 1, 1.2},

		{-1.789, -1, 0},
		{-1.789, 0, -2},
		{-1.789, 1, -1.8},

		{0.000123456, -1, 0},
		{0.000123456, 0, 0},
		{0.000123456, 4, 0.0001},
		{0.001234567, 8, 0.00123457},

		{1234.56789, -3, 1000},
		{1234.56789, -1, 1230},
		{1234.56789, 0, 1235},
		{1234.56789, 1, 1234.6},
		{1234.56789, 3, 1234.568},

		{0, -1, 0},
		{0, 0, 0},
		{0, 2, 0},
	}

	for _, method := range []struct {
		name string
		fun  func(float64, int) float64
	}{
		{"fixed-cockroach", RoundFixed},
		{"fixed-decimal", RoundFixedDecimal},
	} {
		for _, tt := range tests {
			rounded := method.fun(tt.value, tt.precision)
			assert.InDelta(t, tt.expected, rounded, 1e-10, "value: %v, precision: %v",
				tt.value, tt.precision, method.name)
		}
	}
}

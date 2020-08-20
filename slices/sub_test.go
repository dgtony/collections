package slices

import (
	"testing"

	"github.com/dgtony/collections/polymorph"
	"github.com/stretchr/testify/assert"
)

func TestCarve(t *testing.T) {
	slice := []int{1, 2, 3}

	cases := []struct {
		from, to    int
		expected    []int
		errExpected bool
	}{
		{
			from:        0,
			to:          0,
			expected:    []int{1, 2, 3},
			errExpected: false,
		},
		{
			from:        0,
			to:          2,
			expected:    []int{1, 2},
			errExpected: false,
		},
		{
			from:        -1,
			to:          0,
			expected:    []int{3},
			errExpected: false,
		},
		{
			from:        1,
			to:          0,
			expected:    []int{2, 3},
			errExpected: false,
		},
		{
			from:        2,
			to:          2,
			expected:    []int{},
			errExpected: false,
		},
		{
			from:        0,
			to:          3,
			expected:    []int{1, 2, 3},
			errExpected: false,
		},
		{
			from:        -1,
			to:          1,
			errExpected: true,
		},
		{
			from:        -10,
			to:          -1,
			errExpected: true,
		},
		{
			from:        0,
			to:          10,
			errExpected: true,
		},
		{
			from:        2,
			to:          1,
			errExpected: true,
		},
	}

	for _, tt := range cases {
		res, err := Carve(polymorph.FromInts(slice), tt.from, tt.to)

		if tt.errExpected {
			assert.Error(t, err)
		} else {
			assert.Equal(t, tt.expected, polymorph.ToInts(res))
			assert.NoError(t, err)
		}
	}

}

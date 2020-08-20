package slices

import (
	"errors"
)

var ErrBadInterval = errors.New("bad interval")

// Subslicing with negative indexes allowed
func Carve(s []interface{}, from, to int) ([]interface{}, error) {
	var (
		fromIdx = from
		toIdx   = to
	)

	if fromIdx < 0 {
		fromIdx += len(s)
	}

	if toIdx < 0 {
		toIdx += len(s)
	} else if toIdx == 0 {
		// special case: take last `from`
		toIdx = len(s)
	}

	if fromIdx < 0 || fromIdx >= len(s) || toIdx < 0 || toIdx > len(s) || toIdx < fromIdx {
		return nil, ErrBadInterval
	}

	return s[fromIdx:toIdx], nil
}

package maybe

import "fmt"

/* Floats */

type Float struct {
	Value float64
	Valid bool
}

var NothingFloat = Float{Valid: false}

func JustFloat(x float64) Float {
	return Float{Value: x, Valid: true}
}

func (f Float) String() string {
	if !f.Valid {
		return "Nothing"
	}

	return fmt.Sprintf("Just(%f)", f.Value)
}

/* Integers */

type Int struct {
	Value int64
	Valid bool
}

var NothingInt = Int{Valid: false}

func JustInt(x int64) Int {
	return Int{Value: x, Valid: true}
}

func (i Int) String() string {
	if !i.Valid {
		return "Nothing"
	}

	return fmt.Sprintf("Just(%d)", i.Value)
}

/* Functors */

func MapFloat(v Float, f func(float64) float64) Float {
	if !v.Valid {
		return NothingFloat
	}

	return JustFloat(f(v.Value))
}

func MapInt(v Int, f func(int64) int64) Int {
	if !v.Valid {
		return NothingInt
	}

	return JustInt(f(v.Value))
}

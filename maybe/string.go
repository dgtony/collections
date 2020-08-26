package maybe

import "fmt"

type String struct {
	Value string
	Valid bool
}

var NothingString = String{Valid: false}

func JustString(x string) String {
	return String{Value: x, Valid: true}
}

func (s String) String() string {
	if !s.Valid {
		return "Nothing"
	}

	return fmt.Sprintf("Just(%s)", s.Value)
}

func MapString(v String, f func(string) string) String {
	if !v.Valid {
		return NothingString
	}

	return JustString(f(v.Value))
}

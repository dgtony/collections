package polymorph

func FromInts(vs []int) []interface{} {
	untyped := make([]interface{}, len(vs))
	for i := 0; i < len(vs); i++ {
		untyped[i] = vs[i]
	}

	return untyped
}

func FromInt64s(vs []int64) []interface{} {
	untyped := make([]interface{}, len(vs))
	for i := 0; i < len(vs); i++ {
		untyped[i] = vs[i]
	}

	return untyped
}

func FromUint64s(vs []uint64) []interface{} {
	untyped := make([]interface{}, len(vs))
	for i := 0; i < len(vs); i++ {
		untyped[i] = vs[i]
	}

	return untyped
}

func FromFloat64s(vs []float64) []interface{} {
	untyped := make([]interface{}, len(vs))
	for i := 0; i < len(vs); i++ {
		untyped[i] = vs[i]
	}

	return untyped
}

func FromStrings(vs []string) []interface{} {
	untyped := make([]interface{}, len(vs))
	for i := 0; i < len(vs); i++ {
		untyped[i] = vs[i]
	}

	return untyped
}

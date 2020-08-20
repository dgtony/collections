package polymorph

func ToInts(vs []interface{}) []int {
	typed := make([]int, len(vs))
	for i := 0; i < len(vs); i++ {
		typed[i] = vs[i].(int)
	}

	return typed
}

func ToInt64s(vs []interface{}) []int64 {
	typed := make([]int64, len(vs))
	for i := 0; i < len(vs); i++ {
		typed[i] = vs[i].(int64)
	}

	return typed
}

func ToUint64s(vs []interface{}) []uint64 {
	typed := make([]uint64, len(vs))
	for i := 0; i < len(vs); i++ {
		typed[i] = vs[i].(uint64)
	}

	return typed
}

func ToFloat64s(vs []interface{}) []float64 {
	typed := make([]float64, len(vs))
	for i := 0; i < len(vs); i++ {
		typed[i] = vs[i].(float64)
	}

	return typed
}

func ToStrings(vs []interface{}) []string {
	typed := make([]string, len(vs))
	for i := 0; i < len(vs); i++ {
		typed[i] = vs[i].(string)
	}

	return typed
}

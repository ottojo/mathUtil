package mathUtil

// Calculates the sum of a set of values
func Sum(values []float64) float64 {
	var r float64
	for _, v := range values {
		r += v
	}
	return r
}

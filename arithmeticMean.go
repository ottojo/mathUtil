package mathUtil

// Calculates the arithmetic mean of a set of values
func ArithmeticMean(values []float64) float64 {
	return Sum(values) / float64(len(values))
}

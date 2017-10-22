package mathUtil

import "math"

// Calculates the empiric Standard Deviation of a set of values.
func StandardDeviation(values []float64) float64 {
	arithmeticMean := ArithmeticMean(values)
	var squaredDifferences []float64
	for _, v := range values {
		difference := v - arithmeticMean
		squaredDifferences = append(squaredDifferences, difference*difference)
	}
	return math.Sqrt((1 / float64(len(values)-1)) * Sum(squaredDifferences))
}

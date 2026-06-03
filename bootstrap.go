package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/Yinan920/trimmedmean"
)

// runBootstrap estimates and prints the bootstrap standard error of three
// estimators of central tendency for one sample: the ordinary mean, the 0.05
// trimmed mean, and the median. A smaller standard error means the estimator
// varies less under resampling, i.e. it is more stable for this data. For a
// distribution with outliers the trimmed mean and median typically beat the
// ordinary mean, which motivates the trimmed mean as a robust estimator.
func runBootstrap(label string, data []float64, replications int, rng *rand.Rand) {
	means := make([]float64, replications)
	trims := make([]float64, replications)
	medians := make([]float64, replications)
	resample := make([]float64, len(data))

	for r := 0; r < replications; r++ {
		for i := range resample {
			resample[i] = data[rng.Intn(len(data))]
		}
		means[r], _ = trimmedmean.Compute(resample)
		trims[r], _ = trimmedmean.Compute(resample, 0.05)
		medians[r] = median(resample)
	}

	fmt.Printf("\nBootstrap standard errors for %s (%d replications)\n", label, replications)
	fmt.Printf("  mean         : %.6f\n", standardDeviation(means))
	fmt.Printf("  trimmed 0.05 : %.6f\n", standardDeviation(trims))
	fmt.Printf("  median       : %.6f\n", standardDeviation(medians))
}

// median returns the middle value of the data without mutating the caller's
// slice, averaging the two central values for an even count.
func median(values []float64) float64 {
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return sorted[n/2]
	}
	return (sorted[n/2-1] + sorted[n/2]) / 2
}

// standardDeviation returns the sample standard deviation (denominator n-1).
func standardDeviation(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	mean := sum / float64(len(values))

	var sumSquares float64
	for _, value := range values {
		diff := value - mean
		sumSquares += diff * diff
	}
	return math.Sqrt(sumSquares / float64(len(values)-1))
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Function to calculate the median of a slice of numbers
func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2.0
	}
	return data[n/2]
}

// Function to perform bootstrap sampling
func bootstrap(data []float64, numSamples int) []float64 {
	rand.Seed(time.Now().UnixNano())
	bootstrapSamples := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		// Create a resample by picking random elements with replacement
		resample := make([]float64, len(data))
		for j := 0; j < len(data); j++ {
			resample[j] = data[rand.Intn(len(data))]
		}
		// Store the median of this resample
		bootstrapSamples[i] = median(resample)
	}
	return bootstrapSamples
}

func main() {
	// Example data
	data := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Perform bootstrap sampling with 1000 iterations
	bootstrapSamples := bootstrap(data, 1000)

	// Calculate the standard error of the median
	// First, find the median of the bootstrap samples
	stdErr := median(bootstrapSamples)

	// Output the result
	fmt.Printf("Bootstrap Standard Error of the Median: %v\n", stdErr)
}

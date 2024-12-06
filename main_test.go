package main

import (
	"testing"
)

// TestMedian tests the median function for correctness.
func TestMedian(t *testing.T) {
	// Define test cases with different kinds of data
	tests := []struct {
		data     []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},         // Odd number of elements
		{[]float64{1, 2, 3, 4}, 2.5},          // Even number of elements
		{[]float64{5, 1, 3, 2, 4}, 3},         // Unsorted input
		{[]float64{1}, 1},                     // Single element
		{[]float64{}, 0},                      // Empty slice
	}

	// Run each test case
	for _, test := range tests {
		t.Run("Testing Median", func(t *testing.T) {
			// Handle empty slice case
			if len(test.data) == 0 {
				t.Errorf("Median of an empty slice should not be calculated")
				return
			}

			result := median(test.data)
			if result != test.expected {
				t.Errorf("Expected median %v, but got %v", test.expected, result)
			}
		})
	}
}

// TestBootstrap tests the bootstrap sampling function.
func TestBootstrap(t *testing.T) {
	// Define test data
	data := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Perform bootstrap sampling with 1000 iterations
	numSamples := 1000
	bootstrapSamples := bootstrap(data, numSamples)

	// Test that the length of the bootstrap samples is correct
	if len(bootstrapSamples) != numSamples {
		t.Errorf("Expected %d bootstrap samples, but got %d", numSamples, len(bootstrapSamples))
	}

	// Ensure that the bootstrap samples contain valid median values
	for _, sample := range bootstrapSamples {
		if sample == 0 {
			t.Errorf("Bootstrap sample contains zero value, which is unexpected")
		}
	}
}

// BenchmarkBootstrap benchmarks the performance of the bootstrap sampling function.
func BenchmarkBootstrap(b *testing.B) {
	// Define test data
	data := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Benchmark the bootstrap function
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bootstrap(data, 1000)
	}
}

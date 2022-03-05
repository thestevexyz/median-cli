package main

import "testing"

func Test_CalculateMedianEven(t *testing.T) {
	numbers := []float64{
		15.925, 9881.7626, 5555.55, 3051,
	}

	var expect float64 = 4303.275

	result := calculateMedian(numbers)
	if result != expect {
		t.Errorf("calcualateMedian(%v) (with even numbers) = %v; want %v.", numbers, result, expect)
	}
}

func Test_CalculateMedianOdd(t *testing.T) {
	numbers := []float64{
		872.13, 901.178, 1255.56, 765.33, 22451.98,
	}

	var expect float64 = 901.178

	result := calculateMedian(numbers)
	if result != expect {
		t.Errorf("calcualateMedian(%v) (with odd numbers) = %v; want %v.", numbers, result, expect)
	}
}

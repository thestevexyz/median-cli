package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printErrorAndExit("No numbers were provided")
	}

	var argsParsed = []float64{}
	for _, arg := range args {
		float, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			printErrorAndExit(fmt.Sprintf("Failed parsing argument '%s' to float64", arg))
		}

		argsParsed = append(argsParsed, float)
	}

	fmt.Println(calculateMedian(argsParsed))
}

func calculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)

	middleNumber := len(numbers) / 2

	if len(numbers)%2 != 0 {
		return numbers[middleNumber]
	}

	return (numbers[middleNumber-1] + numbers[middleNumber]) / 2
}

func printErrorAndExit(message string) {
	fmt.Printf("Error: %s\n", message)
	os.Exit(1)
}

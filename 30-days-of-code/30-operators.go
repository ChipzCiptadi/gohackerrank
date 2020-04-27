// https://www.hackerrank.com/challenges/30-operators/problem
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Complete the solve function below.
func solve(mealCost float64, tipPercent int32, taxPercent int32) {
	tip := mealCost * float64(tipPercent) / 100.0
	tax := mealCost * float64(taxPercent) / 100.0

	totalCost := mealCost + tip + tax

	fmt.Printf("%.0f", math.Round(totalCost))
}

func operators() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mealCost, err := strconv.ParseFloat(readLine(reader), 64)
	checkError(err)

	tipPercentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tipPercent := int32(tipPercentTemp)

	taxPercentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	taxPercent := int32(taxPercentTemp)

	solve(mealCost, tipPercent, taxPercent)
}

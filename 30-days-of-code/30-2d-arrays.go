// https://www.hackerrank.com/challenges/30-2d-arrays/problem
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func twoDArrays() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	var maxSum int32
	maxSum = math.MinInt32

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	fmt.Printf("%d\n", maxSum)
}

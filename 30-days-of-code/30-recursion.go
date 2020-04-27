// https://www.hackerrank.com/challenges/30-recursion/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Complete the factorial function below.
func factorial(n int32) int32 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func recursion() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := factorial(n)

	fmt.Printf("%d\n", result)
}

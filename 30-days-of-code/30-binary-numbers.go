// https://www.hackerrank.com/challenges/30-binary-numbers/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func binaryNumbers() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	max := 0
	counter := 0

	for n > 0 {
		remainder := n % 2
		n = n / 2

		if remainder == 1 {
			counter++
			if counter > max {
				max = counter
			}
		} else {
			counter = 0
		}
	}

	fmt.Printf("%d\n", max)

}

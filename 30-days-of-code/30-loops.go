// https://www.hackerrank.com/challenges/30-loops/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loops() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var i int32 = 1

	for ; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

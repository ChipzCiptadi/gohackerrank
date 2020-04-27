// https://www.hackerrank.com/challenges/30-review-loop/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reviewLoop() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	T, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var i int64 = 0
	for ; i < T; i++ {
		S := readLine(reader)

		sOdd := ""
		sEven := ""

		for idx, chr := range strings.Split(S, "") {
			if idx%2 == 0 {
				sEven = sEven + chr
			} else {
				sOdd = sOdd + chr
			}
		}

		fmt.Printf("%s %s\n", sEven, sOdd)
	}
}

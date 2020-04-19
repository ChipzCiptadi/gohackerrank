// https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dictionariesAndMaps() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var phoneBook map[string]string
	phoneBook = make(map[string]string)

	for i := 0; i < int(n); i++ {
		query := readLine(reader)
		arrTemp := strings.Split(query, " ")

		phoneBook[arrTemp[0]] = arrTemp[1]
	}

	for {
		name := readLine(reader)
		if name == "" {
			break
		}

		phone := phoneBook[name]
		if phone == "" {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", name, phone)
		}
	}
}

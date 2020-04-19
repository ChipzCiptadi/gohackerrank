package main

import (
	"bufio"
	"io"
	"strings"
)

func main() {
	// helloWorld()
	// dataTypes
	// conditionalStatements()
	// classVsInstance()
	// loops()
	// reviewLoop()
	// arrays()
	// dictionariesAndMaps()
	// recursion()
	// binaryNumbers()
	twoDArrays()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
